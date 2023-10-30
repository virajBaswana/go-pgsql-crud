package router

import (
	"golang-pgsql/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	//check
	// router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newStock", middleware.CreateStock).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/stocks/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/stocks/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}
