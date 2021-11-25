package interfaces

import "github.com/biFebriansyah/gosolid/helpers"

type ProductServices interface {
	GetByID(id string) (*helpers.Response, error)
}
