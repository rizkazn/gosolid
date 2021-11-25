package repository

import (
	"github.com/biFebriansyah/gosolid/models"
	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mocks mock.Mock
}

var data = models.Product{
	Name:  "testing",
	Price: "200",
}

func (pr *ProductRepoMock) FindByID(id int) (*models.Product, error) {
	args := pr.Mocks.Called(id)

	if args.Get(0) == nil {
		data.Id = 0
		return &data, nil
	}

	data.Id = id
	return &data, nil

}
