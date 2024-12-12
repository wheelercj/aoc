with open('input.txt', 'r', encoding='utf8') as file:
    lines = file.readlines()

lefts: list[int] = []
rights: list[int] = []
for line in lines:
    left, right = line.strip().split('   ')
    lefts.append(int(left))
    rights.append(int(right))

lefts = sorted(lefts)
rights = sorted(rights)

sum = 0
for left, right in zip(lefts, rights):
    diff = abs(left - right)
    sum += diff

print(sum)
