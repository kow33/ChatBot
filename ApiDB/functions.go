package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
)

func ServerError(err error, statusCode int, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Println(err.Error())
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

	return true
}

func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			h.ServeHTTP(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			ServerError(errors.New("no authorized"), http.StatusUnauthorized, w)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			ServerError(err, http.StatusUnauthorized, w)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			ServerError(errors.New("no authorized"), http.StatusUnauthorized, w)

			return
		}

		if pair[0] != "bmstuAdmin" || pair[1] != "bmstuPassword" {
			ServerError(errors.New("no authorized"), http.StatusUnauthorized, w)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func LogHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n" +
			"Url: %s\n" +
			"Method: %s\n\n",
			r.URL, r.Method)
		h.ServeHTTP(w, r)
	}
}

func EnvBind() {
	ip = os.Getenv("IP_SERVER")
	port = os.Getenv("PORT_SERVER")
	loginMySql = os.Getenv("MYSQL_LOGIN")
	passwordMySql = os.Getenv("MYSQL_PASSWORD")

	if len(ip) == 0 {
		ip = ""
	}
	if len(port) == 0 {
		port = "8080"
	}
	if len(loginMySql) == 0 {
		loginMySql = "root"
	}
	if len(passwordMySql) == 0 {
		passwordMySql = ""
	}

	addr = ip + ":" + port

	mysqlServerAddr = loginMySql + ":" + passwordMySql + "@/"
}

func LogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}