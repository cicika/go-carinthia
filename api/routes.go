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
  router.GET("/user/:userId/:beaconIdentifier", handler.PassengerCheck)
  
  router.POST("/user/login/", handler.Login)
  router.PUT("/user/", handler.RegisterUser)

  router.GET("/user/", handler.User)
  router.PUT("/user/payment/method/", handler.PaymentMethod) /* not quite certain about this */

  router.POST("/trip/start/", handler.TripStart)
  router.POST("/trip/segment/", handler.TripSegment)
  router.POST("/trip/end/", handler.TripEnd) /* Do we need this one? */
  router.PUT("/trip/passenger/", handler.AddPassenger)
  router.DELETE("/trip/passenger/", handler.RemovePassenger)
  
  router.GET("/report/weekly/", handler.WeeklyReport)
  router.GET("/report/monthly/", handler.MonthlyReport)
  router.GET("/report/current/", handler.CurrentSpendings)
  
  log.Fatal(http.ListenAndServe(":8080", router))
}
