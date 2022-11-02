package api

import (
	"toggl-deck-management-api/domain"

	"github.com/google/uuid"
)

type OpenDeckDTO struct {
	DeckId    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining uint8     `json:"remaining"`
	Cards     []CardDTO `json:"cards"`
}

func createOpenDeckDTO(deck domain.Deck) OpenDeckDTO {
	var cards []CardDTO
	for _, domainCard := range deck.Cards {
		cards = append(cards, createCardDTO(domainCard))
	}
	return OpenDeckDTO{
		DeckId:    deck.DeckId,
		Shuffled:  deck.Shuffled,
		Remaining: domain.CountRemainingCards(deck),
		Cards:     cards,
	}
}
