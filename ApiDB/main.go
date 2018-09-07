package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var ip string
var port string
var addr string
var loginMySql string
var passwordMySql string
var mysqlServerAddr string

func init() {
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

	InitDb("SqlScripts/initDb.sql")
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/api", http.StatusOK))
	r.HandleFunc("/api", HomeHandler)

	r.HandleFunc("/api/v1/schedule/professors", BasicAuth(ProfessorsHandler))
	r.HandleFunc("/api/v1/schedule/professors/{surname}", BasicAuth(ProfessorGetHandler)).Methods("GET")
	r.HandleFunc("/api/v1/schedule/professors/{id}", BasicAuth(ProfessorHandler)).
		Methods("PUT", "DELETE")

	r.HandleFunc("/api/v1/schedule/student_groups", BasicAuth(StudentsGroupsHandler))
	r.HandleFunc("/api/v1/schedule/student_groups/{group_name}", BasicAuth(StudentGroupHandler)).Methods("GET", "PUT", "DELETE")

	r.HandleFunc("/api/v1/other_themes/jokes", BasicAuth(JokesHandler))
	r.HandleFunc("/api/v1/other_themes/jokes/{theme}", BasicAuth(JokeGetHandler)).Methods("GET")
	r.HandleFunc("/api/v1/other_themes/jokes/{id}", BasicAuth(JokeHandler)).Methods( "PUT", "DELETE")

	log.Printf("Server started on %s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln("Server Error: ", err)
	}
}
