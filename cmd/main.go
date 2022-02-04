package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	/*
		r.POST("/task", Add)
		r.DELETE("/task", Delete)
		r.GET("/tasks", List)
		r.GET("/task/:id", Mark)
	*/

	log.Println("start serving :8080")
	http.ListenAndServe(":8080", r)
}
