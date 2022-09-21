package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	mapCards := map[string]int{
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

	return mapCards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var response string
	var dealerCardValue int = ParseCard(dealerCard)
	var myCardSum int = ParseCard(card1) + ParseCard(card2)

	switch {
	case myCardSum == 22:
		response = "P"
	case (myCardSum == 21) && (dealerCardValue >= 10):
		response = "S"
	case (myCardSum == 21) && (dealerCardValue < 10):
		response = "W"
	case (myCardSum >= 17 && myCardSum <= 20):
		response = "S"
	case (myCardSum >= 12 && myCardSum <= 16) && dealerCardValue < 7:
		response = "S"
	case (myCardSum >= 12 && myCardSum <= 16) && dealerCardValue >= 7:
		response = "H"
	case myCardSum <= 11:
		response = "H"
	}
	return response
}
