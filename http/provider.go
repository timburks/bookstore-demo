// GENERATED FILE: DO NOT EDIT!

package bookstore

// To create a server, first write a class that implements this interface.
// Then pass an instance of it to Initialize().
type Provider interface {

	// Return all shelves in the bookstore.
	ListShelves(responses *ListShelvesResponses) (err error)

	// Create a new shelf in the bookstore.
	CreateShelf(parameters *CreateShelfParameters, responses *CreateShelfResponses) (err error)

	// Delete all shelves.
	DeleteShelves() (err error)

	// Get a single shelf resource with the given ID.
	GetShelf(parameters *GetShelfParameters, responses *GetShelfResponses) (err error)

	// Delete a single shelf with the given ID.
	DeleteShelf(parameters *DeleteShelfParameters) (err error)

	// Return all books in a shelf with the given ID.
	ListBooks(parameters *ListBooksParameters, responses *ListBooksResponses) (err error)

	// Create a new book on the shelf.
	CreateBook(parameters *CreateBookParameters, responses *CreateBookResponses) (err error)

	// Get a single book with a given ID from a shelf.
	GetBook(parameters *GetBookParameters, responses *GetBookResponses) (err error)

	// Delete a single book with a given ID from a shelf.
	DeleteBook(parameters *DeleteBookParameters) (err error)
}
