package routes

// Book-management-system\pkg\routes\bookstore-routes.go

import (
	"github.com/gorilla/mux"
	"github.com/mr-yash-dev/Book-management-system/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	// Create a subrouter for /api/v1
	api := router.PathPrefix("/api/v1").Subrouter()

	// Book 	routes
	api.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	api.HandleFunc("/books", controllers.GetBook).Methods("GET")
	api.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	api.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
