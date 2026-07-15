def valid_triangle(f):
    def inner(sides):
        return sum(sides) > 2 * max(sides) and f(sides)
    return inner


@valid_triangle
def equilateral(sides):
    a, b, c = sides

    return a == b == c


@valid_triangle
def isosceles(sides):
    a, b, c = sides
    return (a == b or a == c or b == c)


@valid_triangle
def scalene(sides):
    a, b, c = sides
    return (a != b and a != c and b != c)
