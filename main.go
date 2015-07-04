package main

import (
	c "blogo/go/controller"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Go Go Blogo!! ")

	router := httprouter.New()
	router.GET("/", c.Index)
	router.GET("/board/:id", c.BoardGetId)
	router.ServeFiles("/static/*filepath", http.Dir("resources/static"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
