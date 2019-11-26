// GENERATED FILE: DO NOT EDIT!

package bookstore

// Types used by the API.
// implements the service definition of book
type Book struct {
	Id     string `json:"id,omitempty"`
	Author string `json:"author,omitempty"`
	Title  string `json:"title,omitempty"`
}

// implements the service definition of listBooksResponse
type ListBooksResponse struct {
	Books []Book `json:"books,omitempty"`
}

// implements the service definition of listShelvesResponse
type ListShelvesResponse struct {
	Shelves []Shelf `json:"shelves,omitempty"`
}

// implements the service definition of shelf
type Shelf struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Theme string `json:"theme,omitempty"`
}

// implements the service definition of error
type Error struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ListShelvesResponses holds responses of ListShelves
type ListShelvesResponses struct {
	OK *ListShelvesResponse
}

// CreateShelfParameters holds parameters to CreateShelf
type CreateShelfParameters struct {
	Shelf *Shelf `json:"shelf,omitempty"`
}

// CreateShelfResponses holds responses of CreateShelf
type CreateShelfResponses struct {
	OK *Shelf
}

// GetShelfParameters holds parameters to GetShelf
type GetShelfParameters struct {
	Shelf string `json:"shelf,omitempty"`
}

// GetShelfResponses holds responses of GetShelf
type GetShelfResponses struct {
	OK      *Shelf
	Default *Error
}

// DeleteShelfParameters holds parameters to DeleteShelf
type DeleteShelfParameters struct {
	Shelf string `json:"shelf,omitempty"`
}

// ListBooksParameters holds parameters to ListBooks
type ListBooksParameters struct {
	Shelf string `json:"shelf,omitempty"`
}

// ListBooksResponses holds responses of ListBooks
type ListBooksResponses struct {
	OK      *ListBooksResponse
	Default *Error
}

// CreateBookParameters holds parameters to CreateBook
type CreateBookParameters struct {
	Shelf string `json:"shelf,omitempty"`
	Book  *Book  `json:"book,omitempty"`
}

// CreateBookResponses holds responses of CreateBook
type CreateBookResponses struct {
	OK      *Book
	Default *Error
}

// GetBookParameters holds parameters to GetBook
type GetBookParameters struct {
	Shelf string `json:"shelf,omitempty"`
	Book  string `json:"book,omitempty"`
}

// GetBookResponses holds responses of GetBook
type GetBookResponses struct {
	OK      *Book
	Default *Error
}

// DeleteBookParameters holds parameters to DeleteBook
type DeleteBookParameters struct {
	Shelf string `json:"shelf,omitempty"`
	Book  string `json:"book,omitempty"`
}
