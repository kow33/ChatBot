package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func JokesHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConn("other_themes")
	switch r.Method {
	case http.MethodGet:
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		query := "SELECT * FROM `jokes`"
		var jokes []Joke

		if len(limit) != 0 {
			query += " LIMIT " + limit
			if len(offset) != 0 {
				query += " OFFSET " + offset
			}
		}

		rows, err := db.Query(query)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		for rows.Next() {
			var joke Joke

			err = rows.Scan(&joke.Id, &joke.Theme, &joke.Body)
			if ServerError(err, http.StatusBadGateway, w) {
				return
			}

			jokes = append(jokes, joke)
		}

		if len(jokes) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(jokes)
	case http.MethodPost:
		var joke Joke
		err := json.NewDecoder(r.Body).Decode(&joke)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		query := "INSERT INTO `jokes`(`theme`, `body`) VALUES (?, ?)"

		_, err = db.Query(query, joke.Theme, joke.Body)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func JokeGetHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConn("other_themes")
	theme := mux.Vars(r)["theme"]
	switch r.Method {
	case http.MethodGet:
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		query := "SELECT * FROM `jokes` WHERE `theme` = \"" + theme + "\""
		var jokes []Joke

		if len(limit) != 0 {
			query += " LIMIT " + limit
			if len(offset) != 0 {
				query += " OFFSET " + offset
			}
		}

		rows, err := db.Query(query)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		for rows.Next() {
			var joke Joke

			err = rows.Scan(&joke.Id, &joke.Theme, &joke.Body)
			if ServerError(err, http.StatusBadGateway, w) {
				return
			}

			jokes = append(jokes, joke)
		}

		if len(jokes) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(jokes)
	}
}

func JokeHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConn("other_themes")
	id := mux.Vars(r)["id"]
	switch r.Method {
	case http.MethodDelete:
		query := "DELETE FROM `jokes` WHERE `id` = " + id + ""
		res, err := db.Exec(query)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		count, err := res.RowsAffected()
		if ServerError(err, http.StatusNotFound, w) {
			return
		}
		if count == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	case http.MethodPut:
		var joke Joke
		err := json.NewDecoder(r.Body).Decode(&joke)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		query := "UPDATE `jokes` SET `theme` = ?, `body` = ? WHERE `id` = ?"

		_, err = db.Query(query, joke.Theme, joke.Body, id)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func JokesInfoHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConn("other_themes")

	type JokesInfo struct{
		Theme string `json:"theme"`
	}

	var jokes []JokesInfo

	query := "SELECT DISTINCT(`theme`) as `theme` FROM `jokes`"

	rows, err := db.Query(query)
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	for rows.Next() {
		var joke JokesInfo

		err = rows.Scan(&joke.Theme)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		jokes = append(jokes, joke)
	}

	if len(jokes) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}