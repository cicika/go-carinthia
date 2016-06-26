package handler

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  _"github.com/cicika/go-carinthia/util"
  "encoding/json"
  "net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Printf("%s", r)
  responseString := "{\"AuthToken\":\"auth_token\", \"UserId\": 2}"
  response := model.HttpResponse{200, responseString}
  RespondWith(w, response)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  user := model.User{1, "jonsnow1986@gmail.com", "Jon Snow", 1986, "auth_token", "passenger", ""}
  json, _ := json.Marshal(user)
    
  response := model.HttpResponse{200, string(json[:])}  
  RespondWith(w, response)
}

func User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func PaymentMethod(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  response := model.HttpResponse{200, ""}  
  RespondWith(w, response)
}
