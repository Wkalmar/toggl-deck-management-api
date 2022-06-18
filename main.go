package main

import (
	"toggl-deck-management-api/api"
	"toggl-deck-management-api/domain"
	"toggl-deck-management-api/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateDeckArgs struct {
	Shuffled bool     `form:"shuffled"`
	Cards    []string `form:"cards"`
}

type OpenDeckArgs struct {
	DeckId string `form:"deck_id"`
}

type DrawCardsArgs struct {
	DeckId string `form:"deck_id"`
	Count  uint8  `form:"count"`
}

func main() {
	r := gin.Default()
	r.POST("/create-deck", func(c *gin.Context) {
		var args CreateDeckArgs
		if c.ShouldBind(&args) == nil {
			var domainCards []domain.Card
			for _, card := range args.Cards {
				domainCard, err := domain.ParseCardStringCode(card)
				if err == nil {
					domainCards = append(domainCards, domainCard)
				} else {
					c.String(400, "Invalid request. Invalid card code "+card)
					return
				}
			}
			deck := domain.CreateDeck(args.Shuffled, domainCards...)
			storage.Add(deck)
			dto := api.CreateClosedDeckDTO(deck)
			c.JSON(200, dto)
			return
		} else {
			c.String(400, "Ivalid request. Expecting query of type ?shuffled=<bool>&cards=<card1>,<card2>,...<cardn>")
			return
		}
	})
	r.GET("/open-deck", func(c *gin.Context) {
		var args OpenDeckArgs
		if c.ShouldBind(&args) == nil {
			deckId, err := uuid.Parse(args.DeckId)
			if err != nil {
				c.String(400, "Bad Request. Expecing request in format ?deck_id=<uuid>")
				return
			}
			deck, found := storage.Get(deckId)
			if !found {
				c.String(400, "Bad Request. Deck with given id not found")
				return
			}
			dto := api.CreateOpenDeckDTO(deck)
			c.JSON(200, dto)
			return
		} else {
			c.String(400, "Bad Request. Expecing request in format ?deck_id=<uuid>")
			return
		}
	})
	r.PUT("/draw-cards", func(c *gin.Context) {
		var args DrawCardsArgs
		if c.ShouldBind(&args) == nil {
			deckId, err := uuid.Parse(args.DeckId)
			if err != nil {
				c.String(400, "Bad Request. Expecing request in format ?deck_id=<uuid>")
				return
			}
			deck, found := storage.Get(deckId)
			if !found {
				c.String(400, "Bad Request. Expecting request in format ?deck_id=<uuid>&count=<uint8>")
				return
			}
			cards, err := domain.DrawCards(&deck, args.Count)
			if err != nil {
				c.String(400, "Bad Request. Failed to draw cards from the deck")
				return
			}
			var dto []api.CardDTO
			for _, card := range cards {
				dto = append(dto, api.CreateCardDTO(card))
			}
			storage.Add(deck)
			c.JSON(200, dto)
			return
		} else {
			c.String(400, "Bad Request. Expecting request in format ?deck_id=<uuid>&count=<uint8>")
			return
		}
	})
	r.Run()
}
