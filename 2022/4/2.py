with open("4/input.txt", "r") as file:
    INPUT = file.read()
count = 0
for line in INPUT.splitlines():
    first, second = line.split(",")
    first1, first2 = first.split("-")
    second1, second2 = second.split("-")
    first1 = int(first1)
    first2 = int(first2)
    second1 = int(second1)
    second2 = int(second2)
    if (
        second1 <= first1 <= second2
        or second1 <= first2 <= second2
        or first1 <= second1 <= first2
        or first1 <= second2 <= first2
    ):
        count += 1

print(count)  # 811
