package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html, err := os.Open("./static/api/index.html")
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}
	defer html.Close()

	bHtml, err := ioutil.ReadAll(html)
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bHtml)
}