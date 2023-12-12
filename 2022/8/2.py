DAY = 8


def parse_input(input_text: str) -> list[list[int]]:
    lines = input_text.splitlines()
    trees = [[-1] * len(lines[0]) for _ in lines]
    for i, line in enumerate(lines):
        for j, ch in enumerate(line):
            trees[i][j] = int(ch)
    return trees


def count_visible_trees(trees: list[list[int]], y: int, x: int) -> list[int]:
    c = [0] * 4
    if x == 0 or x == len(trees[0]) - 1 or y == 0 or y == len(trees) - 1:
        return c
    i = y
    j = x
    latest_height = trees[i][j]
    i += 1
    while i < len(trees):
        if trees[i][j] < latest_height:
            c[0] += 1
        else:
            c[0] += 1
            break
        i += 1
    i = y
    latest_height = trees[i][j]
    i -= 1
    while i >= 0:
        if trees[i][j] < latest_height:
            c[1] += 1
        else:
            c[1] += 1
            break
        i -= 1
    i = y
    j = x
    latest_height = trees[i][j]
    j += 1
    while j < len(trees[0]):
        if trees[i][j] < latest_height:
            c[2] += 1
        else:
            c[2] += 1
            break
        j += 1
    j = x
    latest_height = trees[i][j]
    j -= 1
    while j >= 0:
        if trees[i][j] < latest_height:
            c[3] += 1
        else:
            c[3] += 1
            break
        j -= 1
    return c


def product(numbers: list[int]) -> int:
    p = 1
    for n in numbers:
        p *= n
    return p


def get_visible_trees_product(trees: list[list[int]]) -> int:
    best_product = 0
    for i, _ in enumerate(trees):
        for j, _ in enumerate(trees[i]):
            p = product(count_visible_trees(trees, i, j))
            if p > best_product:
                best_product = p
    return best_product


trees = parse_input(INPUT)
p = get_visible_trees_product(trees)
assert p == 8, p


with open(f"{DAY}/input.txt", "r") as file:
    INPUT = file.read()
trees = parse_input(INPUT)
print(f"{get_visible_trees_product(trees) = }")  # 334880
