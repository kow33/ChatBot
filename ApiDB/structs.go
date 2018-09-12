package main

import (
	"encoding/json"
)

type Joke struct {
	Id 				int 		`json:"id"`
	Theme 			string 		`json:"theme"`
	Body 			string 		`json:"body"`
}

type Professor struct {
	Id         		int    		`json:"id"`
	Firstname  		string 		`json:"firstname"`
	Surname    		string 		`json:"surname"`
	Patronymic 		string 		`json:"patronymic"`
	Chair      		string 		`json:"chair"`
	Week       		Week   		`json:"week"`
}

type Group struct {
	Id        		int    		`json:"id"`
	GroupName 		string 		`json:"group_name"`
	Week      		Week   		`json:"week"`
}

type Week struct {
	Monday    		*Day 		`json:"monday"`
	Tuesday   		*Day 		`json:"tuesday"`
	Wednesday 		*Day 		`json:"wednesday"`
	Thursday  		*Day 		`json:"thursday"`
	Friday    		*Day 		`json:"friday"`
	Saturday  		*Day 		`json:"saturday"`
}

type Day struct {
	Lessons 		[]*Lesson 	`json:"lessons"`
}

type Lesson struct {
	Time    		string   	`json:"time"`
	Subject 		Subject 	`json:"subject"`
}

type Subject struct {
	Numerator   	string 		`json:"numerator"`
	Denominator 	string 		`json:"denominator"`
	IsDiffer    	bool   		`json:"is_differ"`
}

func (w *Week) GetWeekInJSON() ([][]byte, error) {
	week := make([][]byte, 6)
	weekArr := []*Day{w.Monday, w.Tuesday, w.Wednesday, w.Thursday, w.Friday, w.Saturday}

	for i, d := range weekArr {
		day, err := json.Marshal(d)
		if err != nil {
			return nil, err
		}
		week[i] = day
	}

	return week, nil
}

func (w *Week) UnmarshalServerWeek(days []string, week [][]byte) error {
	for ind, d := range days {
		if string(week[ind]) == "null" {
			continue
		}
		day := Day{}
		switch d {
		case "monday":
			err := json.Unmarshal(week[ind], &day)
			if err != nil {
				return err
			}
			w.Monday = &day
		case "tuesday":
			err := json.Unmarshal(week[ind], &day)
			if err != nil {
				return err
			}
			w.Tuesday = &day
		case "wednesday":
			err := json.Unmarshal(week[ind], &day)
			if err != nil {
				return err
			}
			w.Wednesday = &day
		case "thursday":
			err := json.Unmarshal(week[ind], &day)
			if err != nil {
				return err
			}
			w.Thursday = &day
		case "friday":
			err := json.Unmarshal(week[ind], &day)
			if err != nil {
				return err
			}
			w.Friday = &day
		case "saturday":
			err := json.Unmarshal(week[ind], &day)
			if err != nil {
				return err
			}
			w.Saturday = &day
		}
	}

	return nil
}
