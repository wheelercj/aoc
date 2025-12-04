def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    lines: list[str] = s.splitlines()

    accessible: int = 0
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if ch == "@":
                surrounding: int = count_surrounding_rolls(x, y, lines)
                if surrounding < 4:
                    accessible += 1

    print(accessible)


def count_surrounding_rolls(x: int, y: int, lines: list[str]) -> int:
    count: int = 0

    if x > 0:
        if lines[y][x - 1] == "@":
            count += 1
    if x < len(lines[y]) - 1:
        if lines[y][x + 1] == "@":
            count += 1
    if y > 0:
        if lines[y - 1][x] == "@":
            count += 1
        if x > 0 and lines[y - 1][x - 1] == "@":
            count += 1
        if x < len(lines[y]) - 1 and lines[y - 1][x + 1] == "@":
            count += 1
    if y < len(lines) - 1:
        if lines[y + 1][x] == "@":
            count += 1
        if x > 0 and lines[y + 1][x - 1] == "@":
            count += 1
        if x < len(lines[y]) - 1 and lines[y + 1][x + 1] == "@":
            count += 1

    return count


if __name__ == "__main__":
    main()
