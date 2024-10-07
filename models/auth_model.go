package models

import "io"

type AuthDetails struct{
	UserID string `json:"userid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthInterface interface{
	Register(*io.ReadCloser) (string , error)
	Login(*io.ReadCloser) (string , error)
}