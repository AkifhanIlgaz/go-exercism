package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int

	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "king", "queen":
		value = 10
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myHand := ParseCard(card1) + ParseCard(card2)
	dealerHand := ParseCard(dealerCard)

	var decision string

	switch {
	case myHand > 21:
		decision = "P"
	case myHand == 21:
		if dealerHand < 10 {
			decision = "W"
		} else {
			decision = "S"
		}
	case myHand >= 17 && myHand <= 20:
		decision = "S"
	case myHand >= 12 && myHand <= 16:
		if dealerHand >= 7 {
			decision = "H"
		} else {
			decision = "S"
		}
	case myHand <= 11:
		decision = "H"
	}

	return decision
}
