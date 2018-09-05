package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func ProfessorsHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConn("schedule")
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

			switch len(days) {
			case 1:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair,
					&week[0])
			case 2:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair,
					&week[0], &week[1])
			case 3:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2])
			case 4:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2], &week[3])
			case 5:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2], &week[3], &week[4])
			case 6:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname, &professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2], &week[3], &week[4], &week[5])
			}

			for ind, day := range days {
				switch day {
				case "monday":
					err = json.Unmarshal(week[ind], &professor.Week.Monday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "tuesday":
					err = json.Unmarshal(week[ind], &professor.Week.Tuesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "wednesday":
					err = json.Unmarshal(week[ind], &professor.Week.Wednesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "thursday":
					err = json.Unmarshal(week[ind], &professor.Week.Thursday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "friday":
					err = json.Unmarshal(week[ind], &professor.Week.Friday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "saturday":
					err = json.Unmarshal(week[ind], &professor.Week.Saturday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				}
			}

			professor.Week.SetIfEmpty()

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

		week := professor.Week.GetWeekInJSON()

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
	db := DbConn("schedule")
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
			switch len(days) {
			case 1:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname,
					&professor.Patronymic, &professor.Chair,
					&week[0])
			case 2:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname,
					&professor.Patronymic, &professor.Chair,
					&week[0], &week[1])
			case 3:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname,
					&professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2])
			case 4:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname,
					&professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2], &week[3])
			case 5:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname,
					&professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2], &week[3], &week[4])
			case 6:
				rows.Scan(&professor.Id, &professor.Firstname, &professor.Surname,
					&professor.Patronymic, &professor.Chair,
					&week[0], &week[1], &week[2], &week[3], &week[4], &week[5])
			}

			for ind, day := range days {
				switch day {
				case "monday":
					err = json.Unmarshal(week[ind], &professor.Week.Monday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "tuesday":
					err = json.Unmarshal(week[ind], &professor.Week.Tuesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "wednesday":
					err = json.Unmarshal(week[ind], &professor.Week.Wednesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "thursday":
					err = json.Unmarshal(week[ind], &professor.Week.Thursday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "friday":
					err = json.Unmarshal(week[ind], &professor.Week.Friday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "saturday":
					err = json.Unmarshal(week[ind], &professor.Week.Saturday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				}
			}

			professor.Week.SetIfEmpty()

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
	db := DbConn("schedule")
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

		week := professor.Week.GetWeekInJSON()
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

func StudentsGroupsHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConn("schedule")
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
			switch len(days) {
			case 1:
				rows.Scan(&group.Id, &group.GroupName, &week[0])
			case 2:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1])
			case 3:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2])
			case 4:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2], &week[3])
			case 5:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2], &week[3], &week[4])
			case 6:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2], &week[3], &week[4], &week[5])
			}

			for ind, day := range days {
				switch day {
				case "monday":
					err = json.Unmarshal(week[ind], &group.Week.Monday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "tuesday":
					err = json.Unmarshal(week[ind], &group.Week.Tuesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "wednesday":
					err = json.Unmarshal(week[ind], &group.Week.Wednesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "thursday":
					err = json.Unmarshal(week[ind], &group.Week.Thursday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "friday":
					err = json.Unmarshal(week[ind], &group.Week.Friday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "saturday":
					err = json.Unmarshal(week[ind], &group.Week.Saturday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				}
			}

			group.Week.SetIfEmpty()

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

		week := group.Week.GetWeekInJSON()
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
	db := DbConn("schedule")
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
			switch len(days) {
			case 1:
				rows.Scan(&group.Id, &group.GroupName, &week[0])
			case 2:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1])
			case 3:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2])
			case 4:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2], &week[3])
			case 5:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2], &week[3], &week[4])
			case 6:
				rows.Scan(&group.Id, &group.GroupName, &week[0], &week[1], &week[2], &week[3], &week[4], &week[5])
			}

			for ind, day := range days {
				switch day {
				case "monday":
					err = json.Unmarshal(week[ind], &group.Week.Monday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "tuesday":
					err = json.Unmarshal(week[ind], &group.Week.Tuesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "wednesday":
					err = json.Unmarshal(week[ind], &group.Week.Wednesday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "thursday":
					err = json.Unmarshal(week[ind], &group.Week.Thursday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "friday":
					err = json.Unmarshal(week[ind], &group.Week.Friday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				case "saturday":
					err = json.Unmarshal(week[ind], &group.Week.Saturday)
					if ServerError(err, http.StatusInternalServerError, w) {
						return
					}
				}
			}

			group.Week.SetIfEmpty()
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

		week := group.Week.GetWeekInJSON()
		_, err = db.Query(query, group.GroupName, week[0], week[1], week[2], week[3], week[4], week[5], groupName)
		if ServerError(err, http.StatusBadGateway, w) {
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}