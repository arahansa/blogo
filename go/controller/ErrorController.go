// ErrorController.go
package controller

import (
	"fmt"
	"net/http"
)

type Error404Handler struct{}
type Error405Handler struct{}
type Error500Handler struct{}

//TODO : make proper page
func (e Error404Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "custom 404")
}

//TODO : make proper page
func (e Error405Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "custom 405")
}

//TODO : make proper page
func (e Error500Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "custom 500")
}
