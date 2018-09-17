package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type data struct {
	Performers []performer `json:"performers"`
	Seniors    []senior    `json:"seniors"`
	MainFuncs  []string    `json:"main_funcs"`
	OtherFuncs []string    `json:"other_funcs"`
	Docs       []doc       `json:"docs"`
	Structures []structure `json:"structures"`
	Examples   []example   `json:"examples"`
}

type performer struct {
	Name string `json:"name"`
}

type senior struct {
	Name string `json:"name"`
}

type doc struct {
	Path    string   `json:"path"`
	Methods []method `json:"methods"`
	Params  []param  `json:"params"`
}

type method struct {
	Type string `json:"type"`
	Desc string `json:"desc"`
}

type param struct {
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	Comments []string `json:"comments"`
	Example  string   `json:"example"`
}

type structure struct {
	Header   string   `json:"header"`
	Comments []string `json:"comments"`
	Json     []string `json:"json"`
}

type example struct {
	Header string `json:"header"`
	From   string `json:"from"`
	Json   string
}

func (e *example) GetJson() error {
	URL := "http://"
	if ip == "" {
		URL += "localhost:" + port
	} else {
		URL += addr
	}
	URL += e.From

	req, err := http.Get(URL)
	if err != nil {
		return err
	}

	rbytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var result bytes.Buffer
	json.Indent(&result, rbytes, "", "  ")

	e.Json = result.String()

	return nil
}