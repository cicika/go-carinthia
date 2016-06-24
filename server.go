package main

import (
  "fmt"
  "github.com/cicika/go-carinthia/api"
  "github.com/cicika/go-carinthia/db"
)

func main() {
  fmt.Printf("Startng up ...")
  
  api.StartHttpServer()  
}
