package handler

import (
  "encoding/json"
  "github.com/cicika/go-carinthia/model"
  "fmt"
  "net/http"
)

func RespondWith(w http.ResponseWriter, response model.HttpResponse) {
  if codeIs20x(response.StatusCode) {
    res, _ := json.Marshal(response.Body)
    fmt.Fprintf(w, "%s", res)
  } else {
    http.Error(w, "", response.StatusCode)
  }
}

func codeIs20x(statusCode int) bool {
  if statusCode >= 200 && statusCode < 250 {
    return true
  } else {
    return false
  }
}
