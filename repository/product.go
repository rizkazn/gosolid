package repository

import (
	"database/sql"

	"github.com/biFebriansyah/gosolid/models"
)

type prodRepo struct {
	db *sql.DB
}

func Product(dbms *sql.DB) *prodRepo {
	return &prodRepo{dbms}
}

func (pr *prodRepo) FindByID(id int) (*models.Product, error) {
	query := `SELECT * FROM public.product WHERE id=$1`

	stm, err := pr.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var users models.Product

	err = stm.QueryRow(id).Scan(&users.Id, &users.Name, &users.Price, &users.CreatedAt, &users.UpdateAt)

	if err != nil {
		return nil, err
	}

	return &users, nil

}

func (pr *prodRepo) Save(name, price string) (*models.Product, error) {
	query := `INSERT INTO public.product ("name", price, created_at, update_at) VALUES($1, $2, $3, $4)`

	stm, err := pr.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	models := models.CreateProds()
	models.Name = name
	models.Price = price

	_, err = stm.Exec(models.Name, models.Price, models.CreatedAt, models.UpdateAt)

	if err != nil {
		return nil, err
	}

	return models, nil

}
