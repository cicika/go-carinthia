package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  "net/http"
)

func Pong(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{200, "Pong"}
  RespondWith(w, response)
}
