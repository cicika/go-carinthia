package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/business"
  "github.com/cicika/go-carinthia/model"
  "github.com/cicika/go-carinthia/util"
  "encoding/json"
  "net/http"
)

var response model.HttpResponse
var extracted map[string]string

func RegisterUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  required := []string{"email", "password", "name"}

  if ValidateRequired(required, r) {
    forExtraction := []string{"email", "password", "name", "beacon_identifier"}
    extracted = ExtractParams(forExtraction, r)
    business.CreateUser(extracted)
    response = model.HttpResponse{200, ""}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
  RespondWith(w, response)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  required := []string{"email", "password"}

  if ValidateRequired(required, r) {
    forExtraction := []string{"email", "password"}
    extracted = ExtractParams(forExtraction, r)
    extracted["password"] = util.MD5(extracted["password"])
    user := business.LoginUser(extracted)
    json, _ := json.Marshal(user)
    
    response = model.HttpResponse{200, string(json[:])}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
  RespondWith(w, response)
}

func User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response = model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func PaymentMethod(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response = model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}
