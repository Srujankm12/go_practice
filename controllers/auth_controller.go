package controllers

import (
	"bunny/models"
	"encoding/json"
	"net/http"
)

type AuthController struct{
	authInterface models.AuthInterface
}

func NewAuthController(auth models.AuthInterface) *AuthController{
	return &AuthController{
		authInterface: auth,
	}
}

func(ac *AuthController) RegisterController(w http.ResponseWriter , r *http.Request) {
	str , err := ac.authInterface.Login(&r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(str)
}

func(ac *AuthController) LoginController(w http.ResponseWriter , r *http.Request) {
	str , err := ac.authInterface.Register(&r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(str)
}

