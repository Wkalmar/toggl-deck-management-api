package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CreateDeckDataItem struct {
	index int
	shape Shape
	rank  Rank
}

var CreateUnshuffledDeckData = []CreateDeckDataItem{
	{0, Spades, Ace},
	{1, Spades, Two},
	{2, Spades, Three},
	{3, Spades, Four},
	{4, Spades, Five},
	{5, Spades, Six},
	{6, Spades, Seven},
	{7, Spades, Eight},
	{8, Spades, Nine},
	{9, Spades, Ten},
	{10, Spades, Jack},
	{11, Spades, Queen},
	{12, Spades, King},
	{13, Diamonds, Ace},
	{14, Diamonds, Two},
	{15, Diamonds, Three},
	{16, Diamonds, Four},
	{17, Diamonds, Five},
	{18, Diamonds, Six},
	{19, Diamonds, Seven},
	{20, Diamonds, Eight},
	{21, Diamonds, Nine},
	{22, Diamonds, Ten},
	{23, Diamonds, Jack},
	{24, Diamonds, Queen},
	{25, Diamonds, King},
	{26, Clubs, Ace},
	{27, Clubs, Two},
	{28, Clubs, Three},
	{29, Clubs, Four},
	{30, Clubs, Five},
	{31, Clubs, Six},
	{32, Clubs, Seven},
	{33, Clubs, Eight},
	{34, Clubs, Nine},
	{35, Clubs, Ten},
	{36, Clubs, Jack},
	{37, Clubs, Queen},
	{38, Clubs, King},
	{39, Hearts, Ace},
	{40, Hearts, Two},
	{41, Hearts, Three},
	{42, Hearts, Four},
	{43, Hearts, Five},
	{44, Hearts, Six},
	{45, Hearts, Seven},
	{46, Hearts, Eight},
	{47, Hearts, Nine},
	{48, Hearts, Ten},
	{49, Hearts, Jack},
	{50, Hearts, Queen},
	{51, Hearts, King},
}

func TestCreateDeck_Unshuffled(t *testing.T) {
	unshuffledDeck := CreateDeck(false)
	assert.False(t, unshuffledDeck.Shuffled)
	for _, item := range CreateUnshuffledDeckData {
		card := unshuffledDeck.Cards[item.index]
		assert.Equal(t, item.rank, card.Rank)
		assert.Equal(t, item.shape, card.Shape)
	}
}

func TestCreateDeck_Shuffled(t *testing.T) {
	shuffledDeck := CreateDeck(true)
	assert.True(t, shuffledDeck.Shuffled)
	cardCount := make(map[Card]int)
	for _, card := range shuffledDeck.Cards {
		value, exists := cardCount[card]
		if exists {
			value++
			cardCount[card] = value
		} else {
			cardCount[card] = 1
		}
	}
	for _, oracle := range unshuffledCards {
		value, exists := cardCount[oracle]
		if exists {
			if value != 1 {
				t.Log("Expecting each card to be present only once")
				t.Fail()
			}
		} else {
			t.Log("Expecting each card to be present")
			t.Fail()
		}
	}
}

func TestCreateDeck_Shuffled_ThenUnshuffled(t *testing.T) {
	shuffledDeck := CreateDeck(true)
	assert.True(t, shuffledDeck.Shuffled)
	unshuffledDeck := CreateDeck(false)
	for _, item := range CreateUnshuffledDeckData {
		card := unshuffledDeck.Cards[item.index]
		assert.Equal(t, item.rank, card.Rank)
		assert.Equal(t, item.shape, card.Shape)
	}
}

func TestCreateDeck_ExactCardsArePassed_Unshuffled(t *testing.T) {
	jackOfDiamonds := CreateCard(Jack, Diamonds)
	aceOfSpades := CreateCard(Ace, Spades)
	queenOfHearts := CreateCard(Queen, Hearts)
	cards := []Card{jackOfDiamonds, aceOfSpades, queenOfHearts}
	deck := CreateDeck(false, cards...)
	for i, inputCard := range cards {
		card := deck.Cards[i]
		assert.Equal(t, inputCard, card)
	}
}

func TestCreateDeck_ExactCardsArePassed_Shuffled(t *testing.T) {
	jackOfDiamonds := CreateCard(Jack, Diamonds)
	aceOfSpades := CreateCard(Ace, Spades)
	queenOfHearts := CreateCard(Queen, Hearts)
	cards := []Card{jackOfDiamonds, aceOfSpades, queenOfHearts}
	deck := CreateDeck(false, cards...)
	deckCardsCount := make(map[Card]int)
	for _, resCard := range deck.Cards {
		value, exists := deckCardsCount[resCard]
		if exists {
			value++
			deckCardsCount[resCard] = value
		} else {
			deckCardsCount[resCard] = 1
		}
	}
	for _, inputCard := range cards {
		value, found := deckCardsCount[inputCard]
		assert.True(t, found, "Expected all cards to be present")
		assert.Equal(t, 1, value, "Expected cards not to be duplicate")
	}
}

func TestDrawCards_SufficientAmount(t *testing.T) {
	deck := CreateDeck(true)
	items, err := DrawCards(&deck, 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(items))
	remaining := CountRemainingCards(deck)
	assert.Equal(t, uint8(50), remaining)
}

func TestDrawCards_SufficientAmount_NoItemsRemain(t *testing.T) {
	deck := CreateDeck(true)
	items, err := DrawCards(&deck, 52)
	assert.Nil(t, err)
	assert.Equal(t, 52, len(items))
	reamining := CountRemainingCards(deck)
	assert.Equal(t, uint8(0), reamining)
}

func TestDrawCards_InsufficientAmount(t *testing.T) {
	deck := CreateDeck(true)
	_, err := DrawCards(&deck, 53)
	assert.NotNil(t, err)
}
