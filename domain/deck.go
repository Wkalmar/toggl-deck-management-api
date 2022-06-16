package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	deck_id  uuid.UUID
	shuffled bool
	cards    []Card
}

var unshuffledCards []Card

func countRemainingCards(d Deck) uint8 {
	return uint8(len(d.cards))
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
		deck_id:  uuid.New(),
		shuffled: shuffled,
		cards:    cards,
	}
}

func DrawCards(deck *Deck, count uint8) ([]Card, error) {
	if count > countRemainingCards(*deck) {
		return nil, errors.New("DrawCards: Insuffucient amount of cards in deck")
	}
	result := deck.cards[:count]
	deck.cards = deck.cards[count:]
	return result, nil
}
