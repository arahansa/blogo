//ArticleController
package controller

import (
	"blogo/go/domain"
	"blogo/go/repository"
	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
	ht "html/template"
	"log"
	"net/http"
	"strconv"
)

var repoArticle = repository.Article{}

const basicLocation string = "resources/templates/board/"

type Article struct{}

func (a *Article) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	articles, pageInfo := repository.GetArticleListForPage(1)
	data := map[string]interface{}{
		"Title":    "Hello World!",
		"articles": articles,
		"pageInfo": pageInfo,
	}
	makeTemplateExcute("list", data, w)
}

func BoardListPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestPage, err := strconv.Atoi(ps.ByName("pageNumber"))
	checkIsNumber(err, w, r)

	articles, pageInfo := repository.GetArticleListForPage(requestPage)
	data := map[string]interface{}{
		"Title":    "Hello World!",
		"articles": articles,
		"pageInfo": pageInfo,
	}
	makeTemplateExcute("list", data, w)

}

func BoardCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	article := new(domain.Article)
	decoder := schema.NewDecoder()
	decoder.Decode(article, r.PostForm)
	log.Println(article)
	_, err := repository.Save(article)
	checkErr(err)
	http.Redirect(w, r, "/", http.StatusFound)
}

func BoardReadId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	checkIsNumber(err, w, r)
	data := map[string]interface{}{"article": repository.GetOneArticle(id)}
	makeTemplateExcute("read", data, w)
}

func BoardUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	article := new(domain.Article)
	decoder := schema.NewDecoder() // hm... I consider to make this variable global variable
	decoder.Decode(article, r.PostForm)
	log.Println("Bofore Updated Article :", article)
	id, err := strconv.Atoi(ps.ByName("id"))
	checkIsNumber(err, w, r)
	article.Id = id

	repository.Update(article)

	data := map[string]interface{}{"article": repository.GetOneArticle(id)}
	makeTemplateExcute("read", data, w)
}

func BoardDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	checkIsNumber(err, w, r)
	repository.DeleteArticle(id)
	http.Redirect(w, r, "/", http.StatusFound)
}

func BoardFormWrite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	makeTemplateExcute("formWrite", nil, w)
}

func BoardFormUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	checkIsNumber(err, w, r)
	data := map[string]interface{}{"article": repository.GetOneArticle(id)}
	makeTemplateExcute("formUpdate", data, w)
}

//Error page
func BoardErrorParsing(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	makeTemplateExcute("error/parsingnumber", nil, w)
}

func makeTemplateExcute(file string, data map[string]interface{}, w http.ResponseWriter) {
	t := ht.Must(ht.ParseFiles(basicLocation + file + ".html"))
	t = ht.Must(t.ParseGlob(basicLocation + "*.tmpl"))
	err := t.Execute(w, data)
	checkErr(err)
}

func checkIsNumber(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		log.Println("error message :", err)
		http.Redirect(w, r, "/board/errorParsing", http.StatusFound)
	}
}
func ErrorServerInternalErrorPage(err error, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/board/error", http.StatusFound)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// TODO : delete these methods
func CreateDummyData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	repoArticle.CreateDummyData()
	http.Redirect(w, r, "/", http.StatusFound)
}

// TODO : delete these methods
func RemoveAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	repository.RemoveAll()
	repository.SetPrimaryKey()
	http.Redirect(w, r, "/", http.StatusFound)
}
