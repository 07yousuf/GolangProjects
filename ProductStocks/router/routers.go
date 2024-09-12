package router
import (
  "ProductStocks/middleware"
  "github.com/gorilla/mux"
)
func Router() *mux.Router{
  route := mux.NewRouter

  route.HandleFunc("/api/stock/{id}",middleware.GetStock().Methods("GET","OPTIONS"))
  route.HandleFunc("/api/stock",middleware.GetStocks().Methods("GET","OPTIONS"))
  route.HandleFunc("/api/stock",middleware.CreateStocks().Methods("POST","OPTIONS"))
  route.HandleFunc("/api/stock/{id}",middleware.UpdateStocks().Methods("PUT","OPTIONS"))
  route.HandleFunc("/api/stock/{id}",middleware.DeleteStocks().Methods("DELETE","OPTIONS"))
}
