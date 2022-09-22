package api

import (
	"strings"
	"toggl-deck-management-api/domain"
	"toggl-deck-management-api/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateDeckArgs struct {
	Shuffled bool   `form:"shuffled"`
	Cards    string `form:"cards"`
}

type OpenDeckArgs struct {
	DeckId string `form:"deck_id"`
}

type DrawCardsArgs struct {
	DeckId string `form:"deck_id"`
	Count  uint8  `form:"count"`
}

// CreateDeckHandler godoc
// @Summary Creates new deck.
// @Description Creates deck that can be either shuffled or unshuffled. It can accept the list of exact cards which can be shuffled or unshuffled as well. In case no cards provided it returns a deck with 52 cards.
// @Accept */*
// @Produce json
// @Param shuffled query bool  true  "indicates whether deck is shuffled"
// @Param cards    query array false "array of card codes i.e. 8C,AS,7D"
// @Success 200 {object} ClosedDeckDTO
// @Router /create-deck [post]
func CreateDeckHandler(c *gin.Context) {
	var args CreateDeckArgs
	if c.ShouldBind(&args) == nil {
		var domainCards []domain.Card
		for _, card := range strings.Split(args.Cards, ",") {
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
		dto := CreateClosedDeckDTO(deck)
		c.JSON(200, dto)
		return
	} else {
		c.String(400, "Ivalid request. Expecting query of type ?shuffled=<bool>&cards=<card1>,<card2>,...<cardn>")
		return
	}
}

// OpenDeckHandler godoc
// @Summary Opens deck.
// @Description Returns a deck with all of its cars revealed.
// @Accept */*
// @Produce json
// @Param deck_id query string  true  "id of the deck"
// @Success 200 {object} OpenDeckDTO
// @Router /open-deck [get]
func OpenDeckHandler(c *gin.Context) {
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
		dto := CreateOpenDeckDTO(deck)
		c.JSON(200, dto)
		return
	} else {
		c.String(400, "Bad Request. Expecing request in format ?deck_id=<uuid>")
		return
	}
}

// DrawCardsHandler godoc
// @Summary Draws cards from a deck.
// @Description Removes given number of cards from a deck and returns them as a response.
// @Accept */*
// @Produce json
// @Param deck_id query string   true  "id of the deck"
// @Param count   query uint8    true  "number of cards to draw"
// @Success 200 {object} []CardDTO
// @Router /draw-cards [put]
func DrawCardsHandler(c *gin.Context) {
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
		var dto []CardDTO
		for _, card := range cards {
			dto = append(dto, CreateCardDTO(card))
		}
		storage.Add(deck)
		c.JSON(200, dto)
		return
	} else {
		c.String(400, "Bad Request. Expecting request in format ?deck_id=<uuid>&count=<uint8>")
		return
	}
}
