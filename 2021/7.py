h_pos = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]
pos = sorted(h_pos)

def cost(positions: list[int], x: int) -> int:
    cost = 0
    for position in positions:
        cost += abs(position - x)
    return cost

mid = int(len(pos)/2)
lowest_cost = cost(pos, pos[mid])
if lowest_cost < cost(pos, pos[mid + 1]):
    for p in pos[mid - 1::-1]:
        if cost(pos, p) < lowest_cost:
            lowest_cost = cost(pos, p)
        else:
            break
else:
    for p in pos[mid + 1:]:
        if cost(pos, p) < lowest_cost:
            lowest_cost = cost(pos, p)
        else:
            break

print(lowest_cost)  # 340987
