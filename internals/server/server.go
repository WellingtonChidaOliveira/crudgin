package server

import (
	"flag"
	"net/http"
	"time"

	"github.com/wellingtonchida/products-with-gin/internals/database"
)

type Server interface{}

type server struct {
	//port
	db database.Service
}

func New() *http.Server {
	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	flag.Parse()

	db := database.New()

	svc := &server{
		db: db,
	}

	server := &http.Server{
		Addr:         *listenAddr,
		Handler:      svc.NewRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
