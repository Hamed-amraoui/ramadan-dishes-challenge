package types

type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	Prays Prays `json:"timings"`
}

type Prays struct {
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
}

type DayTime struct {
	Hour	int
	Min 	int
}

type Dish struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Duration    int      `json:"duration"`
}

type DishInfo struct {
	Name       string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	StartCookingTime       string      `json:"startcookingtime"`
}
