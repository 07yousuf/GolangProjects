package models

  type Stock struct{
    Stockid int64 `json:"StockID"`
    Name string `json:"Name"`
    Company string `json:"Company"`
    Price int64 `json:"Price"`
  }
