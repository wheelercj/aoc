# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''


from functools import reduce
from operator import __mul__
from tqdm import tqdm


def parse_input(input_str: str) -> list[list[int]]:
    return [list(map(int, list(line))) for line in input_str.splitlines()]


def is_adjacent(coords: tuple[int, int],
                basin_coords: list[tuple[int, int]]) -> bool:
    for x, y in basin_coords:
        if x == coords[0] and abs(y - coords[1]) == 1 \
                or y == coords[1] and abs(x - coords[0]) == 1:
            return True
    return False


def find_basins_coords(matrix: list[list[int]]) -> list[list[tuple[int, int]]]:
    basins_coords: list[list[tuple[int, int]]] = []
    for i, row in enumerate(tqdm(matrix)):
        for j, cell in enumerate(row):
            if cell == 9:
                continue
            found = False
            for k, basin_coords in enumerate(basins_coords):
                if is_adjacent((i, j), basin_coords):
                    basins_coords[k].append((i, j))
                    found = True
                    break
            if not found:
                basins_coords.append([(i, j)])

    return basins_coords


def join_adjacent_basins(basins_coords: list[list[tuple[int, int]]]) \
        -> list[list[tuple[int, int]]]:
    joined_basins_coords: list[list[tuple[int, int]]] = [basins_coords[0]]
    for basin_coords in tqdm(basins_coords[1:]):
        found = False
        for x, y in basin_coords:
            found = False
            for i, jbc in enumerate(joined_basins_coords):
                if is_adjacent((x, y), jbc):
                    joined_basins_coords[i].extend(basin_coords)
                    found = True
                    break
            if found:
                break
        if not found:
            joined_basins_coords.append(basin_coords)

    return joined_basins_coords


matrix: list[list[int]] = parse_input(INPUT)
print('Finding basin coords.')
basins_coords: list[list[tuple[int, int]]] = find_basins_coords(matrix)
print('Joining adjacent basins.')
basins_coords: list[list[tuple[int, int]]] = join_adjacent_basins(basins_coords)
print(f'{len(basins_coords) = }')
three_largest: list[list[tuple[int, int]]] = sorted(basins_coords, key=len)[-3:]
result: int = reduce(__mul__, (len(x) for x in three_largest))
print(f'{result = }')  # 916688
