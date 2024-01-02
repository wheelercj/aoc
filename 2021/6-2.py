import collections

fish = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]
fish = collections.Counter(fish)
for i in range(9):
    if i not in fish:
        fish[i] = 0
print(f'{fish = }')

for j in range(256):
    print(f'\rday {j+1}', end='')
    old_fish = fish[0]
    for i in range(8):
        fish[i] = fish[i+1]
    fish[6] += old_fish
    fish[8] = old_fish
print()

print(f'{fish = }')
print(f'{sum(fish.values()) = }')  # 1572358335990
