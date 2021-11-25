package interfaces

import "net/http"

type ProductControll interface {
	GetByID(w http.ResponseWriter, r *http.Request)
}
