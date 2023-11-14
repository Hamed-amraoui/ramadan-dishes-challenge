package utils

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	"os"
	"ramadan/types"
	"strconv"
	"strings"
)

func GetDishes() ([]types.Dish, error) {
	var dishes []types.Dish
	file, _ := os.Open("utils/dishes.json")
	defer file.Close()
	var err = json.NewDecoder(file).Decode(&dishes)
	if err != nil {
		return dishes, err
	}
	return dishes, nil
}

func GetPrayTime(day int) types.Prays {
	data, _ := http.Get("https://api.aladhan.com/v1/hijriCalendarByCity/1445/9?city=Mecca&country=KSA")
	body, _ := io.ReadAll(data.Body)
    
	var response types.Response
	json.Unmarshal(body, &response)

	return response.Data[day].Prays;
}

func ParseDate(date string) types.DayTime {
	var time = strings.Split(date, " ")[0]
	var timeArray = strings.Split(time, ":")
	var hour, _ = strconv.Atoi(timeArray[0])
	var min, _ = strconv.Atoi(timeArray[1])

	return types.DayTime{ Hour: hour, Min: min}
}

func AddCookTime(date types.DayTime, cooktime int) types.DayTime {
	//fmt.Println("AddCookTime")
	date.Min += cooktime
	//fmt.Println(date)

	if date.Min >= 60 {
		date.Hour += date.Min / 60
		// fmt.Println(date)
		date.Min = date.Min % 60
	}

	return date
}

func AfterMaghrib(finishCookingTime types.DayTime, maghrib types.DayTime) bool {
	finishCookingTime = AddCookTime(finishCookingTime, 15)
	return finishCookingTime.Hour > maghrib.Hour || (finishCookingTime.Hour == maghrib.Hour && finishCookingTime.Min >= maghrib.Min)
}

func DiffBetweenFinishingCookAndMaghrib(finishCookingTime types.DayTime, maghrib types.DayTime) int {
	finishCookingTime = AddCookTime(finishCookingTime, 15)
	//fmt.Println(finishCookingTime)
	return (finishCookingTime.Hour - maghrib.Hour) * 60 + finishCookingTime.Min - maghrib.Min;
}


func IngredientInDish(ingredient string, dish types.Dish) bool {
	for _, dishIngredient := range dish.Ingredients {
		if dishIngredient == ingredient {
			return true
		}
	}
	return false
}

