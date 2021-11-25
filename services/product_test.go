package services

import (
	"testing"

	"github.com/biFebriansyah/gosolid/models"
	"github.com/biFebriansyah/gosolid/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepos = repository.ProductRepoMock{Mocks: mock.Mock{}}
var productService = prod_service{rep: &productRepos}

func TestFindByID(t *testing.T) {
	productRepos.Mocks.On("FindByID", 1).Return(1)

	data, err := productService.GetByID("1")

	product := data.Data.(*models.Product)
	assert.Equal(t, 1, product.Id, "Expeted id 1")
	assert.Nil(t, err)
}
