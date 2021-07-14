package serve

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.Handle("/", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
