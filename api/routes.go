package api

import (
  "log"
  "net/http"
  "github.com/cicika/go-carinthia/api/handler"
  "github.com/julienschmidt/httprouter"
)

func StartHttpServer() {
  router := httprouter.New()

  router.GET("/ping", handler.Pong)
  
  router.POST("/login/", handler.Login)
  router.PUT("/user/", handler.RegisterUser)

  router.GET("/user/", handler.BasicAuth(handler.User))
  router.PUT("/user/payment/method", handler.BasicAuth(handler.PaymentMethod)) /* not quite certain about this */

  router.PUT("/trip/start/", handler.BasicAuth(handler.TripStart))
  router.PUT("/trip/segment/", handler.BasicAuth(handler.TripSegment))
  router.PUT("/trip/end/", handler.BasicAuth(handler.TripEnd)) /* Do we need this one? */
  router.PUT("/trip/passenger/", handler.BasicAuth(handler.AddPassenger))
  router.DELETE("/trip/passenger/", handler.BasicAuth(handler.RemovePassenger))
  
  router.GET("/report/weekly/", handler.BasicAuth(handler.WeeklyReport))
  router.GET("/report/monthly/", handler.BasicAuth(handler.MonthlyReport))
  
  log.Fatal(http.ListenAndServe(":8080", router))
}
