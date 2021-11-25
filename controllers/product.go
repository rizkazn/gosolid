package controllers

import (
	"net/http"

	"github.com/biFebriansyah/gosolid/interfaces"
	"github.com/gorilla/mux"
)

type prods_controll struct {
	rep interfaces.ProductServices
}

func ProdsCtrl(rps interfaces.ProductServices) *prods_controll {
	return &prods_controll{rps}
}

func (pr *prods_controll) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	res, err := pr.rep.GetByID(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res.Send(w)
}
