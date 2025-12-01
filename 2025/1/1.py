from pathlib import Path


def main():
    input_s: str = Path("input.txt").read_text(encoding="utf8").strip()

    zeros: int = 0
    pos: int = 50
    for line in input_s.splitlines():
        dir: str = line[0]
        num: int = int(line[1:])

        if num == 0:
            continue

        if dir == "L":
            pos -= num
        elif dir == "R":
            pos += num
        else:
            raise ValueError(dir)

        pos %= 100

        if pos == 0:
            zeros += 1

    print(zeros)


if __name__ == "__main__":
    main()
