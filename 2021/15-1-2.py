# Copied from https://www.youtube.com/watch?v=8_9I6fNR7Z8
import heapq


def main():
    coords: dict[tuple[int, int], int] = parse_input(INPUT)
    last_x, last_y = max(coords)
    
    best_at: dict[tuple[int, int], int] = {}
    todo = [(0, (0, 0))]
    while todo:
        cost, last_coord = heapq.heappop(todo)
        if last_coord in best_at and cost >= best_at[last_coord]:
            continue
        else:
            best_at[last_coord] = cost
        
        if last_coord == (last_x, last_y):
            print(f'{cost = }')  # 609
            break

        for candidate in next_p(*last_coord):
            if candidate in coords:
                todo.append((cost + coords[candidate], candidate))

    print(f'{best_at[(last_x, last_y)] = }')


def parse_input(input_str: str) -> dict[tuple[int, int], int]:
    coords = {}
    for y, line in enumerate(input_str.splitlines()):
        for x, char in enumerate(line):
            coords[(x, y)] = int(char)
    return coords


def next_p(x, y):
    yield x + 1, y
    yield x, y + 1

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''


INPUT = '''\
'''


if __name__ == '__main__':
    main()
