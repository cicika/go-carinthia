package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  "net/http"
)

func User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func PaymentMethod(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}
