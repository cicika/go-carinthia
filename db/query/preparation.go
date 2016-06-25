package query

import (
  "database/sql"
  _ "github.com/lib/pq"
  "log"
)

func QueryMethod(db *sql.DB, pc int, params []string, st string)(func(*sql.DB, []string, string) *sql.Rows) {
  mapping := map[int]func(*sql.DB, []string, string) *sql.Rows {
    1: query1,
    2: query2,
    3: query3,
    4: query4,
    5: query5,
    6: query6,
  }

  return mapping[pc]
}

func query1(db *sql.DB, params []string, st string)(*sql.Rows) {
  rows, err := db.Query(st, params[0])
  if (err != nil) {
    log.Fatal(err)
  }
  return rows
}

func query2(db *sql.DB, params []string, st string)(*sql.Rows) {
  rows, err := db.Query(st, params[0], params[1])
  if (err != nil) {
    log.Fatal(err)
  }
  return rows 
}

func query3(db *sql.DB, params []string, st string)(*sql.Rows) {
  rows, err := db.Query(st, params[0], params[1], params[2])
  if (err != nil) {
    log.Fatal(err)
  }
  return rows
}

func query4(db *sql.DB, params []string, st string)(*sql.Rows) {
  rows, err := db.Query(st, params[0], params[1], params[2], params[3])
  if (err != nil) {
    log.Fatal(err)
  }
  return rows
}

func query5(db *sql.DB, params []string, st string)(*sql.Rows) {
  rows, err := db.Query(st, params[0], params[1], params[2], params[3], params[4])
  if (err != nil) {
    log.Fatal(err)
  }
  return rows
}

func query6(db *sql.DB, params []string, st string)(*sql.Rows) {
  rows, err := db.Query(st, params[0], params[1], params[2], params[3], params[4], params[5])
  if (err != nil) {
    log.Fatal(err)
  }
  return rows
}
