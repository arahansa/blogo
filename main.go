package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	c "blogo/go/controller"
)

func main() {
	router := httprouter.New()
	router.GET("/", c.Index)
	router.GET("/board/:id", c.BoardGetId)
	log.Fatal(http.ListenAndServe(":8080", router))
}
