package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"ramadan/types"
	"ramadan/utils"
	"strconv"
)

// cooktime handler
func CooktimeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ingredient := query.Get("ingredient")
	selectedDay := query.Get("day")
	if ingredient == "" || selectedDay == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dishes, err := utils.GetDishes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	day, _ := strconv.Atoi(selectedDay)
	prays := utils.GetPrayTime(day)
	asr := utils.ParseDate(prays.Asr)
	maghrib := utils.ParseDate(prays.Maghrib)

	var result []types.DishInfo

	for _, dish := range dishes {
		if utils.IngredientInDish(ingredient, dish) {
			// Calculate cook time for the current dish
			minutes := utils.DiffBetweenFinishingCookAndMaghrib(
				utils.AddCookTime(asr, dish.Duration),
				maghrib,
			)
			StartCookingTime := "After Asr"
			if minutes > 0 {
				StartCookingTime = "Before Asr"
			}

			AbsMin := float64(minutes)

			dishInfo := types.DishInfo{
				Name:       dish.Name,
				Ingredients: dish.Ingredients,
				StartCookingTime: fmt.Sprintf("%.0f minutes %s", math.Abs(AbsMin), StartCookingTime),
			}
			result = append(result, dishInfo)
		}
	}

	responseJSON, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}


// suggest handler 
func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	selectedDay := query.Get("day")
	if selectedDay == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dishes, err := utils.GetDishes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	day, _ := strconv.Atoi(selectedDay)
	prays := utils.GetPrayTime(day)
	asr := utils.ParseDate(prays.Asr)
	maghrib := utils.ParseDate(prays.Maghrib)

	// Generate a random index
	randIndex := rand.Intn(len(dishes))

	dish := dishes[randIndex]

	// Calculate cook time for the random dish
	minutes := utils.DiffBetweenFinishingCookAndMaghrib(
		utils.AddCookTime(asr, dish.Duration),
		maghrib,
	)
	StartCookingTime := "After Asr"
	if minutes > 0 {
		StartCookingTime = "Before Asr"
	}

	AbsMin := float64(minutes)

	result := types.DishInfo{
		Name:             dish.Name,
		Ingredients:      dish.Ingredients,
		StartCookingTime: fmt.Sprintf("%.0f minutes %s", math.Abs(AbsMin), StartCookingTime),
	}

	responseJSON, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}



