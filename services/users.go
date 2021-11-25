package services

import "github.com/biFebriansyah/gosolid/interfaces"

type user_service struct {
	rep interfaces.UsersRepoFaces
}

func UserServices(reps interfaces.UsersRepoFaces) *user_service {
	return &user_service{reps}
}
