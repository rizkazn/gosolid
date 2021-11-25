package controllers

import (
	"github.com/biFebriansyah/gosolid/interfaces"
)

type users_controll struct {
	rep interfaces.UsersServiceFaces
}

func UsersCtrl(rps interfaces.UsersServiceFaces) *users_controll {
	return &users_controll{rps}
}
