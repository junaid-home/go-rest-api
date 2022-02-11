package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	controller "api/src/controllers"
	middleware "api/src/middlewares"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	FoodController := controller.FoodController{}

	router.HandleFunc("/food/all", middleware.CheckAuth(FoodController.GetAllFoodItems(db))).Methods(http.MethodGet)
	router.HandleFunc("/food", middleware.CheckAuth(FoodController.AddNewFoodItem(db))).Methods(http.MethodPost)
	router.HandleFunc("/food/{name}", middleware.CheckAuth(FoodController.GetSingleFoodItem(db))).Methods(http.MethodGet)
	router.HandleFunc("/food/{id}", middleware.CheckAuth(FoodController.DeleteSingleFoodItem(db))).Methods(http.MethodDelete)

	UserController := controller.UserController{}

	router.HandleFunc("/auth/login", UserController.LoginUser(db)).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", UserController.SignupUser(db)).Methods(http.MethodPost)

	return router
}
