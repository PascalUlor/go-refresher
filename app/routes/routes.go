package routes


import (
	"example.com/middlewares"
	"example.com/controllers"

	"github.com/gorilla/mux"
)

var controller = controllers.Controllers{}

type Server struct {
	Router *mux.Router
}

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	//TODO default routes
	server.Router.HandleFunc("/", middlewares.ResponseJSON(controller.Default)).Methods("GET")
}