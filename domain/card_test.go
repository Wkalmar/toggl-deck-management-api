package domain

import "testing"

type GetCardStringCodeDataItem struct {
	rank  Rank
	shape Shape
	code  string
}

var GetCardStringCodeData = []GetCardStringCodeDataItem{
	{Ace, Spades, "AS"},
	{Two, Spades, "2S"},
	{Three, Spades, "3S"},
	{Four, Spades, "4S"},
	{Five, Spades, "5S"},
	{Six, Spades, "6S"},
	{Seven, Spades, "7S"},
	{Eight, Spades, "8S"},
	{Nine, Spades, "9S"},
	{Ten, Spades, "10S"},
	{Jack, Spades, "JS"},
	{Queen, Spades, "QS"},
	{King, Spades, "KS"},
	{Ace, Hearts, "AH"},
	{Two, Hearts, "2H"},
	{Three, Hearts, "3H"},
	{Four, Hearts, "4H"},
	{Five, Hearts, "5H"},
	{Six, Hearts, "6H"},
	{Seven, Hearts, "7H"},
	{Eight, Hearts, "8H"},
	{Nine, Hearts, "9H"},
	{Ten, Hearts, "10H"},
	{Jack, Hearts, "JH"},
	{Queen, Hearts, "QH"},
	{King, Hearts, "KH"},
	{Ace, Clubs, "AC"},
	{Two, Clubs, "2C"},
	{Three, Clubs, "3C"},
	{Four, Clubs, "4C"},
	{Five, Clubs, "5C"},
	{Six, Clubs, "6C"},
	{Seven, Clubs, "7C"},
	{Eight, Clubs, "8C"},
	{Nine, Clubs, "9C"},
	{Ten, Clubs, "10C"},
	{Jack, Clubs, "JC"},
	{Queen, Clubs, "QC"},
	{King, Clubs, "KC"},
	{Ace, Diamonds, "AD"},
	{Two, Diamonds, "2D"},
	{Three, Diamonds, "3D"},
	{Four, Diamonds, "4D"},
	{Five, Diamonds, "5D"},
	{Six, Diamonds, "6D"},
	{Seven, Diamonds, "7D"},
	{Eight, Diamonds, "8D"},
	{Nine, Diamonds, "9D"},
	{Ten, Diamonds, "10D"},
	{Jack, Diamonds, "JD"},
	{Queen, Diamonds, "QD"},
	{King, Diamonds, "KD"},
}

func TestGetCardStringCode(t *testing.T) {
	for _, item := range GetCardStringCodeData {
		card := CreateCard(item.rank, item.shape)
		code := GetCardStringCode(card)
		if code != item.code {
			t.Log("Expected code to be " + item.code + " but was " + code)
			t.Fail()
		}
	}
}
