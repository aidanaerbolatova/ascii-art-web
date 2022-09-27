package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleInput)
	mux.HandleFunc("/ascii-art-web", HandleOutput)
	mux.Handle("/templates/css/", http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))))
	mux.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))
	fmt.Printf("Starting server at post: 8080\nhttp://localhost:8080/\n")
	err := http.ListenAndServe(":8080", mux)
	return err
}
