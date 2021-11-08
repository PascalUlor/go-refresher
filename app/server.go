package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	// "example.com/controllers"
	"example.com/routes"
)


func init() {
	gotenv.Load()
}

//Server app server

var server = routes.Server{}
// type Server struct {
// 	Router *mux.Router
// }

//TODO app server
func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
	done := make(chan bool)
	go http.ListenAndServe(":" + port, server.Router)
	log.Printf("Server started at port %v", port)
	<-done
	log.Fatal(http.ListenAndServe(":" + port, server.Router))
}