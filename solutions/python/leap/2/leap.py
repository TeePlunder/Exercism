def leap_year(year):
    return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)


print(f"1997 leap? => {leap_year(1997)}")
print(f"1996 leap? => {leap_year(1996)}")
print(f"1900 leap? => {leap_year(1900)}")
print(f"2000 leap? => {leap_year(2000)}")
print(f"2028 leap? => {leap_year(2028)}")
