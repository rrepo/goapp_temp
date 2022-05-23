package controllers

import (
	_ "fmt"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w,"hello gopher","layout","top")
}


func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

