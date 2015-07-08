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
	router.GET("/board/page/:pageNumber", c.BoardListPage)

	router.POST("/board", c.BoardCreate)
	router.GET("/board/read/:id", c.BoardReadId)
	router.POST("/board/update/:id", c.BoardUpdate)
	router.GET("/board/delete/:id", c.BoardDelete)

	router.GET("/board/formWrite", c.BoardFormWrite)
	router.GET("/board/formUpdate/:id", c.BoardFormUpdate)
	
	router.GET("/board/createDummyData", c.CreateDummyData)
	router.GET("/board/RemoveAll", c.RemoveAll)
	// !-- End Board Route --! //

	router.ServeFiles("/static/*filepath", http.Dir("resources/static"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
