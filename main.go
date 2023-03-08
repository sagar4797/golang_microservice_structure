package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/sagar4797/microservice/controller"
	"github.com/sagar4797/microservice/db"
	"github.com/sagar4797/microservice/logger"
	"github.com/sagar4797/microservice/routers"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	logrus.SetFormatter(logger.Logger{
		Service:   "Service",
		Version:   os.Getenv("APP_VERSION"),
		Formatter: formatter,
	})
}

func main() {
	dbconn, err := db.Connect()
	if err != nil {
		logrus.Fatalf("Postgresql init: %s", err)
	} else {
		logrus.Infof("Postgres connected, Status: %#v", dbconn.Stats())
	}
	defer dbconn.Close()

	cf := controller.NewControllerFactory(dbconn)
	router := routers.Setup(cf)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 40 * time.Second,
	}
	logrus.Info("Server initializing..")
	if err = s.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
	return
}
