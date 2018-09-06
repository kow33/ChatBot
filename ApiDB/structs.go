package main

import (
	"encoding/json"
	"net/http"
)

type Professor struct {
	Id         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Chair      string `json:"chair"`
	Week       Week   `json:"week"`
}

type Group struct {
	Id        int    `json:"id"`
	GroupName string `json:"group_name"`
	Week      Week   `json:"week"`
}

type Week struct {
	Monday    *Day `json:"monday"`
	Tuesday   *Day `json:"tuesday"`
	Wednesday *Day `json:"wednesday"`
	Thursday  *Day `json:"thursday"`
	Friday    *Day `json:"friday"`
	Saturday  *Day `json:"saturday"`
}

type Day struct {
	Lessons []*Lesson `json:"lessons"`
	IsEmpty bool      `json:"is_empty"`
}

type Lesson struct {
	Time    string   `json:"time"`
	Subject *Subject `json:"subject"`
}

type Subject struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
	IsDiffer    bool   `json:"is_differ"`
}

func (w *Week) GetWeekInJSON() [][]byte {
	week := make([][]byte, 6)

	day, err := json.Marshal(w.Monday)
	if err != nil {
		panic(err.Error())
	}
	week[0] = day

	day, err = json.Marshal(w.Tuesday)
	if err != nil {
		panic(err.Error())
	}
	week[1] = day

	day, err = json.Marshal(w.Wednesday)
	if err != nil {
		panic(err.Error())
	}
	week[2] = day

	day, err = json.Marshal(w.Thursday)
	if err != nil {
		panic(err.Error())
	}
	week[3] = day

	day, err = json.Marshal(w.Friday)
	if err != nil {
		panic(err.Error())
	}
	week[4] = day

	day, err = json.Marshal(w.Saturday)
	if err != nil {
		panic(err.Error())
	}
	week[5] = day

	return week
}

func (w *Week) UnmarshalServerWeek(days []string, week [][]byte, writter http.ResponseWriter) {
	for ind, d := range days {
		day := Day{}
		switch d {
		case "monday":

			err := json.Unmarshal(week[ind], &day)
			if ServerError(err, http.StatusInternalServerError, writter) {
				return
			}
			w.Monday = &day
		case "tuesday":
			err := json.Unmarshal(week[ind], &day)
			if ServerError(err, http.StatusInternalServerError, writter) {
				return
			}
			w.Tuesday = &day
		case "wednesday":
			err := json.Unmarshal(week[ind], &day)
			if ServerError(err, http.StatusInternalServerError, writter) {
				return
			}
			w.Wednesday = &day
		case "thursday":
			err := json.Unmarshal(week[ind], &day)
			if ServerError(err, http.StatusInternalServerError, writter) {
				return
			}
			w.Thursday = &day
		case "friday":
			err := json.Unmarshal(week[ind], &day)
			if ServerError(err, http.StatusInternalServerError, writter) {
				return
			}
			w.Friday = &day
		case "saturday":
			err := json.Unmarshal(week[ind], &day)
			if ServerError(err, http.StatusInternalServerError, writter) {
				return
			}
			w.Saturday = &day
		}
	}
}
