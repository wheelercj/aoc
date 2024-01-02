def main():
    str_coords, str_folds = INPUT.split('\n\n')

    coords: list[list[int]] = []
    max_x = -1
    max_y = -1
    for line in str_coords.splitlines():
        coord = line.split(',')
        x, y = int(coord[0]), int(coord[1])
        coords.append([x, y])
        if x > max_x:
            max_x = x
        if y > max_y:
            max_y = y

    matrix: list[list[str]] = [
        ['.' for _ in range(max_x + 1)]
        for _ in range(max_y + 1)]

    for x, y in coords:
        matrix[y][x] = '#'

    folds: list[list[str, int]] = []
    for str_fold in str_folds.splitlines():
        axis, value = str_fold.split()[-1].split('=')
        folds.append([axis, int(value)])

    for axis, value in folds[:1]:  # This doesn't work with more than one fold.
        if axis == 'y':
            bottom_half = matrix[value+1:]
            top_half = matrix[:value]
            top_half.reverse()
            new_matrix = []
            for top_row, bottom_row in zip(top_half, bottom_half):
                new_row = []
                for top_cell, bottom_cell in zip(top_row, bottom_row):
                    if '#' in (top_cell, bottom_cell):
                        new_row.append('#')
                    else:
                        new_row.append('.')
                new_matrix.append(new_row)
            new_matrix.reverse()
        elif axis == 'x':
            left_half = [matrix[i][:value] for i, _ in enumerate(matrix)]
            for row in left_half:
                row.reverse()
            right_half = [matrix[i][value+1:] for i, _ in enumerate(matrix)]
            new_matrix = []
            for left_row, right_row in zip(left_half, right_half):
                new_row = []
                for left_cell, right_cell in zip(left_row, right_row):
                    if '#' in (left_cell, right_cell):
                        new_row.append('#')
                    else:
                        new_row.append('.')
                new_matrix.append(new_row)
        else:
            raise ValueError
        matrix = new_matrix

    dot_count = 0
    for row in matrix:
        dot_count += row.count('#')
    print(f'{dot_count = }')  # 827


def print_matrix(matrix: list[list[str]]) -> None:
    for row in matrix:
        for cell in row:
            print(cell, end='')
        print()

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

SAMPLE_OUTPUT = 0

INPUT = '''\
'''


if __name__ == '__main__':
    main()
