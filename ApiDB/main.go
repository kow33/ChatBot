package main

import (
	"flag"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var ip string
var port string
var addr string
var loginMySql string
var passwordMySql string
var mysqlServerAddr string
var templates *template.Template

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

	templates = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/api", http.StatusOK))
	r.HandleFunc("/api", HomeHandler)

	r.HandleFunc("/api/v1/schedule/professors", BasicAuth(LogHandlerFunc(ProfessorsHandler)))
	r.HandleFunc("/api/v1/schedule/professors/{surname}", LogHandlerFunc(ProfessorGetHandler)).
		Methods("GET")
	r.HandleFunc("/api/v1/schedule/professors/{id}", BasicAuth(LogHandlerFunc(ProfessorHandler))).
		Methods("PUT", "DELETE")
	r.HandleFunc("/api/v1/schedule/info/professors", LogHandlerFunc(ProfessorsInfoHandler)).
		Methods("GET")

	r.HandleFunc("/api/v1/schedule/student_groups", BasicAuth(LogHandlerFunc(StudentsGroupsHandler)))
	r.HandleFunc("/api/v1/schedule/student_groups/{group_name}", BasicAuth(LogHandlerFunc(StudentGroupHandler))).
		Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/api/v1/schedule/info/student_groups", LogHandlerFunc(StudentsInfoHandler)).
		Methods("GET")

	r.HandleFunc("/api/v1/other_themes/jokes", BasicAuth(LogHandlerFunc(JokesHandler)))
	r.HandleFunc("/api/v1/other_themes/jokes/{theme}", LogHandlerFunc(JokeGetHandler)).
		Methods("GET")
	r.HandleFunc("/api/v1/other_themes/jokes/{id}", BasicAuth(LogHandlerFunc(JokeHandler))).
		Methods( "PUT", "DELETE")
	r.HandleFunc("/api/v1/other_themes/info/jokes", LogHandlerFunc(JokesInfoHandler)).
		Methods("GET")

	r.HandleFunc("/api/v1/add_professors", BasicAuth(LogHandlerFunc(AddProfessorHandler)))
	r.HandleFunc("/professors", LogHandlerFunc(ProfessorTemplateHandler)).
		Methods("GET")
	r.HandleFunc("/jokes", LogHandlerFunc(JokesTemplateHandler)).
		Methods("GET")

	log.Printf("Server started on %s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln("Server Error: ", err)
	}
}
