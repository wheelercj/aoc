# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''


def parse_input(input_str: str) -> list[list[int]]:
    return [list(map(int, list(line))) for line in input_str.splitlines()]


def plus_sliding_window(matrix: list[list[int]]) \
        -> tuple[int, int, int, int, int]:
    for i, row in enumerate(matrix):
        for j, center in enumerate(row):
            if i - 1 >= 0:
                up = matrix[i-1][j]
            else:
                up = None
            try:
                down = matrix[i+1][j]
            except IndexError:
                down = None
            if j - 1 >= 0:
                left = matrix[i][j-1]
            else:
                left = None
            try:
                right = matrix[i][j+1]
            except IndexError:
                right = None
            
            yield up, down, left, right, center


matrix = parse_input(INPUT)
risk_levels = []
for up, down, left, right, center in plus_sliding_window(matrix):
    if up is not None and up <= center \
            or down is not None and down <= center \
            or left is not None and left <= center \
            or right is not None and right <= center:
        continue
    risk_levels.append(1 + center)

print(f'{sum(risk_levels) = }')  # 522
