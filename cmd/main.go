package main

import (
	"avito_internship_task/internal/handler"
	"avito_internship_task/internal/middleware"
	"avito_internship_task/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	db := utils.InitDB()

	authHandler := handler.CreateAuthHandler(db)
	transactionsHandler := handler.CreateTransactionsHandler(db)
	merchHandler := handler.CreateMerchHandler(db)

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/auth", authHandler.Login).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/merch", merchHandler.GetMerchList).Methods("GET")
	protected.HandleFunc("/buy/{item}", merchHandler.BuyMerch).Methods("GET")
	protected.HandleFunc("/sendCoin", transactionsHandler.TransferCoins).Methods("POST")
	protected.HandleFunc("/info", transactionsHandler.GetTransactionsInfo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
