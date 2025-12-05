def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    fresh_s, available_s = s.split("\n\n")

    fresh: list[tuple[int, ...]] = [
        tuple(int(num) for num in line.split("-")) for line in fresh_s.splitlines()
    ]

    fresh_count: int = 0
    for available in available_s.splitlines():
        if is_fresh(int(available), fresh):
            fresh_count += 1

    print(fresh_count)


def is_fresh(available: int, fresh: list[tuple[int, ...]]) -> bool:
    for start, end in fresh:
        if start <= available <= end:
            return True

    return False


if __name__ == "__main__":
    main()
