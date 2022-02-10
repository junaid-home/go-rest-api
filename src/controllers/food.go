package controller

import "net/http"

type FoodController struct{}

func (f FoodController) GetAllFoodItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok Google!!!"))
	}
}
