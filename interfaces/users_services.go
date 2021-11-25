package interfaces

import "github.com/biFebriansyah/gosolid/helpers"

type UsersServiceFaces interface {
	FindPassword(username string) (string, error)
	Save(name, email, username, password string) (helpers.Response, error)
	GetAll() (helpers.Response, error)
	GetById(id string) (helpers.Response, error)
}
