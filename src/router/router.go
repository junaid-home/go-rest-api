package router

import (
	"github.com/gorilla/mux"

	controller "api/src/controllers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	FoodController := controller.FoodController{}

	router.HandleFunc("/food/all", FoodController.GetAllFoodItems()).Methods("GET")

	return router
}
