package interfaces

import (
	"github.com/biFebriansyah/gosolid/models"
)

type ProductRepos interface {
	FindByID(id int) (*models.Product, error)
}
