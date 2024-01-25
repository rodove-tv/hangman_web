package main

import (
	"fmt"
	"net/http"
	"sample-app/site/internal/handler"
)

func main() {
	port := ":5500"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/hangman", handler.Hangman)
	http.HandleFunc("/setting", handler.Setting)
	fmt.Println("(http://localhost:5500)- Server started on port", port)
	http.ListenAndServe(port, nil)

}
