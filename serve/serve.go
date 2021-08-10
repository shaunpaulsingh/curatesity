package serve

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("htdocs/css/"))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))
	is := http.FileServer(http.Dir("htdocs/img/"))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", is))

	r.HandleFunc("/", HomeHandler)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "htdocs/index.gohtml")
}
