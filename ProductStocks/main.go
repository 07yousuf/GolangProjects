package main
import (
  "fmt"
  "net/http"
  "ProductStocks/router"
)
func main(){
  r := mux.NewRouter()
  fmt.Println("Server starting at port at 8002")
  log.Fatal(http.ListenAndServer("localhost:8002",r))
}
