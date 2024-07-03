package main

import (
	"github.com/wellingtonchida/products-with-gin/internals/server"
)

func main() {
	server := server.New()
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
