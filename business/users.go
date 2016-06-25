package business

import (
  "log"
  "strings"
  "github.com/cicika/go-carinthia/db"
  "github.com/cicika/go-carinthia/model"
  "github.com/cicika/go-carinthia/util"
)

func CreateUser(paramMap map[string]string) {
  params := []string{paramMap["email"],
                       util.MD5(paramMap["password"]),
                       paramMap["name"], 
                       authToken(paramMap["email"], paramMap["name"]),
                       "passenger", "" }

  db.Execute(len(params), params, db.UserQ("CreateUser"))
}

func UpdateUser(paramMap map[string]string){
  params := []string{paramMap["auth_token"]}
  db.Execute(1, params, db.UserQ("UpdateUser"))
}

func LoginUser(paramMap map[string]string)(model.User) {
  params := []string{paramMap["email"], util.MD5(paramMap["password"])}
  rows := db.Execute(len(params), params, db.UserQ("CheckLoginDetails"))
  var id int64
  var email, name, auth_token, utype string
  var users []model.User
  for rows.Next() {
    err := rows.Scan(&id, &email, &name, &auth_token, &utype)
      if err != nil {
        log.Fatal(err)
      }
      users = append(users, model.User{id, email, name, auth_token, utype, ""})
  }
  return users[0]
}

func authToken(email string, name string)(string) {
  return util.MD5(strings.Join([]string{email}, name))
}