package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"api/src/helpers"
	"api/src/models"
)

type FoodController struct{}

var error = helpers.CustomError{}

func (f FoodController) GetAllFoodItems(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		FoodItems := []models.Food{}

		if results := db.Find(&FoodItems); results.Error != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Fetch Food Items from database!")
			return
		}

		helpers.RespondWithJSON(w, FoodItems)
	}
}

func (f FoodController) GetSingleFoodItem(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		FoodItem := models.Food{}

		if results := db.Where("name = ?", params["name"]).First(&FoodItem); results.Error != nil || results.RowsAffected < 1 {
			error.ApiError(w, http.StatusNotFound, "Didn't Find food item with name = "+params["name"])
			return
		}

		helpers.RespondWithJSON(w, FoodItem)
	}
}

func (f FoodController) AddNewFoodItem(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		FoodItem := models.Food{}
		json.NewDecoder(r.Body).Decode(&FoodItem)

		if len(FoodItem.Name) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
			return
		}

		if FoodItem.Quantity == 0 {
			error.ApiError(w, http.StatusBadRequest, "Quantity Shouldn't be zero!")
			return
		}

		if len(FoodItem.Selling_Price) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Invalid Selling Price!")
			return
		}

		if result := db.Create(&FoodItem); result.Error != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Add new Food Item in database!")
			return
		}

		helpers.RespondWithJSON(w, FoodItem)
	}
}

func (f FoodController) DeleteSingleFoodItem(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		FoodItem := models.Food{}

		if results := db.Where("id = ?", params["id"]).First(&FoodItem); results.Error != nil || results.RowsAffected < 1 {
			error.ApiError(w, http.StatusNotFound, "Didn't Find food item with id = "+params["id"])
			return
		}

		if results := db.Delete(&FoodItem); results.Error != nil || results.RowsAffected < 1 {
			error.ApiError(w, http.StatusInternalServerError, "Failed to Delete Item from the database!")
			return
		}

		helpers.RespondWithJSON(w, FoodItem)
	}
}
