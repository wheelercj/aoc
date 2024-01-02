moves = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]

horizontal = 0
depth = 0

for move in moves:
    match move:
        case('forward', x):
            horizontal += x
        case('up', x):
            depth -= x
        case('down', x):
            depth += x

print(f'{horizontal = }')  # 2050
print(f'{depth = }')  # 826
print(f'{horizontal * depth = }')  # 1693300
