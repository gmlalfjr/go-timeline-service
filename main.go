package main

import (
	"github.com/gmlalfjr/timeline-service/app"
)

func main() {
	r := app.RunServer()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
