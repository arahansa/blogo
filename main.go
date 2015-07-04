package main

import (
	c "blogo/go/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Running blogo! ")

	router := httprouter.New()
	router.GET("/", c.Index)
	router.GET("/board/:id", c.BoardGetId)
	log.Fatal(http.ListenAndServe(":8080", router))
}
