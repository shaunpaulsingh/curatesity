package serve

import (
	"curatesity/htdocs/sites"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var sitesHeld sites.Sites

var funcs = template.FuncMap{"get": sitesHeld.GetSite}

func Init() {
	sitesHeld.Init()

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("htdocs/css/"))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))
	is := http.FileServer(http.Dir("htdocs/img/"))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", is))

	r.HandleFunc("/", HomeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("index.gohtml").Funcs(funcs).ParseFiles("htdocs/index.gohtml"))
	err := t.Execute(w, sitesHeld)
	if err != nil {
		panic(err)
	}
}
