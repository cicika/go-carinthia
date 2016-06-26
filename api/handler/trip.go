package handler

import (
  _ "fmt"
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  _ "github.com/cicika/go-carinthia/business"
  "net/http"
)

func TripStart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  response := model.HttpResponse{200, ""}
  RespondWith(w, response)
}

func TripSegment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  response := model.HttpResponse{400, "BadRequest"}

  RespondWith(w, response)
}

func TripEnd(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  response := model.HttpResponse{200, ""}
  
  RespondWith(w, response)
}

func AddPassenger(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func RemovePassenger(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}
