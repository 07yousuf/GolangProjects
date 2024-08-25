package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/07yousuf/BasicBookManagementSystem/PKG/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
}
type Author struct {
}

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	r.Handle("/", r)
	fmt.Println("starting server at port 8010")
	log.Fatal(http.ListenAndServe("localhost:8010", r))
}
