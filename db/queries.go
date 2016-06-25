package db

import "github.com/cicika/go-carinthia/db/query"

func UserQ(q string)string{
  mapping := map[string]string {
    "CreateUser": query.CreateUser,
    "UpdateAuthToken": query.UpdateAuthToken,
    "CheckLoginDetails" : query.CheckLoginDetails,
  }
  return mapping[q]
}
