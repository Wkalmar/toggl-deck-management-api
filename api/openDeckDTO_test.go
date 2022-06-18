package api

import (
	"strconv"
	"testing"
	"toggl-deck-management-api/domain"
)

func TestCreateOpenDeckDTO(t *testing.T) {
	deck := domain.CreateDeck(false, domain.CreateCard(domain.Ace, domain.Clubs), domain.CreateCard(domain.Four, domain.Clubs))
	actual := CreateOpenDeckDTO(deck)
	if actual.Remaining != 2 {
		t.Log("Expected to have 2 cards remaining but was " + strconv.Itoa(int(actual.Remaining)))
		t.Fail()
	}
	if actual.Shuffled {
		t.Log("Expected deck not to be shuffled")
		t.Fail()
	}
	expectedCard0 := CardDTO{
		Rank:  "ACE",
		Shape: "CLUBS",
		Code:  "AC",
	}
	if actual.Cards[0] != expectedCard0 {
		t.Log("Invalid first card in the array of cards")
		t.Fail()
	}
	expectedCard1 := CardDTO{
		Rank:  "4",
		Shape: "CLUBS",
		Code:  "4C",
	}
	if actual.Cards[1] != expectedCard1 {
		t.Log("Invalid second card in the array of cards")
		t.Fail()
	}
}
