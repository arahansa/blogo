package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	var templates = template.Must(template.New("index").ParseFiles("index.html"))
	data := map[string]interface{}{"Title": "Hello World!"}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatal(err)
	}

	//templates.Execute(w, data) //error

}
