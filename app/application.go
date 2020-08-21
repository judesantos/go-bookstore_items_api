package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/judesantos/go-bookstore_items_api/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	// setup db client

	elasticsearch.Init()

	// setup routes
	mapUrls()

	// setup rest service

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("\nItems service, listening on port 8081...")

	// start rest service

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
