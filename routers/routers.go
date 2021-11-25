package routers

import (
	"net/http"

	"github.com/biFebriansyah/gosolid/configs/db"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// inisialisasi endpoint
	mainRutes := r.PathPrefix("/api/v1").Subrouter().StrictSlash(false)
	// mainRutes.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodOptions)

	// db init
	dbms, _ := db.New()
	ProductRoute(mainRutes, dbms)

	return r

}

func notFound(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Where yu go?"))
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello from API"))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Origin", "http://foo.com")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}
