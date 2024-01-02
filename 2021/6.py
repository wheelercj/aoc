
fish = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]

new_fish = 0
for j in range(80):
    print(f'\rday {j}', end='')
    if new_fish:
        fish.extend([8] * new_fish)
        new_fish = 0
    for i in range(len(fish)):
        fish[i] -= 1
        if fish[i] < 0:
            fish[i] = 6
            new_fish += 1

fish.extend([8] * new_fish)

print(f'\n{len(fish) = }')  # 346063
