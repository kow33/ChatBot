package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

var templ = template.Must(template.ParseFiles("templates/home.gohtml"))

func DocsHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := ioutil.ReadFile("Docs/docs.json")
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}

	var data data
	err = json.Unmarshal(jsonFile, &data)
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}

	for i := 0; i < len(data.Examples); i++ {
		err = data.Examples[i].GetJson()
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
	}

	err = templ.Execute(w, data)
	if ServerError(err, http.StatusNotFound, w) {
		return
	}
}