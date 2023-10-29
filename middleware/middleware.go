package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-pgsql/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env files")
	}

	db, err := sql.Open("postgres", os.Getenv("PSQL_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to db")
	return db

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("unable to decode request body")

	}
	insertId := insertStock(stock)
	res := Response{
		ID:      insertId,
		Message: "Successfully created new Stock",
	}
	json.NewEncoder(w).Encode(res)
}
func GetAllStocks() {

}
func GetStock() {

}
func UpdateStock() {

}
func DeleteStock() {

}

func insertStock(stock models.Stock) int64 {
	return 100
}
