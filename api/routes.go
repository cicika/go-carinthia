package api

import (
  "log"
  "net/http"
  "github.com/cicika/go-carinthia/api/handler"
  "github.com/julienschmidt/httprouter"
)

func StartHttpServer() {
  router := httprouter.New()

  router.GET("/user/{userId}", handler.User)
  router.POST("/user/", handler.RegisterUser)
  router.POST("/user/payment/", handler.PaymentMethod) /* not quite certain about this */

  router.POST("/trip/start/", handler.TripStart)
  router.POST("/trip/segment/", handler.TripSegment)
  router.POST("/trip/end/", handler.TripEnd) /* Do we need this one? */
  router.POST("/trip/passenger/", handler.AddPassenger)
  router.DELETE("/trip/passenger/", handler.RemovePassenger)
  
  router.GET("/report/weekly/", handler.WeeklyReport)
  router.GET("/report/monthly/", handler.MonthlyReport)
  
  router.GET("/ping", handler.Pong)

  log.Fatal(http.ListenAndServe(":8080", router))
}
