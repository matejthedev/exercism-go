package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealerCardParsed := ParseCard(dealerCard)
	var action string
	switch {
	case sum < 12:
		action = "H"
	case sum == 22:
		action = "P"
	case sum > 11 && sum < 17 && dealerCardParsed > 6:
		action = "H"
	case sum > 11 && sum < 17:
		action = "S"
	case sum > 16 && sum < 21:
		action = "S"
	case sum == 21 && dealerCardParsed > 9:
		action = "S"
	case sum == 21 && dealerCardParsed < 10:
		action = "W"
	}

	return action
}
