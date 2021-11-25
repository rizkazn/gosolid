package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/biFebriansyah/gosolid/helpers"
	"github.com/biFebriansyah/gosolid/models"
	"github.com/biFebriansyah/gosolid/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var prodRepo = services.ProductServMock{Mocks: mock.Mock{}}
var prodsCtrl = ProdsCtrl(&prodRepo)

func TestGetByID(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/{id}", prodsCtrl.GetByID)
	prodRepo.Mocks.On("GetByID", "1").Return("1")

	r.ServeHTTP(w, httptest.NewRequest("GET", "/test/1", nil))

	var prods models.Product
	respone := helpers.Response{
		Data: &prods,
	}

	err := json.Unmarshal(w.Body.Bytes(), &respone)

	if err != nil {
		t.Fatal("salah")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.False(t, respone.IsError, "is Error mustbe false")
}
