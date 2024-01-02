# shortest_path: list = []
lowest_cost = 99999999999999


def main():
    global lowest_cost
    matrix = [[int(x) for x in list(line)] for line in INPUT.splitlines()]
    explore(0, 0, matrix, [], 0, 0, 0)
    lowest_cost -= matrix[0][0]
    print(f'\n{lowest_cost = }')
    # print(f'{shortest_path = }')


def explore(x: int,
            y: int,
            matrix: list[list[int]],
            this_path: list[tuple[int, int]],
            this_cost: int,
            backtracks: int,
            backtrack_limit: int) -> None:
    # global shortest_path
    global lowest_cost
    this_path = list(this_path)
    print(f'\rlen: {len(this_path)}', end='')

    this_path.append((x, y))
    this_cost += matrix[y][x]

    if this_cost > lowest_cost:
        return
    
    width = len(matrix[0])
    height = len(matrix)
    if x == width - 1 and y == height - 1:  # if at end
        if this_cost < lowest_cost:
            lowest_cost = this_cost
            print(f'\n{lowest_cost = }')
            # shortest_path = list(this_path)
        return

    if y + 1 < height and (x, y + 1) not in this_path:
        explore(x,
                y + 1,
                matrix,
                this_path,
                this_cost,
                backtracks,
                backtrack_limit)
    if x + 1 < width and (x + 1, y) not in this_path:
        explore(x + 1,
                y,
                matrix,
                this_path,
                this_cost,
                backtracks,
                backtrack_limit)
    if backtracks < backtrack_limit:
        if x - 1 >= 0 and (x - 1, y) not in this_path:
            backtracks += 1
            explore(x - 1,
                    y,
                    matrix,
                    this_path,
                    this_cost,
                    backtracks,
                    backtrack_limit)
        if y - 1 >= 0 and (x, y - 1) not in this_path:
            backtracks += 1
            explore(x,
                    y - 1,
                    matrix,
                    this_path,
                    this_cost,
                    backtracks,
                    backtrack_limit)

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''


if __name__ == '__main__':
    main()
