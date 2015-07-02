//ArticleController
package controller

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func BoardGetId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "게시글 아이디 : , %s!\n", ps.ByName("id"))
}
