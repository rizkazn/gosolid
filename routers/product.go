package routers

import (
	"database/sql"

	"github.com/biFebriansyah/gosolid/controllers"
	"github.com/biFebriansyah/gosolid/repository"
	"github.com/biFebriansyah/gosolid/services"
	"github.com/gorilla/mux"
)

func ProductRoute(r *mux.Router, db *sql.DB) {
	repos := repository.Product(db)
	servic := services.ProdService(repos)
	ctrls := controllers.ProdsCtrl(servic)
	route := r.PathPrefix("/prods").Subrouter()

	// users endpoint
	route.HandleFunc("/{id}", ctrls.GetByID).Methods("GET")
}
