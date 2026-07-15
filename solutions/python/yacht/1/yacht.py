from collections import Counter
from typing import List

# Score categories.
# Change the values as you see fit.
YACHT = (lambda dice: 50 if len(set(dice)) == 1 else 0)
ONES = (lambda dice: check_dice_for_number(dice, 1))
TWOS = (lambda dice: check_dice_for_number(dice, 2))
THREES = (lambda dice: check_dice_for_number(dice, 3))
FOURS = (lambda dice: check_dice_for_number(dice, 4))
FIVES = (lambda dice: check_dice_for_number(dice, 5))
SIXES = (lambda dice: check_dice_for_number(dice, 6))
FULL_HOUSE = (lambda dice: full_house(dice))
FOUR_OF_A_KIND = (lambda dice: four_of_kind(dice))
LITTLE_STRAIGHT = (lambda dice: 30 if sorted(dice) == [1, 2, 3, 4, 5] else 0)
BIG_STRAIGHT = (lambda dice: 30 if sorted(dice) == [2, 3, 4, 5, 6] else 0)
CHOICE = sum


def check_dice_for_number(dice: List[int], searchedNumber: int):
    return searchedNumber * dice.count(searchedNumber)


def four_of_kind(dice: List[int]):
    dice.sort()
    four_times_found_element = 0
    for number in dice:
        if (dice.count(number) >= 4):
            four_times_found_element = number
    return 4 * four_times_found_element


def full_house(dice: List[int]):
    if (sorted(Counter(dice).values()) == [2, 3]):
        return sum(dice)
    return 0


def score(dice: List[int], category):
    """checks the score of the dice

    Args:
        dice (List of Numbers): list of numbers, always 5
        category (one of the variables):
    """
    return category(dice)
