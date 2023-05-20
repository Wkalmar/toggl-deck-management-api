package storage

import (
	"sync"
	"toggl-deck-management-api/domain"

	"github.com/google/uuid"
)

var data = sync.Map{}

func Add(deck domain.Deck) {
	data.Store(deck.DeckId, deck)
}

func Get(id uuid.UUID) (*domain.Deck, bool) {
	item, found := data.Load(id)
	if found {
		deck := item.(domain.Deck)
		return &deck, found
	}
	return nil, found
}
