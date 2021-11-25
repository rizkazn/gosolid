package repository

import (
	"database/sql"

	"github.com/biFebriansyah/gosolid/models"
)

type initRepo struct {
	db *sql.DB
}

func Users(dbms *sql.DB) *initRepo {
	return &initRepo{dbms}
}

func (r *initRepo) FindAll() (*models.Users, error) {
	query := `SELECT * FROM public.users`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var data models.Users
	var users models.User

	for rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.UserName, &users.Email, &users.Password, &users.CreatedAt, &users.UpdateAt)

		if err != nil {
			return nil, err
		}

		data = append(data, users)
	}

	return &data, nil

}

func (r *initRepo) FindByID(id string) (*models.User, error) {
	query := `SELECT * FROM public.users WHERE id=$1`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var users models.User

	err = stm.QueryRow(id).Scan(&users.Id, &users.Name, &users.UserName, &users.Email, &users.Password, &users.CreatedAt, &users.UpdateAt)

	if err != nil {
		return nil, err
	}

	return &users, nil

}

func (r *initRepo) FindByUsername(username string) (*models.User, error) {
	query := `SELECT * FROM public.users WHERE username = $1`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var users models.User

	err = stm.QueryRow(username).Scan(&users.Id, &users.Name, &users.UserName, &users.Email, &users.Password, &users.CreatedAt, &users.UpdateAt)

	if err != nil {
		return nil, err
	}

	return &users, nil

}

func (r *initRepo) Save(user *models.User) error {
	query := `INSERT INTO public.users("name", username, email, "password", update_at, created_at) VALUES($1, $2, $3, $4, $5, $6)`

	stm, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stm.Exec(user.Name, user.UserName, user.Email, user.Password, user.CreatedAt, user.UpdateAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *initRepo) GetPassword(username string) (string, error) {
	query := `SELECT password FROM public.users WHERE username = $1`
	rows, err := r.db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var passwords string

	for rows.Next() {
		err := rows.Scan(&passwords)

		if err != nil {
			return "", err
		}
	}

	return passwords, nil

}
