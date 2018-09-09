package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func StudentsGroupsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	switch r.Method {
	case http.MethodGet:
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		days := strings.Split(strings.ToLower(r.URL.Query().Get("days")), ",")
		query := "SELECT * FROM `student_groups`"
		if len(r.URL.Query().Get("days")) != 0 {
			query = "SELECT `id`, `group_name`, " + strings.Join(days, ", ") + " FROM `student_groups`"
		} else {
			days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		}
		var groups []Group

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
			var group Group

			resPointers := make([]interface{}, len(days)+2)
			resPointers[0] = &group.Id
			resPointers[1] = &group.GroupName
			for i := 0; i < len(days); i++ {
				resPointers[i+2] = &week[i]
			}

			err = rows.Scan(resPointers...)
			if ServerError(err, http.StatusBadGateway, w) {
				return
			}

			err = group.Week.UnmarshalServerWeek(days, week)
			if ServerError(err, http.StatusInternalServerError, w) {
				return
			}

			groups = append(groups, group)
		}

		if len(groups) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(groups)
	case http.MethodPost:
		var group Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		query := "INSERT INTO `student_groups`(`group_name`, " +
			"`monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?)"

		week, err := group.Week.GetWeekInJSON()
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		_, err = db.Query(query, group.GroupName, week[0], week[1], week[2], week[3], week[4], week[5])
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func StudentGroupHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	groupName := mux.Vars(r)["group_name"]
	switch r.Method {
	case http.MethodGet:
		days := strings.Split(strings.ToLower(r.URL.Query().Get("days")), ",")
		query := "SELECT * FROM `student_groups` WHERE `group_name` = \"" + groupName + "\""
		if len(r.URL.Query().Get("days")) != 0 {
			query = "SELECT `id`, `group_name`, " + strings.Join(days, ", ") + " FROM `student_groups` " +
				"WHERE `group_name` = \"" + groupName + "\""
		} else {
			days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		}
		var group Group

		rows, err := db.Query(query)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}
		if rows.Next() {
			week := make([][]byte, len(days))

			resPointers := make([]interface{}, len(days)+2)
			resPointers[0] = &group.Id
			resPointers[1] = &group.GroupName
			for i := 0; i < len(days); i++ {
				resPointers[i+2] = &week[i]
			}

			err = rows.Scan(resPointers...)
			if ServerError(err, http.StatusBadGateway, w) {
				return
			}

			err = group.Week.UnmarshalServerWeek(days, week)
			if ServerError(err, http.StatusInternalServerError, w) {
				return
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(group)
	case http.MethodDelete:
		query := "DELETE FROM `student_groups` WHERE `group_name` = \"" + groupName + "\""
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
		var group Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		query := "UPDATE `student_groups` " +
			"SET `group_name` = ?, `monday` = ?, `tuesday` = ?, " +
			"`wednesday` = ?, `thursday` = ?, `friday` = ?, `saturday` = ? " +
			"WHERE `group_name` = ?"

		week, err := group.Week.GetWeekInJSON()
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		_, err = db.Query(query, group.GroupName, week[0], week[1], week[2], week[3], week[4], week[5], groupName)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func StudentsInfoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := DbConn("schedule")
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	type GroupInfo struct{
		GroupName string `json:"group_name"`
	}

	var groups []GroupInfo

	query := "SELECT `group_name` FROM `student_groups`"

	rows, err := db.Query(query)
	if ServerError(err, http.StatusBadGateway, w) {
		return
	}

	for rows.Next() {
		var group GroupInfo

		err = rows.Scan(&group.GroupName)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		groups = append(groups, group)
	}

	if len(groups) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(groups)
}