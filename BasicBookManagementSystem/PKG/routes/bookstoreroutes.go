package routes

import (
	"github.com/07yousuf/BasicBookManagementSystem/PKG/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(route *mux.Router) {
	route.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	route.HandleFunc("/books/", controllers.CreateBooks).Methods("POST")
	route.HandleFunc("/books/{BookId}", controllers.GetBook).Methods("GET")
	route.HandleFunc("/books/{BookId}", controllers.UpdateBook).Methods("PUT")
	route.HandleFunc("/books/{BookId}", controllers.DeleteBook).Methods("DELETE")
}
