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
            for _ in range(num):
                pos -= 1
                if pos == 0:
                    zeros += 1
                elif pos == -1:
                    pos = 99
        elif dir == "R":
            for _ in range(num):
                pos += 1
                if pos == 100:
                    pos = 0
                    zeros += 1
        else:
            raise ValueError(dir)

    print(zeros)


if __name__ == "__main__":
    main()
