def equilateral(sides):
    a, b, c = sides

    return isValidTriangle(sides) and a == b == c


def isosceles(sides):
    a, b, c = sides
    return isValidTriangle(sides) and (a == b or a == c or b == c)


def scalene(sides):
    a, b, c = sides
    return isValidTriangle(sides) and (a != b and a != c and b != c)


def isValidTriangle(sides):
    return noBelowZeroNumbers(sides) and validShape(sides)


def noBelowZeroNumbers(sides):
    noBelowZeroFound = True
    for s in sides:
        if (s <= 0):
            noBelowZeroFound = False
    return noBelowZeroFound


def validShape(sides):
    a, b, c = sides
    return (a + b >= c) and (b + c >= a) and (a + c >= b)
