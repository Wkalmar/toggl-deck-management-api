package api

import (
	"testing"
	"toggl-deck-management-api/domain"
)

type CreateCardDTODataItem struct {
	input  domain.Card
	actual CardDTO
}

var CreateCardDTOData = []CreateCardDTODataItem{
	{domain.CreateCard(domain.Ace, domain.Spades), CardDTO{"ACE", "SPADES", "AS"}},
	{domain.CreateCard(domain.Two, domain.Spades), CardDTO{"2", "SPADES", "2S"}},
	{domain.CreateCard(domain.Three, domain.Spades), CardDTO{"3", "SPADES", "3S"}},
	{domain.CreateCard(domain.Four, domain.Spades), CardDTO{"4", "SPADES", "4S"}},
	{domain.CreateCard(domain.Five, domain.Spades), CardDTO{"5", "SPADES", "5S"}},
	{domain.CreateCard(domain.Six, domain.Spades), CardDTO{"6", "SPADES", "6S"}},
	{domain.CreateCard(domain.Seven, domain.Spades), CardDTO{"7", "SPADES", "7S"}},
	{domain.CreateCard(domain.Eight, domain.Spades), CardDTO{"8", "SPADES", "8S"}},
	{domain.CreateCard(domain.Nine, domain.Spades), CardDTO{"9", "SPADES", "9S"}},
	{domain.CreateCard(domain.Ten, domain.Spades), CardDTO{"10", "SPADES", "10S"}},
	{domain.CreateCard(domain.Jack, domain.Spades), CardDTO{"JACK", "SPADES", "JS"}},
	{domain.CreateCard(domain.Queen, domain.Spades), CardDTO{"QUEEN", "SPADES", "QS"}},
	{domain.CreateCard(domain.King, domain.Spades), CardDTO{"KING", "SPADES", "KS"}},
	{domain.CreateCard(domain.Ace, domain.Hearts), CardDTO{"ACE", "HEARTS", "AH"}},
	{domain.CreateCard(domain.Two, domain.Hearts), CardDTO{"2", "HEARTS", "2H"}},
	{domain.CreateCard(domain.Three, domain.Hearts), CardDTO{"3", "HEARTS", "3H"}},
	{domain.CreateCard(domain.Four, domain.Hearts), CardDTO{"4", "HEARTS", "4H"}},
	{domain.CreateCard(domain.Five, domain.Hearts), CardDTO{"5", "HEARTS", "5H"}},
	{domain.CreateCard(domain.Six, domain.Hearts), CardDTO{"6", "HEARTS", "6H"}},
	{domain.CreateCard(domain.Seven, domain.Hearts), CardDTO{"7", "HEARTS", "7H"}},
	{domain.CreateCard(domain.Eight, domain.Hearts), CardDTO{"8", "HEARTS", "8H"}},
	{domain.CreateCard(domain.Nine, domain.Hearts), CardDTO{"9", "HEARTS", "9H"}},
	{domain.CreateCard(domain.Ten, domain.Hearts), CardDTO{"10", "HEARTS", "10H"}},
	{domain.CreateCard(domain.Jack, domain.Hearts), CardDTO{"JACK", "HEARTS", "JH"}},
	{domain.CreateCard(domain.Queen, domain.Hearts), CardDTO{"QUEEN", "HEARTS", "QH"}},
	{domain.CreateCard(domain.King, domain.Hearts), CardDTO{"KING", "HEARTS", "KH"}},
	{domain.CreateCard(domain.Ace, domain.Clubs), CardDTO{"ACE", "CLUBS", "AC"}},
	{domain.CreateCard(domain.Two, domain.Clubs), CardDTO{"2", "CLUBS", "2C"}},
	{domain.CreateCard(domain.Three, domain.Clubs), CardDTO{"3", "CLUBS", "3C"}},
	{domain.CreateCard(domain.Four, domain.Clubs), CardDTO{"4", "CLUBS", "4C"}},
	{domain.CreateCard(domain.Five, domain.Clubs), CardDTO{"5", "CLUBS", "5C"}},
	{domain.CreateCard(domain.Six, domain.Clubs), CardDTO{"6", "CLUBS", "6C"}},
	{domain.CreateCard(domain.Seven, domain.Clubs), CardDTO{"7", "CLUBS", "7C"}},
	{domain.CreateCard(domain.Eight, domain.Clubs), CardDTO{"8", "CLUBS", "8C"}},
	{domain.CreateCard(domain.Nine, domain.Clubs), CardDTO{"9", "CLUBS", "9C"}},
	{domain.CreateCard(domain.Ten, domain.Clubs), CardDTO{"10", "CLUBS", "10C"}},
	{domain.CreateCard(domain.Jack, domain.Clubs), CardDTO{"JACK", "CLUBS", "JC"}},
	{domain.CreateCard(domain.Queen, domain.Clubs), CardDTO{"QUEEN", "CLUBS", "QC"}},
	{domain.CreateCard(domain.King, domain.Clubs), CardDTO{"KING", "CLUBS", "KC"}},
	{domain.CreateCard(domain.Ace, domain.Diamonds), CardDTO{"ACE", "DIAMONDS", "AD"}},
	{domain.CreateCard(domain.Two, domain.Diamonds), CardDTO{"2", "DIAMONDS", "2D"}},
	{domain.CreateCard(domain.Three, domain.Diamonds), CardDTO{"3", "DIAMONDS", "3D"}},
	{domain.CreateCard(domain.Four, domain.Diamonds), CardDTO{"4", "DIAMONDS", "4D"}},
	{domain.CreateCard(domain.Five, domain.Diamonds), CardDTO{"5", "DIAMONDS", "5D"}},
	{domain.CreateCard(domain.Six, domain.Diamonds), CardDTO{"6", "DIAMONDS", "6D"}},
	{domain.CreateCard(domain.Seven, domain.Diamonds), CardDTO{"7", "DIAMONDS", "7D"}},
	{domain.CreateCard(domain.Eight, domain.Diamonds), CardDTO{"8", "DIAMONDS", "8D"}},
	{domain.CreateCard(domain.Nine, domain.Diamonds), CardDTO{"9", "DIAMONDS", "9D"}},
	{domain.CreateCard(domain.Ten, domain.Diamonds), CardDTO{"10", "DIAMONDS", "10D"}},
	{domain.CreateCard(domain.Jack, domain.Diamonds), CardDTO{"JACK", "DIAMONDS", "JD"}},
	{domain.CreateCard(domain.Queen, domain.Diamonds), CardDTO{"QUEEN", "DIAMONDS", "QD"}},
	{domain.CreateCard(domain.King, domain.Diamonds), CardDTO{"KING", "DIAMONDS", "KD"}},
}

func TestCreateCardDTO(t *testing.T) {
	for _, item := range CreateCardDTOData {
		actual := CreateCardDTO(item.input)
		if actual != item.actual {
			t.Log("Incorrect value when creating CardDTO")
			t.Fail()
		}
	}
}

type TestParseCardStringCodeDataItem struct {
	code    string
	isError bool
	card    domain.Card
}

var TestParseCardStringCodeData = []TestParseCardStringCodeDataItem{
	{"13F", true, domain.Card{}},
	{"9F", true, domain.Card{}},
	{"13S", true, domain.Card{}},
	{"8C", false, domain.CreateCard(domain.Eight, domain.Clubs)},
	{"10S", false, domain.CreateCard(domain.Ten, domain.Spades)},
}

func TestParseCardStringCode(t *testing.T) {
	for _, data := range TestParseCardStringCodeData {
		actual, err := ParseCardStringCode(data.code)
		if data.isError && err == nil {
			t.Log("Expected to fail for " + data.code)
			t.Fail()
		} else if !data.isError && err != nil {
			t.Log("Expected to succeed for " + data.code)
			t.Fail()
		}
		if actual != data.card {
			t.Log("Expected " + GetCardStringCode(data.card) + " but was " + GetCardStringCode(actual))
			t.Fail()
		}
	}
}
