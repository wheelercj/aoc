from tqdm import tqdm


def main():
    matrix = parse_input(INPUT)
    total_flash_count = 0
    for _ in tqdm(range(100)):
        increment_all(matrix)
        new_flashes = True
        while new_flashes:
            flash_count = flash(matrix)
            if not flash_count:
                new_flashes = False
            total_flash_count += flash_count
    print(f'{total_flash_count = }')  # 1585


def parse_input(nums: str) -> list[list[int]]:
    return [[int(x) for x in row] for row in nums.splitlines()]


def increment_all(matrix: list[list[int]]) -> None:
    for i, row in enumerate(matrix):
        for j, _ in enumerate(row):
            matrix[i][j] += 1


def increment_adjacent(x: int, y: int, matrix: list[list[int]]) -> None:
    if x - 1 >= 0:
        if matrix[x-1][y]:
            matrix[x-1][y] += 1
    try:
        if matrix[x+1][y]:
            matrix[x+1][y] += 1
    except IndexError:
        pass
    if y - 1 >= 0:
        if matrix[x][y-1]:
            matrix[x][y-1] += 1
    try:
        if matrix[x][y+1]:
            matrix[x][y+1] += 1
    except IndexError:
        pass
    if x - 1 >= 0 and y - 1 >= 0:
        if matrix[x-1][y-1]:
            matrix[x-1][y-1] += 1
    try:
        if matrix[x+1][y+1]:
            matrix[x+1][y+1] += 1
    except IndexError:
        pass
    if x - 1 >= 0:
        try:
            if matrix[x-1][y+1]:
                matrix[x-1][y+1] += 1
        except IndexError:
            pass
    if y - 1 >= 0:
        try:
            if matrix[x+1][y-1]:
                matrix[x+1][y-1] += 1
        except IndexError:
            pass


def flash(matrix: list[list[int]]) -> int:
    flash_count = 0
    for i, row in enumerate(matrix):
        for j, cell in enumerate(row):
            if cell > 9:
                flash_count += 1
                increment_adjacent(i, j, matrix)
                matrix[i][j] = 0
    return flash_count

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''


INPUT = '''\
'''


if __name__ == '__main__':
    main()
