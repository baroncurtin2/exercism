package blackjack

var CardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return CardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	playerVal := card1Val + card2Val
	dealerCardVal := ParseCard(dealerCard)

	var action string

	switch {
	case playerVal == 22:
		action = "P"
	case playerVal == 21 && dealerCardVal != 10 && dealerCardVal != 11:
		action = "W"
	case playerVal == 21 && (dealerCardVal == 10 || dealerCardVal == 11):
		action = "S"
	case playerVal >= 17 && playerVal <= 20:
		action = "S"
	case playerVal >= 12 && playerVal <= 16 && dealerCardVal >= 7:
		action = "H"
	case playerVal >= 12 && playerVal <= 16 && dealerCardVal < 7:
		action = "S"
	case playerVal <= 11:
		action = "H"
	}

	return action
}
