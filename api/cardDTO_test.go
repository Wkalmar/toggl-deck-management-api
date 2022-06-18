package api

import (
	"testing"
	"toggl-deck-management-api/domain"
)

func TestCreateCardDTO(t *testing.T) {
	input := domain.Card{
		Suite: domain.Clubs,
		Value: domain.Seven,
	}
	actual := CreateCardDTO(input)
	expected := CardDTO{
		Rank:  "7",
		Shape: "CLUBS",
		Code:  "7C",
	}
	if actual != expected {
		t.Log("Incorrect value when creating CardDTO")
		t.Fail()
	}
}
