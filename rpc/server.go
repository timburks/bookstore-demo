// Copyright 2019 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

//
// The server type implements a Bookstore server.
// Shelves and books are stored using the Cloud Datastore API.
// https://cloud.google.com/datastore/
//
type server struct{}

// NewStore creates a new data storage connection.
func (s *server) newDataStoreClient(ctx context.Context) (*datastore.Client, error) {
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	projectID := credentials.ProjectID
	if projectID == "" {
		return nil, errors.New("unable to determine project ID")
	}
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *server) ListShelves(ctx context.Context, _ *empty.Empty) (*ListShelvesResponse, error) {
	log.Printf("list shelves")
	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	q := datastore.NewQuery("shelf")
	var shelves []*Shelf
	keys, err := client.GetAll(ctx, q, &shelves)
	for i, k := range keys {
		shelves[i].Id = strconv.FormatInt(k.ID, 10)
	}
	responses := &ListShelvesResponse{
		Shelves: shelves,
	}
	return responses, nil
}

func (s *server) CreateShelf(ctx context.Context, parameters *CreateShelfParameters) (*Shelf, error) {
	log.Printf("create shelf %+v", parameters)
	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelf := parameters.Shelf
	if shelf.Id == "" && shelf.Name == "" && shelf.Theme == "" {
		return nil, status.Error(codes.InvalidArgument, "empty request body")
	}
	var k *datastore.Key
	if shelf.Id == "" {
		k = datastore.IncompleteKey("shelf", nil)
	} else {
		shelfid, _ := strconv.ParseInt(shelf.Id, 10, 64)
		k = &datastore.Key{Kind: "shelf", ID: shelfid}
	}
	k, err = client.Put(ctx, k, shelf)
	if err != nil {
		return nil, err
	}
	shelf.Id = strconv.FormatInt(k.ID, 10)
	responses := shelf
	return responses, nil
}

func (s *server) DeleteShelves(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	log.Printf("delete shelves")

	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	q := datastore.NewQuery("shelf")
	var shelves []*Shelf
	keys, err := client.GetAll(ctx, q, &shelves)
	if err != nil {
		return nil, err
	}
	err = client.DeleteMulti(ctx, keys)
	if err != nil {
		return nil, err
	}
	q = datastore.NewQuery("book")
	var books []*Book
	keys, err = client.GetAll(ctx, q, &books)
	if err != nil {
		return nil, err
	}
	err = client.DeleteMulti(ctx, keys)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, err
}

func (s *server) GetShelf(ctx context.Context, parameters *GetShelfParameters) (*Shelf, error) {
	log.Printf("get shelf %+v", parameters)

	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelfid, _ := strconv.ParseInt(parameters.Shelf, 10, 64)
	k := &datastore.Key{Kind: "shelf", ID: shelfid}
	var shelf Shelf
	err = client.Get(ctx, k, &shelf)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	shelf.Id = strconv.FormatInt(k.ID, 10)
	return &shelf, nil
}

func (s *server) DeleteShelf(ctx context.Context, parameters *DeleteShelfParameters) (*empty.Empty, error) {
	log.Printf("delete shelf %+v", parameters)

	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelfid, _ := strconv.ParseInt(parameters.Shelf, 10, 64)
	k := &datastore.Key{Kind: "shelf", ID: shelfid}
	// delete all the books with this shelf
	q := datastore.NewQuery("book").Ancestor(k)
	var books []*Book
	keys, err := client.GetAll(ctx, q, &books)
	err = client.DeleteMulti(ctx, keys)
	// delete the shelf
	err = client.Delete(ctx, k)
	return &empty.Empty{}, err
}

func (s *server) ListBooks(ctx context.Context, parameters *ListBooksParameters) (responses *ListBooksResponse, err error) {
	log.Printf("list books %+v", parameters)

	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelfid, _ := strconv.ParseInt(parameters.Shelf, 10, 64)
	ancestor := &datastore.Key{Kind: "shelf", ID: shelfid}
	q := datastore.NewQuery("book").Ancestor(ancestor)
	var books []*Book
	keys, err := client.GetAll(ctx, q, &books)
	for i, k := range keys {
		books[i].Id = strconv.FormatInt(k.ID, 10)
	}
	responses = &ListBooksResponse{
		Books: books,
	}
	return responses, nil
}

func (s *server) CreateBook(ctx context.Context, parameters *CreateBookParameters) (*Book, error) {
	log.Printf("create book %+v", parameters)

	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelfid, _ := strconv.ParseInt(parameters.Shelf, 10, 64)
	ancestor := &datastore.Key{Kind: "shelf", ID: shelfid}

	book := parameters.Book
	var k *datastore.Key
	if book.Id == "" {
		k = datastore.IncompleteKey("book", ancestor)
	} else {
		bookid, _ := strconv.ParseInt(book.Id, 10, 64)
		k = &datastore.Key{Kind: "book", ID: bookid, Parent: ancestor}
	}
	k, err = client.Put(ctx, k, book)
	if err != nil {
		return nil, err
	}
	book.Id = strconv.FormatInt(k.ID, 10)
	return book, nil
}

func (s *server) GetBook(ctx context.Context, parameters *GetBookParameters) (*Book, error) {
	log.Printf("get book %+v", parameters)

	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelfid, _ := strconv.ParseInt(parameters.Shelf, 10, 64)
	ancestor := &datastore.Key{Kind: "shelf", ID: shelfid}
	bookid, _ := strconv.ParseInt(parameters.Book, 10, 64)
	k := &datastore.Key{Kind: "book", ID: bookid, Parent: ancestor}
	var book Book
	err = client.Get(ctx, k, &book)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	book.Id = strconv.FormatInt(k.ID, 10)
	return &book, nil
}

func (s *server) DeleteBook(ctx context.Context, parameters *DeleteBookParameters) (*empty.Empty, error) {
	log.Printf("delete book %+v", parameters)
	client, err := s.newDataStoreClient(ctx)
	if err != nil {
		return nil, err
	}
	shelfid, _ := strconv.ParseInt(parameters.Shelf, 10, 64)
	ancestor := &datastore.Key{Kind: "shelf", ID: shelfid}
	bookid, _ := strconv.ParseInt(parameters.Book, 10, 64)
	k := &datastore.Key{Kind: "book", ID: bookid, Parent: ancestor}
	err = client.Delete(ctx, k)
	return &empty.Empty{}, err
}

// RunServer ...
func RunServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Printf("\nServer listening on port %v \n", port)
	RegisterBookstoreServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
