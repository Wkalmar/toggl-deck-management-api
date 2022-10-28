package domain

type Card struct {
	Rank  Rank
	Shape Shape
}

func CreateCard(rank Rank, shape Shape) Card {
	return Card{
		Rank:  rank,
		Shape: shape,
	}
}
