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
	router.HandleFunc("/food/{name}", FoodController.GetSingleFoodItem(db)).Methods(http.MethodGet)
	router.HandleFunc("/food/{id}", FoodController.DeleteSingleFoodItem(db)).Methods(http.MethodDelete)

	UserController := controller.UserController{}

	router.HandleFunc("/auth/login", UserController.LoginUser(db)).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", UserController.SignupUser(db)).Methods(http.MethodPost)

	return router
}
