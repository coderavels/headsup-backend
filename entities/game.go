package entities

type Deck struct {
	ID    int
	Name  string
	Cards []*Card
}

type Card struct {
	ID   int
	Name string
}
