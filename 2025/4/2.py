def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    lines: list[list[str]] = [list(line) for line in s.splitlines()]

    q: list[tuple[int, int]] = []
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if ch == "@":
                q.append((x, y))

    removed: int = 0
    while q:
        x, y = q.pop(0)
        if lines[y][x] == "@":
            neighbors: list[tuple[int, int]] = get_neighbor_rolls(x, y, lines)
            if len(neighbors) < 4:
                removed += 1
                lines[y][x] = "."
                q.extend(neighbors)

    print(removed)


def get_neighbor_rolls(x: int, y: int, lines: list[list[str]]) -> list[tuple[int, int]]:
    coords: list[tuple[int, int]] = []

    for j in range(y - 1, y + 2):
        if j < 0 or j >= len(lines):
            continue

        for i in range(x - 1, x + 2):
            if i < 0 or i >= len(lines[j]):
                continue
            elif i == x and j == y:
                continue

            if lines[j][i] == "@":
                coords.append((i, j))

    return coords


if __name__ == "__main__":
    main()
