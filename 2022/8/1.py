DAY = 8


def parse_input(input_text: str) -> list[list[int]]:
    lines = input_text.splitlines()
    trees = [[-1] * len(lines[0]) for _ in lines]
    for i, line in enumerate(lines):
        for j, ch in enumerate(line):
            trees[i][j] = int(ch)
    return trees


def reverse_laterally(trees: list[list[int]]) -> None:
    for i, _ in enumerate(trees):
        trees[i].reverse()


def count_visible_trees(input_text: str) -> int:
    trees = parse_input(input_text)
    trees_counted = [[False] * len(trees[0]) for _ in trees]
    visible_trees = _count_visible_trees_from_West(trees, trees_counted)
    reverse_laterally(trees)
    reverse_laterally(trees_counted)
    visible_trees += _count_visible_trees_from_West(trees, trees_counted)
    visible_trees += _count_visible_trees_from_North(trees, trees_counted)
    trees.reverse()
    trees_counted.reverse()
    visible_trees += _count_visible_trees_from_North(trees, trees_counted)
    if len(trees) < 10:
        reverse_laterally(trees_counted)
        trees_counted.reverse()
        print(f"{trees_counted = }")
    return visible_trees


def _count_visible_trees_from_West(trees: list[list[int]], trees_counted: list[list[bool]]) -> int:
    visible_trees = 0
    for i, row in enumerate(trees):
        latest_height = row[0]
        if not trees_counted[i][0]:
            visible_trees += 1
            trees_counted[i][0] = True
        for j, tree_height in enumerate(row):
            if tree_height > latest_height:
                latest_height = tree_height
                if not trees_counted[i][j]:
                    visible_trees += 1
                    trees_counted[i][j] = True
    return visible_trees


def _count_visible_trees_from_North(trees: list[list[int]], trees_counted: list[list[bool]]) -> int:
    visible_trees = 0
    for j, _ in enumerate(trees[0]):
        if j == 0 or j == len(trees[0]) - 1:
            continue
        latest_height = trees[0][j]
        if not trees_counted[0][j]:
            visible_trees += 1
            trees_counted[0][j] = True
        for i, _ in enumerate(trees):
            if i == 0 or i == len(trees) - 1:
                continue
            if trees[i][j] > latest_height:
                latest_height = trees[i][j]
                if not trees_counted[i][j]:
                    visible_trees += 1
                    trees_counted[i][j] = True
    return visible_trees


result = count_visible_trees(INPUT)
assert result == 21, result

with open(f"{DAY}/input.txt", "r") as file:
    INPUT = file.read()
print(f"{count_visible_trees(INPUT) = }")  # 1792
