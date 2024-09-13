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
    params := mux.Vars(r)
    id,err := strconv.Atoi(params["id"])
    if err != nil{
      log.Fatalf("unable to convert the string into int. %v",err)
    }
    stock, err := getStock(int64(id))
    if err != nil {
      log.Fatalf("unable to get stock %v", err)
    }
    json.NewEncoder(w).Encode(stock)
  }

  GetAllStock(w http.ResponseWriter, r http.Request){
    stocks, err := getAllStocks()
    if err != nil {
      log.Fatalf("unable to get all stock %v",err)
    }
    json.NewEncoder(w).Encode(stocks)
  }

  CreateStock(w http.ResponseWriter, r http.Request){
    var stock models.Stock
    err:= json.NewDecoder(r.Body).Decode(&stock)
    if err != nil{
      log.Fatal("Unable to decode the request body. %v",err)
    }
    insertID := response{
      ID: insertID,
      Message: "Stock created Succesfully",
    }
    json.NewEncoder(w).Encode(res)
  }

  UpdateStocks(w http.ResponseWriter, r http.Request){
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
      log.Fatalf("unable to convert the string into int. %v", err)
    }
      var stock models.Stock
      err = json.NewDecoder(r.Body).Decode(&stock)
      
      if err != nil{
        log.Fatalf("unable to decode the request body. %v",err)
      }
      updateRows := updateStock(int64(id), stock)
      msg := fmt.Sprint("Stock updated Succesfully, total row/record affected %v",updatedRows)
      res := response{
        ID: int64(id),
        Message: msg,
      }
      json.NewEncoder(w).Encode(res)
  }
  DeleteStocks(w http.ResponseWriter, r http.Request){
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])

    if err != nil{
      log.Fatalf("unable to convert the string into int. %v",err)
    }
    deletedRows := deleteStock(int64(id))
    msg := fmt.Sprintf("Stocks updated Succesfully. Total rows/record affected %v",deletedRows)
    res := response{
      ID: int64(id),
      Message: msg,
    }
    json.NewEncoder(w).Encode(res)
  }

  //Handler functions---

  func GetAllStocks()([]models.Stocks, error)
  db := createConnection()
  defer db.close()
  var stocks []models.stocks
  sqlStatement := `SELECTS * FROM stocks`
  rows, err := db.Query(sqlStatement)
  if err != nil {
    log.Fatalf("unable to excute the query. %v", err)
  }
  defer rows.Close()
  for rows.Next(){
    var stock models.Stock 
     err = rows.Scan(&stock.StockID,&stock.Name,&stock.Price,&stock.Company)
     if err != nil {
      
     }
  }
}
