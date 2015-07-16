// webmvcrouter.go
package config

import (
	control "blogo/go/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var article = control.Article{}

func GetRouter() *httprouter.Router {

	router := httprouter.New()

	// !-- Begin Board L-CRUD-Form --!  //
	router.GET("/", article.List)
	router.GET("/board/page/:pageNumber", control.BoardListPage)

	router.POST("/board", control.BoardCreate)
	router.GET("/board/read/:id", control.BoardReadId)
	router.POST("/board/update/:id", control.BoardUpdate)
	router.GET("/board/delete/:id", control.BoardDelete)

	router.GET("/board/formWrite", control.BoardFormWrite)
	router.GET("/board/formUpdate/:id", control.BoardFormUpdate)

	router.GET("/board/errorParsing", control.BoardErrorParsing)
	router.GET("/board/createDummyData", control.CreateDummyData)
	router.GET("/board/RemoveAll", control.RemoveAll)

	router.HandleMethodNotAllowed = false
	router.NotFound = new(control.Error404Handler)
	router.MethodNotAllowed = new(control.Error405Handler)

	// !-- End Board Route --! //

	router.ServeFiles("/static/*filepath", http.Dir("resources/static"))

	return router
}
