package services

import (
	"github.com/biFebriansyah/gosolid/helpers"
	"github.com/stretchr/testify/mock"
)

type ProductServMock struct {
	Mocks mock.Mock
}

var data = helpers.Response{
	Status: "200",
}

func (pr *ProductServMock) GetByID(id string) (*helpers.Response, error) {
	args := pr.Mocks.Called(id)

	if args.Get(0) == nil {
		data.IsError = true
		return &data, nil
	}

	data.IsError = false
	return &data, nil
}
