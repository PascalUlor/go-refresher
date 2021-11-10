package controllers

import (
	"net/http"

	"example.com/helpers"
)


type Controllers struct {}

//TODO default route
func (controller *Controllers) Default(w http.ResponseWriter, r *http.Request) {
	helpers.JSON(w, http.StatusOK, "Welcome to my Crypto App")
}