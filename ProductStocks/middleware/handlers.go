package middleware
import (
  "net/http"
  "encoding/json"
  "ProductStocks/models"
  "fmt"
  "log"
  "os"
  "strconv"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)
//response format
type response struct {
  ID      int64 `json:"id,omitempty"`
  Message string`json:"message,omitempty`
}
//connection with postgres db
func createConnection() *sql.DB{
  //load .env file 
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  //Open the connection
  db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
    if err != nil {
    panic(err)
  }
  //connection check
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Succesfully connected")
  //return the connection
  return db
}

  GetStock(w http.ResponseWriter, r http.Request){

  }
  GetStocks(w http.ResponseWriter, r http.Request){

  }
  CreateStocks(w http.ResponseWriter, r http.Request){

  }
  UpdateStocks(w http.ResponseWriter, r http.Request){

  }
  DeleteStocks(w http.ResponseWriter, r http.Request){

  }
}
