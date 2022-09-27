package server

import (
	"ascii/ascii"
	"net/http"
	"strconv"
	"text/template"
)

func HandleOutput(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if r.Method == http.MethodGet {
			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		rw.Header().Set("Allow: ", http.MethodPost)
		http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/ascii.html")
	if err != nil {
		ServerError(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	text := r.FormValue("data")
	banner := r.FormValue("banner")
	if text == "" || banner != "Standard" && banner != "Shadow" && banner != "Thinkertoy" {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	result, errAscii := (ascii.Ascii(text, banner))
	if errAscii == 400 {
		ServerError(rw, badRequestError, errAscii)
		return
	} else if errAscii == 500 {
		ServerError(rw, internalSrvError, http.StatusInternalServerError)
		return
	}
	if r.FormValue("download") == "download" {
		rw.Header().Add("Content-Disposition", "attachment; filename=ascii.txt")
		rw.Header().Add("Content-Type", "text/plain")
		rw.Header().Add("Content-Length", strconv.Itoa(len(result)))
		rw.Write([]byte(result))
		return
	} else if r.FormValue("generate") != "generate" {
		http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	res := struct {
		Text string
	}{
		Text: result,
	}
	err = tmpl.Execute(rw, res)
	if err != nil {
		ServerError(rw, err.Error(), http.StatusInternalServerError)
	}
}
