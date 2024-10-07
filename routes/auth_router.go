package routes

import (
	"bunny/controllers"
	"bunny/repository"
	"database/sql"
	"net/http"
)


func AuthRouter(db *sql.DB , mux *http.ServeMux){
	repo := repository.NewAuthRepo(db)
	cont := controllers.NewAuthController(repo)

	mux.HandleFunc("/login" , cont.LoginController)
	mux.HandleFunc("/register" , cont.RegisterController)
}