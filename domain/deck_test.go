package domain

import (
	"strconv"
	"testing"
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
	if unshuffledDeck.Shuffled {
		t.Log("Property shuffled of unshuffled deck expected to be false")
		t.Fail()
	}
	for _, item := range CreateUnshuffledDeckData {
		card := unshuffledDeck.Cards[item.index]
		if card.Shape != item.shape || card.Rank != item.rank {
			t.Log("First card expected to be " + GetCardStringCode(CreateCard(item.rank, item.shape)) + " but was " + GetCardStringCode(card))
			t.Fail()
		}
	}
}

func TestCreateDeck_Shuffled(t *testing.T) {
	shuffledDeck := CreateDeck(true)
	if !shuffledDeck.Shuffled {
		t.Log("Property shuffled of shuffled deck expected to be true")
		t.Fail()
	}
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
	if !shuffledDeck.Shuffled {
		t.Log("Property shuffled of shuffled deck expected to be true")
		t.Fail()
	}
	unshuffledDeck := CreateDeck(false)
	for _, item := range CreateUnshuffledDeckData {
		card := unshuffledDeck.Cards[item.index]
		if card.Shape != item.shape || card.Rank != item.rank {
			t.Log("Card expected to be " + GetCardStringCode(CreateCard(item.rank, item.shape)) + " but was " + GetCardStringCode(card))
			t.Fail()
		}
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
		if card != inputCard {
			t.Log("Card expected to be " + GetCardStringCode(inputCard) + " but was " + GetCardStringCode(card))
			t.Fail()
		}
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
		if !found {
			t.Log("Expected all cards to be present")
			t.Fail()
		}
		if value != 1 {
			t.Log("Expected cards not to be duplicate")
			t.Fail()
		}
	}
}

func TestDrawCards_SufficientAmount(t *testing.T) {
	deck := CreateDeck(true)
	items, err := DrawCards(&deck, 2)
	if err != nil {
		t.Log("Expected to succeed")
		t.Fail()
	}
	if len(items) != 2 {
		t.Log("Expected to return 2 items but was " + strconv.Itoa((len(items))))
		t.Fail()
	}
	reamining := countRemainingCardsCore(deck)
	if reamining != 50 {
		t.Log("Excpected 50 items to remain but was " + strconv.Itoa(int(reamining)))
		t.Fail()
	}
}

func TestDrawCards_SufficientAmount_NoItemsRemain(t *testing.T) {
	deck := CreateDeck(true)
	items, err := DrawCards(&deck, 52)
	if err != nil {
		t.Log("Expected to succeed")
		t.Fail()
	}
	if len(items) != 52 {
		t.Log("Expected to return 2 items but was " + strconv.Itoa((len(items))))
		t.Fail()
	}
	reamining := countRemainingCardsCore(deck)
	if reamining != 0 {
		t.Log("Excpected 50 items to remain but was " + strconv.Itoa(int(reamining)))
		t.Fail()
	}
}

func TestDrawCards_InsufficientAmount(t *testing.T) {
	deck := CreateDeck(true)
	_, err := DrawCards(&deck, 53)
	if err == nil {
		t.Log("Expected to fail")
		t.Fail()
	}

}
