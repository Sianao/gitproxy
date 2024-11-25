package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sianao/gitproxy/handler"
	"github.com/sianao/gitproxy/router"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	// r := gin.Default()
	router := router.NewRouter(l)
	srv := &http.Server{
		Handler:      handler.NewHandler(router, l),
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
