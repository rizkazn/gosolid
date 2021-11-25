package services

import (
	"strconv"

	"github.com/biFebriansyah/gosolid/helpers"
	"github.com/biFebriansyah/gosolid/interfaces"
)

type prod_service struct {
	rep interfaces.ProductRepos
}

func ProdService(reps interfaces.ProductRepos) *prod_service {
	return &prod_service{reps}
}

func (pr *prod_service) GetByID(id string) (*helpers.Response, error) {
	theId, _ := strconv.Atoi(id)
	data, err := pr.rep.FindByID(theId)

	if err != nil {
		return nil, err
	}

	respon := helpers.Response{
		Status:  "200",
		IsError: false,
		Data:    data,
	}

	return &respon, nil
}
