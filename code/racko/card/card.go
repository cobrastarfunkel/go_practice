package card

type Card struct {
	Next, Previous *Card
	Value          int
}
