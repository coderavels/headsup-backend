package handlers

import (
	"net/http"
	"strconv"

	"github.com/coderavels/headsup-backend/entities"
	"github.com/gin-gonic/gin"
)

var decks = []*entities.Deck{
	{
		ID:   0,
		Name: "Actors",
		Cards: []*entities.Card{
			{
				ID:   0,
				Name: "Amitabh Bachchan",
			},
			{
				ID:   1,
				Name: "Shahrukh Khan",
			},
			{
				ID:   2,
				Name: "Abhishek Bachchan",
			},
			{
				ID:   3,
				Name: "Bobby Deol",
			},
			{
				ID:   4,
				Name: "Sunny Deol",
			},
			{
				ID:   5,
				Name: "Dharmendra",
			},
		},
	},
	{
		ID:   1,
		Name: "Movies",
		Cards: []*entities.Card{
			{
				ID:   0,
				Name: "Border",
			},
			{
				ID:   1,
				Name: "Shanghai",
			},
			{
				ID:   2,
				Name: "Kabhi Khushi Kabhi Gham",
			},
			{
				ID:   3,
				Name: "Dilwale",
			},
			{
				ID:   4,
				Name: "Kabul Express",
			},
			{
				ID:   5,
				Name: "Mission Kashmir",
			},
			{
				ID:   6,
				Name: "Mrityudaata",
			},
		},
	},
	{
		ID:   2,
		Name: "Food",
		Cards: []*entities.Card{
			{
				ID:   0,
				Name: "Chole Bhature",
			},
			{
				ID:   1,
				Name: "Palak Paneer",
			},
			{
				ID:   2,
				Name: "Puri Aloo",
			},
			{
				ID:   3,
				Name: "Pizza",
			},
			{
				ID:   4,
				Name: "Rajma Chawal",
			},
			{
				ID:   5,
				Name: "Dal Chawal",
			},
			{
				ID:   6,
				Name: "Akhuni",
			},
		},
	},
	{
		ID:   3,
		Name: "Cities",
		Cards: []*entities.Card{
			{
				ID:   0,
				Name: "Bangalore",
			},
			{
				ID:   1,
				Name: "Delhi",
			},
			{
				ID:   2,
				Name: "Kurseong",
			},
			{
				ID:   3,
				Name: "Tezpur",
			},
			{
				ID:   4,
				Name: "Mumbai",
			},
			{
				ID:   5,
				Name: "Srinagar",
			},
		},
	},
	{
		ID:   4,
		Name: "Animals",
		Cards: []*entities.Card{
			{

				Name: "Cat",
			},
			{
				ID:   1,
				Name: "Dog",
			},
			{
				ID:   2,
				Name: "Crocodile",
			},
			{
				ID:   3,
				Name: "Yak",
			},
			{
				ID:   4,
				Name: "Ox",
			},
			{
				ID:   5,
				Name: "Cow",
			},
		},
	},
}

type deckCategory struct {
	ID   int
	Name string
}

func GetDecks(c *gin.Context) {
	var deckCategories []*deckCategory
	for _, d := range decks {
		deckCategories = append(deckCategories, &deckCategory{ID: d.ID, Name: d.Name})
	}
	c.IndentedJSON(http.StatusOK, deckCategories)
}

func GetDeckCards(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var deck *entities.Deck

	for _, d := range decks {
		if d.ID == id {
			deck = d
		}
	}

	if deck == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "deck not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, deck.Cards)
}
