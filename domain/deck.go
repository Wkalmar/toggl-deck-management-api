package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("Create deck.", "shuffled", shuffled, "cards", cards)
	if len(cards) == 0 {
		cards = initCards()
	}
	if shuffled {
		shuffleCards(cards)
	}
	result := Deck{
		DeckId:   uuid.New(),
		Shuffled: shuffled,
		Cards:    cards,
	}
	sugar.Infow("Create deck completed", "deck", result)
	return result
}

func DrawCards(deck *Deck, count uint8) ([]Card, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("Draw cards.", "deck", deck, "count", count)
	if count > CountRemainingCards(*deck) {
		err := errors.New("insuffucient amount of cards in deck")
		sugar.Errorw("Draw cards completed.", "err", err, "deck", deck)
		return nil, err
	}
	result := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	sugar.Infow("Draw cards completed.", "result", result, "deck", deck)
	return result, nil
}
