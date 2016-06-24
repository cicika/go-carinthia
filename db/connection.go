package db

import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/lib/pq"
  "github.com/cicika/go-carinthia/db/query"
)

func OpenConnection()(*sql.DB) {
  dbConn, err := sql.Open("postgres", "postgres://carinthia:23456789!@localhost:5432/carinthia?sslmode=disable")

  if err != nil {
    log.Fatal(err)
  }

  return dbConn
}

func CreateTables() {
  fmt.Printf("Creating database tables...")
  
  dbConn := OpenConnection()
  _, err := dbConn.Query(query.CreateUserTable)

  if err != nil {
    log.Fatal(err)
  }
}