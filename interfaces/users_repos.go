package interfaces

import "github.com/biFebriansyah/gosolid/models"

type UsersRepoFaces interface {
	FindAll() (*models.Users, error)
	FindByID(id string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Save(user *models.User) error
	GetPassword(username string) (string, error)
}
