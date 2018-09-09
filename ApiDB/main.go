package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var ip string
var port string
var addr string
var loginMySql string
var passwordMySql string
var mysqlServerAddr string

func init() {
	EnvBind()

	var dropScriptPath string
	var initScriptPath string
	var isNeedInit bool
	var isNeedDrop bool

	flag.BoolVar(&isNeedInit, "init_db", false, "True - init databases, false - otherwise")
	flag.BoolVar(&isNeedDrop, "drop_db", false, "True - drop if exists and init databases, " +
		"false - do nothing")
	flag.StringVar(&dropScriptPath, "dscript", "SqlScripts/dropDb.sql",
		"Enter path to custom drop script")
	flag.StringVar(&initScriptPath, "iscript", "SqlScripts/initDb.sql",
		"Enter path to custom init script")

	flag.Parse()
	if isNeedDrop {
		DropDb(dropScriptPath)
		InitDb(initScriptPath)
	} else if isNeedInit {
		InitDb(initScriptPath)
	}
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/api", http.StatusOK))
	r.HandleFunc("/api", HomeHandler)

	r.HandleFunc("/api/v1/schedule/professors", BasicAuth(ProfessorsHandler))
	r.HandleFunc("/api/v1/schedule/professors/{surname}", ProfessorGetHandler).Methods("GET")
	r.HandleFunc("/api/v1/schedule/professors/{id}", BasicAuth(ProfessorHandler)).
		Methods("PUT", "DELETE")
	r.HandleFunc("/api/v1/schedule/info/professors", ProfessorsInfoHandler).Methods("GET")

	r.HandleFunc("/api/v1/schedule/student_groups", BasicAuth(StudentsGroupsHandler))
	r.HandleFunc("/api/v1/schedule/student_groups/{group_name}", BasicAuth(StudentGroupHandler)).
		Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/api/v1/schedule/info/student_groups", StudentsInfoHandler).Methods("GET")

	r.HandleFunc("/api/v1/other_themes/jokes", BasicAuth(JokesHandler))
	r.HandleFunc("/api/v1/other_themes/jokes/{theme}", JokeGetHandler).Methods("GET")
	r.HandleFunc("/api/v1/other_themes/jokes/{id}", BasicAuth(JokeHandler)).Methods( "PUT", "DELETE")
	r.HandleFunc("/api/v1/other_themes/info/jokes", JokesInfoHandler).Methods("GET")

	log.Printf("Server started on %s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln("Server Error: ", err)
	}
}
