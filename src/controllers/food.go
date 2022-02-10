package controller

import (
	"api/src/helpers"
	"api/src/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type FoodController struct{}

func (f FoodController) GetAllFoodItems(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok Google!!!"))
	}
}

func (f FoodController) AddNewFoodItem(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		error := helpers.CustomError{}

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

		json.NewEncoder(w).Encode(FoodItem)
	}
}
