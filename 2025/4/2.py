def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    lines: list[list[str]] = [list(line) for line in s.splitlines()]

    removed: int = 0
    q: list[tuple[int, int]] = []
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if ch == "@":
                q.append((x, y))

    while q:
        coord: tuple[int, int] = q.pop(0)
        x: int = coord[0]
        y: int = coord[1]
        if lines[y][x] == "@":
            surrounding: list[tuple[int, int]] = get_surrounding_rolls(x, y, lines)
            if len(surrounding) < 4:
                removed += 1
                lines[y][x] = "."
                q.extend(surrounding)

    print(removed)


def get_surrounding_rolls(x: int, y: int, lines: list[list[str]]) -> list[tuple[int, int]]:
    coords: list[tuple[int, int]] = []

    if x > 0:
        if lines[y][x - 1] == "@":
            coords.append((x - 1, y))
    if x < len(lines[y]) - 1:
        if lines[y][x + 1] == "@":
            coords.append((x + 1, y))
    if y > 0:
        if lines[y - 1][x] == "@":
            coords.append((x, y - 1))
        if x > 0 and lines[y - 1][x - 1] == "@":
            coords.append((x - 1, y - 1))
        if x < len(lines[y]) - 1 and lines[y - 1][x + 1] == "@":
            coords.append((x + 1, y - 1))
    if y < len(lines) - 1:
        if lines[y + 1][x] == "@":
            coords.append((x, y + 1))
        if x > 0 and lines[y + 1][x - 1] == "@":
            coords.append((x - 1, y + 1))
        if x < len(lines[y]) - 1 and lines[y + 1][x + 1] == "@":
            coords.append((x + 1, y + 1))

    return coords


if __name__ == "__main__":
    main()
