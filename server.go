package main

import (
  "fmt"
  "github.com/cicika/go-carinthia/api"
)

func main() {
  fmt.Printf("Startng up ...")
  
  api.StartHttpServer()  
}
