package serve

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		fs := http.FileServer(http.Dir("htdocs/css/"))
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))
		is := http.FileServer(http.Dir("htdocs/img/"))
		r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", is))

		r.HandleFunc("/", HomeHandler)

		log.Fatal(listener(portStr, r))
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "htdocs/index.gohtml")
}
