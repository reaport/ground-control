package main

import (
	"net/http"

	"github.com/reaport/ground-control/internal/controller"
	"github.com/reaport/ground-control/pkg/api"
)

func main() {
	controller := controller.New(nil)
	server, err := api.NewServer(controller)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
