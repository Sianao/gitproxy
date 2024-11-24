package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sianao/gitproxy/handler"
	"github.com/sianao/gitproxy/router"
)

func main() {
	router := router.NewRouter()
	srv := &http.Server{
		Handler:      handler.NewHandler(router),
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
