package main

import (
	c "blogo/go/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	log.Println("Go Go Blogo!! ")
	router := httprouter.New()

	// !-- Begin Board L-CRUD-Form --! //
	router.GET("/", c.BoardList)

	router.POST("/board", c.BoardCreate)
	router.GET("/board/:id", c.BoardReadId)
	router.POST("/board/:id", c.BoardUpdate)
	router.GET("/boardDelete/:id", c.BoardDelete)

	router.GET("/boardFormWrite", c.BoardFormWrite)
	router.GET("/boardFormUpdate/:id", c.BoardFormUpdate)
	// !-- End Board Route --! //

	router.ServeFiles("/static/*filepath", http.Dir("resources/static"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
