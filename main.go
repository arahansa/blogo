package main

import (
	"blogo/go/config"
	"log"
	"net/http"
)

func main() {
	log.Println("Go Go Blogo!! ")
	http.ListenAndServe(":3000", config.GetRouter())
}
