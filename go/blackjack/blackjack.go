package blackjack

const (
	optionStand   = "S"
	optionHit     = "H"
	optionSplit   = "P"
	optionAutoWin = "W"
)

var cardValues = map[string]int{
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
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	parsedCard1 := ParseCard(card1)
	parsedCard2 := ParseCard(card2)
	parsedDealerCard := ParseCard(dealerCard)

	cardsSum := parsedCard1 + parsedCard2

	switch {
	case cardsSum == 22:
		return optionSplit
	case cardsSum == 21:
		if parsedDealerCard == 11 || parsedDealerCard == 10 {
			return optionStand
		} else {
			return optionAutoWin
		}
	case cardsSum >= 17 && cardsSum <= 20:
		return optionStand
	case cardsSum >= 12 && cardsSum <= 26:
		if parsedDealerCard >= 7 {
			return optionHit
		}

		return optionStand
	default:
		return optionHit
	}
}
