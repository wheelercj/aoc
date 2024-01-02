moves = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]

horizontal = 0
depth = 0
aim = 0

for move in moves:
    match move:
        case('forward', x):
            horizontal += x
            depth += aim * x
        case('up', x):
            aim -= x
        case('down', x):
            aim += x

print(f'{horizontal = }')  # 2050
print(f'{depth = }')  # 906321
print(f'{aim = }')  # 826
print(f'{horizontal * depth = }')  # 1857958050
