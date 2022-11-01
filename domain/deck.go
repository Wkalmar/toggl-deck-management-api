package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	DeckId   uuid.UUID
	Shuffled bool
	Cards    []Card
}

var unshuffledCards []Card

func CountRemainingCards(d Deck) uint8 {
	return uint8(len(d.Cards))
}

func generateUnshuffledCards() []Card {
	var res []Card

	for i := Spades; i <= Hearts; i++ {
		for j := Ace; j <= King; j++ {
			res = append(res, CreateCard(j, i))
		}
	}
	return res
}

func shuffleCards(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func initCards() []Card {
	cards := make([]Card, len(unshuffledCards))
	copy(cards, unshuffledCards)
	return cards
}

func CreateDeck(shuffled bool, cards ...Card) Deck {
	if len(cards) == 0 {
		cards = initCards()
	}
	if shuffled {
		shuffleCards(cards)
	}

	return Deck{
		DeckId:   uuid.New(),
		Shuffled: shuffled,
		Cards:    cards,
	}
}

func DrawCards(deck *Deck, count uint8) ([]Card, error) {
	if count > CountRemainingCards(*deck) {
		return nil, errors.New("DrawCards: Insuffucient amount of cards in deck")
	}
	result := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	return result, nil
}
