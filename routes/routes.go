package routes

import (
	"github.com/gorilla/mux"
	"ramadan/handlers"
)

func SetupRouter() *mux.Router {
  r := mux.NewRouter()

	r.HandleFunc("/suggest", handlers.SuggestHandler).Methods("GET")
	r.HandleFunc("/cooktime", handlers.CooktimeHandler).Methods("GET")
	
	return r
}
