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

func Get(id uuid.UUID) (domain.Deck, bool) {
	item, found := data.Load(id)
	return item.(domain.Deck), found
}
