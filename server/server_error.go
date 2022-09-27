package server

import (
	"net/http"
	"text/template"
)

type Response struct {
	Status  int
	Message string
}

const (
	badRequestError  = "Status code: 400. Bad request."
	notFoundError    = "Status code: 404. Not Found."
	internalSrvError = "Status code: 500. Internal Server Error."
)

func ServerError(w http.ResponseWriter, msg string, status int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := &Response{
		Message: msg,
		Status:  status,
	}
	w.WriteHeader(resp.Status)
	err = tmpl.Execute(w, resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
