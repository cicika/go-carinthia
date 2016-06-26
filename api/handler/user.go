package handler

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/business"
  "github.com/cicika/go-carinthia/model"
  "github.com/cicika/go-carinthia/util"
  "encoding/json"
  "net/http"
  _"net/http/httputil"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

}

func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  //fmt.Printf("%s", "                      ")
  //dump, _ := httputil.DumpRequest(r, true)

  fmt.Printf("%s", r.Body)
  var response model.HttpResponse
  var extracted map[string]string
  required := []string{"Email", "Password", "Name", "BirthYear"}

  if ValidateRequired(required, r) == true {
    forExtraction := []string{"Email", "Password", "Name", "BirthYear"}
    extracted = ExtractParams(forExtraction, r)
    authToken, userId := business.CreateUser(extracted)
    responseString := "{\"AuthToken\":" +"\"" + authToken + "\", \"UserId\":" + string(userId) + "\"}"
    response = model.HttpResponse{200, responseString}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
  RespondWith(w, response)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Printf("%s", r)
  var response model.HttpResponse
  var extracted map[string]string

  required := []string{"Email", "Password"}

  if ValidateRequired(required, r) == true {
    forExtraction := []string{"Email", "Password"}
    extracted = ExtractParams(forExtraction, r)
    extracted["Password"] = util.MD5(extracted["Password"])
    user := business.LoginUser(extracted)
    json, _ := json.Marshal(user)
    
    response = model.HttpResponse{200, string(json[:])}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
  RespondWith(w, response)
}

func User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func PaymentMethod(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  var response model.HttpResponse
  var extracted map[string]string

  required := []string{"CardNumber", "CardHolder", "Expiration", "Cvv"}

  if ValidateRequired(required, r) == true {
    forExtraction := []string{"CardNumber", "CardHolder", "Expiration", "Cvv"}
    extracted = ExtractParams(forExtraction, r)
    extracted["UserId"] = params.ByName("UserId")
    business.AddPaymentMethod(extracted)
    
    response = model.HttpResponse{200, ""}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
  RespondWith(w, response)
}
