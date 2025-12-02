from pathlib import Path


def main():
    input_s: str = Path("input.txt").read_text(encoding="utf8").strip()

    sum: int = 0
    for range_ in input_s.split(","):
        first_id, last_id = range_.split("-")
        for id in range(int(first_id), int(last_id) + 1):
            id_s = str(id)

            if len(id_s) % 2 == 1:
                continue

            first_half: str = id_s[: len(id_s) // 2]
            second_half: str = id_s[len(id_s) // 2 :]

            if first_half == second_half:
                sum += id

    print(sum)


if __name__ == "__main__":
    main()
