package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	// "ramadan/types"
	"ramadan/routes"
)

func main() {
	// Setup the server
	r := routes.SetupRouter()

	// Start the server
	fmt.Println("Server started on port 3000")
	http.ListenAndServe(":3000", r)

	// response, _ := http.Get("https://api.aladhan.com/v1/hijriCalendarByCity/1445/9?city=Mecca&country=KSA")
	// responseData, _ := ioutil.ReadAll(response.Body)
    
    // var responseJson types.Response
    // json.Unmarshal(responseData, &responseJson)

    // fmt.Println(responseJson.Data[0].Prays)


	// var asr = utils.ParseDate("15:54 (+03)")
	// var finishCookingTime = utils.AddCookTime(asr, 150)
	// var maghrib = utils.ParseDate("18:25 (+03)")
	// fmt.Println(utils.AfterMaghrib(finishCookingTime, maghrib))
}
