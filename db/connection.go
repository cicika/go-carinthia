package db

import (
  "database/sql"
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

func Execute(pc int, params []string, q string)(*sql.Rows){
  db := OpenConnection()
  queryMethod := query.QueryMethod(db, pc, params, q)

  rows := queryMethod(db, params, q)
  return rows
}

/* 
  returns affected row ID
*/
func ExecuteWithId(pc int, params []string, q string)([]int64){
  var ids []int64
  var id int64
  rows := Execute(pc, params, q)
  for rows.Next() {
    err := rows.Scan(&id)
    if err != nil {
      log.Fatal(err)
    }
    ids = append(ids, id)
  }

  return ids
}
