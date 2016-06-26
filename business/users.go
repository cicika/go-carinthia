package business

import (
  "log"
  "strings"
  "github.com/cicika/go-carinthia/db"
  "github.com/cicika/go-carinthia/model"
  "github.com/cicika/go-carinthia/util"
)

func CreateUser(paramMap map[string]string)(string){
  token := authToken(paramMap["Email"], paramMap["Name"])
  params := []string{paramMap["Email"],
                     util.MD5(paramMap["Password"]),
                     paramMap["Name"],
                     paramMap["BirthYear"],
                     token,
                     "passenger", "" }

  db.Execute(len(params), params, db.UserQ("CreateUser"))

  paramMap["AuthToken"] = token
  UpdateUser(paramMap)
  return token
}

func UpdateUser(paramMap map[string]string){
  params := []string{paramMap["AuthToken"]}
  db.Execute(1, params, db.UserQ("UpdateUser"))
}

func LoginUser(paramMap map[string]string)(model.User) {
  params := []string{paramMap["Email"], util.MD5(paramMap["Password"])}
  rows := db.Execute(len(params), params, db.UserQ("CheckLoginDetails"))
  var id int64
  var age int32
  var email, password, name, auth_token, utype string
  var users []model.User
  for rows.Next() {
    err := rows.Scan(&id, &email, &password, &name, &age, &auth_token, &utype)
      if err != nil {
        log.Fatal(err)
      }
      users = append(users, model.User{id, email, name, age, auth_token, utype, ""})
  }
  return users[0]
}

func AddPaymentMethod(paramMap map[string]string){
  params := []string{paramMap["UserId"], 
                     lastFour([]byte(paramMap["CardNumber"])),
                     paramMap["CardHolder"], paramMap["Expiration"],
                     paramMap["Cvv"]}
  db.Execute(len(params), params, db.UserQ("AddPaymentMethod"))

}

func authToken(email string, name string)(string) {
  return util.MD5(strings.Join([]string{email}, name))
}

func lastFour(input []byte)(string) {
  start := len(input) - 5
  end := len(input) - 1
  return string(append(input[:start], input[end:]...)[:])
}
