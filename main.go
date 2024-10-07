package main

import (
	"bunny/routes"
	"log"
	"net/http"
)

func main(){
	// db := database.NewDbDetails("postgres" , "url")
	// conn , err := db.ConnectToDatabase()
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()
	mux := http.NewServeMux()
	routes.AuthRouter(nil , mux)

	log.Println("Server Is Started...")
	log.Fatal(http.ListenAndServe(":8000" , mux))
}