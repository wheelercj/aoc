from collections import Counter

with open('input.txt', 'r', encoding='utf8') as file:
    lines = file.readlines()

lefts: list[int] = []
rights: list[int] = []
for line in lines:
    left, right = line.strip().split('   ')
    lefts.append(int(left))
    rights.append(int(right))

right_counts = Counter(rights)

similarity_score = 0
for left in lefts:
    similarity_score += left * right_counts[left]

print(similarity_score)

Counter()
