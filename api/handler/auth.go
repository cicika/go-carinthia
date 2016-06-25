package handler

import (
  "github.com/cicika/go-carinthia/db"
  "github.com/julienschmidt/httprouter"
  "fmt"
  "net/http"
  "strings"
)

func BasicAuth(h httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    const basicAuthPrefix string = "GoCarinthia "

    auth := r.Header.Get("Authorization")
    fmt.Printf(auth)

    if strings.HasPrefix(auth, basicAuthPrefix) {
      token := strings.SplitAfter(auth, basicAuthPrefix)[1]
      users := checkToken(token)
      
      if len(users) > 0 {
        h(w, r, append(ps, httprouter.Param{"uid", string(users[0])}))
        return
      }
    }

    w.Header().Set("WWW-Authenticate", "GoCarinthia realm=Restricted")
    http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
  }
}

func checkToken(token string) []int64 {
  return db.ExecuteWithId(1, []string{token}, db.UserQ("CheckToken"))
}
