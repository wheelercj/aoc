h_pos = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]
pos = sorted(h_pos)

def one_cost(a: int, b: int) -> int:
    distance = abs(a - b)
    cost = 0
    for i in range(distance):
        cost += i + 1
    return cost

def total_cost(positions: list[int], target: int) -> int:
    cost = 0
    for position in positions:
        cost += one_cost(position, target)
    return cost

lowest = 9999999999
first = pos[0]
last = pos[-1]
for p in range(first, last+1):
    cost = total_cost(pos, p)
    print(f'\r{cost = }         {lowest = }      ', end='')
    if cost < lowest:
        lowest = cost

# 96987874
