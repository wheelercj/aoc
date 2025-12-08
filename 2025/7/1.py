def main():
    with open("input.txt", "r", encoding="utf8") as file:
        input_s: str = file.read().strip()

    grid: list[list[str]] = [list(line) for line in input_s.splitlines()]

    s_coords: tuple[int, int] | None = None
    for y, line in enumerate(grid):
        for x, ch in enumerate(line):
            if ch == "S":
                s_coords = (x, y)
    assert s_coords, s_coords

    split_count: int = 0
    q: list[tuple[int, int]] = [s_coords]
    while q:
        x, y = q.pop(0)
        if y + 1 == len(grid):
            continue
        elif grid[y + 1][x] == "^":
            split_count += 1
            if x - 1 >= 0:
                q.append((x - 1, y + 1))
            if x + 1 < len(grid[y + 1]):
                q.append((x + 1, y + 1))
        elif grid[y + 1][x] == ".":
            q.append((x, y + 1))
            grid[y + 1][x] = "|"
        elif grid[y + 1][x] == "|":
            continue
        else:
            raise ValueError(grid[y + 1][x])

    print(split_count)


if __name__ == "__main__":
    main()
