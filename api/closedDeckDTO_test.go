package api

import (
	"testing"
	"toggl-deck-management-api/domain"

	"github.com/stretchr/testify/assert"
)

type TestCreateClosedDeckDTODataItem struct {
	shuffled  bool
	cards     []domain.Card
	remaining uint8
}

var TestCreateClosedDeckDTOData = []TestCreateClosedDeckDTODataItem{
	{false, []domain.Card{}, 52},
	{true, []domain.Card{}, 52},
	{false, []domain.Card{domain.CreateCard(domain.Ace, domain.Clubs), domain.CreateCard(domain.Five, domain.Diamonds)}, 2},
}

func TestCreateClosedDeckDTO(t *testing.T) {
	for _, item := range TestCreateClosedDeckDTOData {
		deck := domain.CreateDeck(item.shuffled, item.cards...)
		dto := CreateClosedDeckDTO(deck)
		assert.Equal(t, item.shuffled, dto.Shuffled)
	}

}
