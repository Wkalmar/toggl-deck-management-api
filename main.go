package main

import (
	"toggl-deck-management-api/api"
	"toggl-deck-management-api/domain"
	"toggl-deck-management-api/storage"

	"github.com/gin-gonic/gin"
)

type CreateDeckArgs struct {
	Shuffled bool     `form:"shuffled"`
	Cards    []string `form:"cards"`
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
	r.Run()
}
