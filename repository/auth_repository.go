package repository

import (
	"database/sql"
	"io"
)

type AuthRepo struct{
	DB *sql.DB
}

func NewAuthRepo(sql *sql.DB) *AuthRepo{
	return &AuthRepo{
		DB: sql,
	}
}

func(ar *AuthRepo) Register(r *io.ReadCloser) (string , error) {
	return "Register" , nil
}

func(ar *AuthRepo) Login(r *io.ReadCloser)(string ,error){
	return "Login" , nil
}


