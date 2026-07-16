package blackjack

// | card  | value | card    | value |
// | :---: | :---: | :-----: | :---: |
// |  ace  |  11   | eight   |   8   |
// |  two  |   2   | nine    |   9   |
// | three |   3   |  ten    |  10   |
// | four  |   4   | jack    |  10   |
// | five  |   5   | queen   |  10   |
// |  six  |   6   | king    |  10   |
// | seven |   7   | *other* |   0   |

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "ten", "jack", "queen", "king":
		value = 10
	case "nine":
		value = 9
	case "eight":
		value = 8
	case "seven":
		value = 7
	case "six":
		value = 6
	case "five":
		value = 5
	case "four":
		value = 4
	case "three":
		value = 3
	case "two":
		value = 2
	default:
		value = 0
	}

	return value
}

// - If you have a pair of aces you must always split them.
// - If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a face card (Jack/Queen/King) or a ten then you automatically win.
//   If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
// - If your cards sum up to a value within the range [17, 20] you should always stand.
// - If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
// - If your cards sum up to 11 or lower you should always hit.

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	stand := "S"
	hit := "H"
	split := "P"
	automaticallyWin := "W"

	parsedCard1 := ParseCard(card1)
	parsedCard2 := ParseCard(card2)
	parsedCardSum := parsedCard1 + parsedCard2
	parsedDealerCard := ParseCard(dealerCard)

	if parsedCardSum == 22 {
		return split
	}

	if parsedCard1+parsedCard2 == 21 {
		if parsedDealerCard == 11 || parsedDealerCard == 10 {
			return stand
		}
		return automaticallyWin
	}

	if 17 <= parsedCardSum && parsedCardSum <= 20 {
		return stand
	}

	if 12 <= parsedCardSum && parsedCardSum <= 16 {
		if parsedDealerCard >= 7 {
			return hit
		}

		return stand
	}

	if parsedCardSum <= 11 {
		return hit
	}

	return "0"
}
