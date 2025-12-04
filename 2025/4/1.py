def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    lines: list[str] = s.splitlines()

    accessible: int = 0
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if ch == "@":
                neighbors: int = count_neighbor_rolls(x, y, lines)
                if neighbors < 4:
                    accessible += 1

    print(accessible)


def count_neighbor_rolls(x: int, y: int, lines: list[str]) -> int:
    count: int = 0

    for j in range(y - 1, y + 2):
        if j < 0 or j >= len(lines):
            continue

        for i in range(x - 1, x + 2):
            if i < 0 or i >= len(lines[j]):
                continue
            elif i == x and j == y:
                continue

            if lines[j][i] == "@":
                count += 1

    return count


if __name__ == "__main__":
    main()
