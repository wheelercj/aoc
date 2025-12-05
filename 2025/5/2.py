def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    fresh_s, _ = s.split("\n\n")

    fresh: list[list[int]] = [
        [int(num) for num in line.split("-")] for line in fresh_s.splitlines()
    ]

    is_changed: bool = True
    while is_changed:
        is_changed = False

        i: int = 0
        while i < len(fresh):
            start1, end1 = fresh[i]

            is_deduped: bool = False
            for j, range2 in enumerate(fresh):
                if i == j:
                    continue
                start2, end2 = range2

                if start1 <= start2 <= end1:
                    fresh[j][0] = start1
                    is_deduped = True
                if start1 <= end2 <= end1:
                    fresh[j][1] = end1
                    is_deduped = True

            if is_deduped:
                del fresh[i]
                is_changed = True
            else:
                i += 1

    fresh_count: int = 0
    for start1, end1 in fresh:
        fresh_count += end1 - start1 + 1

    print(fresh_count)


if __name__ == "__main__":
    main()
