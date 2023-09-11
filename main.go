package main

import (
	"github.com/gin-gonic/gin"

	"github.com/coderavels/headsup-backend/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/decks", handlers.GetDecks)
	router.GET("/decks/:id/cards", handlers.GetDeckCards)

	router.Run("localhost:8080")
}
