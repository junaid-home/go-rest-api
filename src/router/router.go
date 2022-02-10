package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	controller "api/src/controllers"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	FoodController := controller.FoodController{}

	router.HandleFunc("/food/all", FoodController.GetAllFoodItems(db)).Methods(http.MethodGet)
	router.HandleFunc("/food", FoodController.AddNewFoodItem(db)).Methods(http.MethodPost)

	return router
}
