package server

import (
	"net/http"
	"text/template"
)

func HandleInput(rw http.ResponseWriter, r *http.Request) {
	var (
		tmpl *template.Template
		err  error
	)

	if r.URL.Path != "/" {
		ServerError(rw, notFoundError, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		rw.Header().Set("Allow: ", http.MethodGet)
		http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if tmpl, err = template.ParseFiles("templates/ascii.html"); err != nil {
		ServerError(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(rw, nil); err != nil {
		ServerError(rw, internalSrvError, http.StatusInternalServerError)
		return
	}
}
