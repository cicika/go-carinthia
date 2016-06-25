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

func ValidateRequired(forValidation []string, r *http.Request)(bool) {
  var reqData map[string]string
  formData := postData(r)
  json.Unmarshal([]byte(formData), &reqData)
  cnt := len(forValidation)
  for i := 0; i < cnt; i++ {
    if reqData[forValidation[i]] == "" {
      return false
    }
  }
  return true
}

func ExtractParams(forExtraction []string, r *http.Request)(map[string]string) {
  var reqData map[string]string
  formData := postData(r)
  json.Unmarshal([]byte(formData), &reqData)
  var paramMap map[string]string
  paramMap = make(map[string]string)
  cnt := len(forExtraction)
  for i := 0; i < cnt; i++ {
    paramMap[forExtraction[i]] = reqData[forExtraction[i]]
  }
  return paramMap
}

func postData(r *http.Request)(string) {
 // r.ParseForm()
  fmt.Printf("%s", r)
  reqMap := r.Form

  keys := make([]string, len(reqMap))
  i := 0
  for k := range reqMap {
    keys[i] = k
    i++
  }
  return keys[0]
}
