package controller

import (
	"encoding/json"
	"net/http"

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
