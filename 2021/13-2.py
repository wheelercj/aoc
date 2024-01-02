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

    for axis, value in folds:
        if axis == 'y':
            top_half = matrix[:value]
            top_half.reverse()
            bottom_half = matrix[value+1:]
            new_matrix = []
            longer = len(top_half) if len(top_half) > len(bottom_half) \
                else len(bottom_half)
            for i in range(longer):
                if i < len(top_half):
                    top_row = top_half[i]
                else:
                    top_row = ['.'] * len(top_half[0])
                if i < len(bottom_half):
                    bottom_row = bottom_half[i]
                else:
                    bottom_half = ['.'] * len(bottom_half[0])

                new_row = []
                assert len(top_row) == len(bottom_row)
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
            assert len(left_half) == len(right_half)
            for left_row, right_row in zip(left_half, right_half):
                new_row = []
                assert len(left_row) == len(right_row)
                for left_cell, right_cell in zip(left_row, right_row):
                    if '#' in (left_cell, right_cell):
                        new_row.append('#')
                    else:
                        new_row.append('.')
                new_matrix.append(new_row)
        else:
            raise ValueError
        matrix = new_matrix

    for row in matrix:
        row.reverse()
    print_matrix(matrix)


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
