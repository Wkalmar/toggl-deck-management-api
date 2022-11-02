package api

import (
	"testing"
	"toggl-deck-management-api/domain"

	"github.com/stretchr/testify/assert"
)

func TestCreateOpenDeckDTO(t *testing.T) {
	deck := domain.CreateDeck(false, domain.CreateCard(domain.Ace, domain.Clubs), domain.CreateCard(domain.Four, domain.Clubs))
	actual := createOpenDeckDTO(deck)
	assert.Equal(t, uint8(2), actual.Remaining)
	assert.False(t, actual.Shuffled)
	expectedCard0 := CardDTO{
		Rank:  "ACE",
		Shape: "CLUBS",
		Code:  "AC",
	}
	assert.Equal(t, expectedCard0, actual.Cards[0])
	expectedCard1 := CardDTO{
		Rank:  "4",
		Shape: "CLUBS",
		Code:  "4C",
	}
	assert.Equal(t, expectedCard1, actual.Cards[1])
}
