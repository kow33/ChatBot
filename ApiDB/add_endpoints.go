package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ProfessorTemp struct {
	Days []DayTemp
}

type DayTemp struct {
	Name string
	Lessons []LessonTemp
}

type LessonTemp struct {
	Number int
	Time string
}

func transformBool(str string) bool {
	if str == "on" {
		return false
	}
	return true
}

func AddProfessorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		lessonData := []LessonTemp {
			{1,"8:30 - 10:05"},
			{2, "10:15 - 11:50"},
			{3, "12:00 - 13:35"},
			{4, "13:50 - 15:25"},
			{5, "15:40 - 17:15"},
			{6, "17:25 - 19:00"},
			{7, "19:10 - 20:45"},
		}
		data := ProfessorTemp{
			[]DayTemp{
				{"Monday", lessonData},
				{"Tuesday", lessonData},
				{"Wednesday", lessonData},
				{"Thursday", lessonData},
				{"Friday", lessonData},
				{"Saturday", lessonData},
			},
		}

		err := templates.ExecuteTemplate(w, "professor.gohtml", data)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
	case http.MethodPost:
		r.ParseForm()

		var professor Professor

		professor.Firstname = r.FormValue("firstname")
		professor.Surname = r.FormValue("surname")
		professor.Patronymic = r.FormValue("patronymic")
		professor.Chair = r.FormValue("chair")

		days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

		for _, dayStr := range days {
			if transformBool(r.FormValue(dayStr + "_is_empty")) {
				var lessons []Lesson
				for i := 0; i < 7; i++ {
					lesson := Lesson{}
					lesson.Time = r.FormValue(dayStr + "_lesson_" + strconv.Itoa(i+1) + "_time")
					lesson.Subject.Numerator =
						r.FormValue(dayStr + "_lesson_" + strconv.Itoa(i+1) + "_numerator")
					lesson.Subject.Denominator =
						r.FormValue(dayStr + "_lesson_" + strconv.Itoa(i+1) + "_denominator")
					lesson.Subject.IsDiffer = lesson.Subject.Numerator != lesson.Subject.Denominator

					lessons = append(lessons, lesson)
				}

				switch dayStr {
				case "Monday":
					professor.Week.Monday = &Day{lessons}
				case "Tuesday":
					professor.Week.Tuesday = &Day{lessons}
				case "Wednesday":
					professor.Week.Wednesday = &Day{lessons}
				case "Thursday":
					professor.Week.Thursday = &Day{lessons}
				case "Friday":
					professor.Week.Friday = &Day{lessons}
				case "Saturday":
					professor.Week.Saturday = &Day{lessons}
				}
			}
		}

		jsonData, err := json.Marshal(professor)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
		URL := "http://"
		if ip == "" {
			URL += "localhost:" + port
		} else {
			URL += addr
		}

		URL += "/api/v1/schedule/professors"

		req, err := http.NewRequest("POST", URL, bytes.NewReader(jsonData))
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
		req.Header.Set("content-type", "application/json")
		req.SetBasicAuth("bmstuAdmin", "bmstuPassword")

		resp, err := (&http.Client{}).Do(req)
		if ServerError(err, http.StatusBadRequest, w) {
			return
		}

		bodyText, err := ioutil.ReadAll(resp.Body)

		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Header().Set("content-type", resp.Header.Get("content-type"))
		w.Write(bodyText)
	}
}

func AddJokeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		form, err := ioutil.ReadFile("static/api/addJoke.html")
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(form)
	case http.MethodPost:
		r.ParseForm()

		var joke Joke
		joke.Theme = r.FormValue("theme")
		joke.Body = r.FormValue("body")

		jsonData, err := json.Marshal(joke)
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
		URL := "http://"
		if ip == "" {
			URL += "localhost:" + port
		} else {
			URL += addr
		}

		URL += "/api/v1/other_themes/jokes"

		req, err := http.NewRequest("POST", URL, bytes.NewReader(jsonData))
		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}
		req.Header.Set("content-type", "application/json")
		req.SetBasicAuth("bmstuAdmin", "bmstuPassword")

		resp, err := (&http.Client{}).Do(req)
		if ServerError(err, http.StatusBadRequest, w) {
			return
		}

		bodyText, err := ioutil.ReadAll(resp.Body)

		if ServerError(err, http.StatusInternalServerError, w) {
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Header().Set("content-type", resp.Header.Get("content-type"))
		w.Write(bodyText)
	}
}