package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func ProfessorsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	switch r.Method {
	case http.MethodGet:
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		days := strings.Split(strings.ToLower(r.URL.Query().Get("days")), ",")
		query := "SELECT * FROM `professors`"
		if len(r.URL.Query().Get("days")) != 0 {
			query = "SELECT `id`, `firstname`, `surname`, `patronymic`, `chair`, " + strings.Join(days, ", ") + " FROM `professors`"
		} else {
			days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		}
		var professors []Professor

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
			week := make([][]byte, len(days))
			var professor Professor

			resPointers := make([]interface{}, len(days)+5)
			resPointers[0] = &professor.Id
			resPointers[1] = &professor.Firstname
			resPointers[2] = &professor.Surname
			resPointers[3] = &professor.Patronymic
			resPointers[4] = &professor.Chair
			for i := 0; i < len(days); i++ {
				resPointers[i+5] = &week[i]
			}

			err = rows.Scan(resPointers...)
			if ServerError(err, http.StatusBadGateway, w) {
				return
			}

			err = professor.Week.UnmarshalServerWeek(days, week)
			if ServerError(err, http.StatusInternalServerError, w) {
				return
			}

			professors = append(professors, professor)
		}

		if len(professors) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(professors)

	case http.MethodPost:
		var professor Professor
		err := json.NewDecoder(r.Body).Decode(&professor)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		query := "INSERT INTO `professors`(`firstname`, `surname`, `patronymic`, `chair`, " +
			"`monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

		week, err := professor.Week.GetWeekInJSON()
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		_, err = db.Query(query, professor.Firstname, professor.Surname, professor.Patronymic, professor.Chair,
			week[0], week[1], week[2], week[3], week[4], week[5])
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func ProfessorGetHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	surname := mux.Vars(r)["surname"]
	switch r.Method {
	case http.MethodGet:
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		days := strings.Split(strings.ToLower(r.URL.Query().Get("days")), ",")
		query := "SELECT * FROM `professors` WHERE `surname` = \"" + surname + "\""
		if len(r.URL.Query().Get("days")) != 0 {
			query = "SELECT `id`, `firstname`, `surname`, `patronymic`, `chair`, " + strings.Join(days, ", ") + " FROM `professors` " +
				"WHERE `surname` = \"" + surname + "\""
		} else {
			days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		}
		var professors []Professor

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
			week := make([][]byte, len(days))
			var professor Professor

			resPointers := make([]interface{}, len(days)+5)
			resPointers[0] = &professor.Id
			resPointers[1] = &professor.Firstname
			resPointers[2] = &professor.Surname
			resPointers[3] = &professor.Patronymic
			resPointers[4] = &professor.Chair
			for i := 0; i < len(days); i++ {
				resPointers[i+5] = &week[i]
			}

			err = rows.Scan(resPointers...)
			if ServerError(err, http.StatusBadGateway, w) {
				return
			}

			err = professor.Week.UnmarshalServerWeek(days, week)
			if ServerError(err, http.StatusInternalServerError, w) {
				return
			}

			professors = append(professors, professor)
		}

		if len(professors) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(professors)
	}
}

func ProfessorHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}

	switch r.Method {
	case http.MethodDelete:
		query := "DELETE FROM `professors` WHERE `id` = " + strconv.Itoa(id) + ""
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
		var professor Professor
		err := json.NewDecoder(r.Body).Decode(&professor)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		query := "UPDATE `professors` " +
			"SET `firstname` = ?, `surname` = ?, `patronymic` = ?, `chair` = ?, " +
			"`monday` = ?, `tuesday` = ?, " +
			"`wednesday` = ?, `thursday` = ?, `friday` = ?, `saturday` = ? " +
			"WHERE `id` = ?"

		week, err := professor.Week.GetWeekInJSON()
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		res, err := db.Exec(query, professor.Firstname, professor.Surname, professor.Patronymic, professor.Chair,
			week[0], week[1], week[2], week[3], week[4], week[5], id)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}
		count, err := res.RowsAffected()
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}
		if count == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ProfessorsInfoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	type ProfessorInfo struct{
		Firstname string `json:"firstname"`
		Surname string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Chair string `json:"chair"`
	}

	var professors []ProfessorInfo

	query := "SELECT `firstname`, `surname`, `patronymic`, `chair` FROM `professors`"

	rows, err := db.Query(query)
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	for rows.Next() {
		var professor ProfessorInfo

		err = rows.Scan(&professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		professors = append(professors, professor)
	}

	if len(professors) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(professors)
}

func ProfessorTemplateHandler(w http.ResponseWriter, r *http.Request) {
	URL := "http://"
	if ip == "" {
		URL += "localhost:" + port
	} else {
		URL += addr
	}
	URL += "/api/v1/schedule/professors"

	req, err := http.Get(URL)
	if ServerError(err, req.StatusCode, w) {
		return
	}
	var professors []Professor

	err = json.NewDecoder(req.Body).Decode(&professors)
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}

	err = templates.ExecuteTemplate(w, "professors.gohtml", struct {
		Professors []Professor
		Times []string
	}{
		professors,
		professors[0].Week.GetTimesFromWeek(),
	})
	if ServerError(err, http.StatusInternalServerError, w) {
		return
	}
}