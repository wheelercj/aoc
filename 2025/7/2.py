def main():
    with open("input.txt", "r", encoding="utf8") as file:
        input_s: str = file.read().strip()

    grid: list[list[str | int]] = [list(line) for line in input_s.splitlines()]

    s_coords: tuple[int, int] | None = None
    for y, line in enumerate(grid):
        for x, ch in enumerate(line):
            if ch == "S":
                s_coords = (x, y)
                grid[y][x] = 1
    assert s_coords, s_coords

    total_timeline_count: int = 0
    q: list[tuple[int, int]] = [s_coords]
    while q:
        x, y = q.pop(0)

        timeline_count: str | int = grid[y][x]
        assert isinstance(timeline_count, int), timeline_count

        if y + 1 == len(grid):
            total_timeline_count += timeline_count
            continue

        if isinstance(grid[y + 1][x], int):
            grid[y + 1][x] += timeline_count  # type: ignore
        elif grid[y + 1][x] == ".":
            grid[y + 1][x] = timeline_count
            q.append((x, y + 1))
        elif grid[y + 1][x] == "^":
            can_go_left: bool = x > 0 and grid[y + 1][x - 1] != "^"
            can_go_right: bool = x < len(grid[y + 1]) - 1 and grid[y + 1][x + 1] != "^"
            assert can_go_left or can_go_right

            if can_go_left:
                if isinstance(grid[y + 1][x - 1], int):
                    grid[y + 1][x - 1] += timeline_count  # type: ignore
                else:
                    grid[y + 1][x - 1] = timeline_count
                    q.append((x - 1, y + 1))
            if can_go_right:
                if isinstance(grid[y + 1][x + 1], int):
                    grid[y + 1][x + 1] += timeline_count  # type: ignore
                else:
                    grid[y + 1][x + 1] = timeline_count
                    q.append((x + 1, y + 1))
        else:
            raise ValueError(grid[y + 1][x])

    print(total_timeline_count)


if __name__ == "__main__":
    main()
