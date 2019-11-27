// GENERATED FILE: DO NOT EDIT!

package bookstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client represents an API client.
type Client struct {
	service string
	APIKey  string
	client  *http.Client
}

// NewClient creates an API client.
func NewClient(service string, c *http.Client) *Client {
	client := &Client{}
	client.service = service
	if c != nil {
		client.client = c
	} else {
		client.client = http.DefaultClient
	}
	return client
}

// Return all shelves in the bookstore.
func (client *Client) ListShelves() (
	response *ListShelvesResponses,
	err error,
) {
	path := client.service + "/shelves"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &ListShelvesResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &ListShelvesResponse{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	default:
		break
	}
	return
}

// Create a new shelf in the bookstore.
func (client *Client) CreateShelf(
	shelf Shelf,
) (
	response *CreateShelfResponses,
	err error,
) {
	path := client.service + "/shelves"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(shelf)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &CreateShelfResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Shelf{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	default:
		break
	}
	return
}

// Delete all shelves.
func (client *Client) DeleteShelves() (
	err error,
) {
	path := client.service + "/shelves"
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return
}

// Get a single shelf resource with the given ID.
func (client *Client) GetShelf(
	shelf string,
) (
	response *GetShelfResponses,
	err error,
) {
	path := client.service + "/shelves/{shelf}"
	path = strings.Replace(path, "{shelf}", fmt.Sprintf("%v", shelf), 1)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &GetShelfResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Shelf{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

// Delete a single shelf with the given ID.
func (client *Client) DeleteShelf(
	shelf string,
) (
	err error,
) {
	path := client.service + "/shelves/{shelf}"
	path = strings.Replace(path, "{shelf}", fmt.Sprintf("%v", shelf), 1)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return
}

// Return all books in a shelf with the given ID.
func (client *Client) ListBooks(
	shelf string,
) (
	response *ListBooksResponses,
	err error,
) {
	path := client.service + "/shelves/{shelf}/books"
	path = strings.Replace(path, "{shelf}", fmt.Sprintf("%v", shelf), 1)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &ListBooksResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &ListBooksResponse{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

// Create a new book on the shelf.
func (client *Client) CreateBook(
	shelf string,
	book Book,
) (
	response *CreateBookResponses,
	err error,
) {
	path := client.service + "/shelves/{shelf}/books"
	path = strings.Replace(path, "{shelf}", fmt.Sprintf("%v", shelf), 1)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(book)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &CreateBookResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Book{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

// Get a single book with a given ID from a shelf.
func (client *Client) GetBook(
	shelf string,
	book string,
) (
	response *GetBookResponses,
	err error,
) {
	path := client.service + "/shelves/{shelf}/books/{book}"
	path = strings.Replace(path, "{shelf}", fmt.Sprintf("%v", shelf), 1)
	path = strings.Replace(path, "{book}", fmt.Sprintf("%v", book), 1)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &GetBookResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Book{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

// Delete a single book with a given ID from a shelf.
func (client *Client) DeleteBook(
	shelf string,
	book string,
) (
	err error,
) {
	path := client.service + "/shelves/{shelf}/books/{book}"
	path = strings.Replace(path, "{shelf}", fmt.Sprintf("%v", shelf), 1)
	path = strings.Replace(path, "{book}", fmt.Sprintf("%v", book), 1)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return
}
