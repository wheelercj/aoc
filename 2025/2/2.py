import re
from pathlib import Path


def main():
    input_s: str = Path("input.txt").read_text(encoding="utf8").strip()

    silly_pattern: re.Pattern = re.compile(r"^(.+)\1+$")

    sum: int = 0
    for range_ in input_s.split(","):
        first_id, last_id = range_.split("-")
        for id in range(int(first_id), int(last_id) + 1):
            if silly_pattern.match(str(id)):
                sum += id

    print(sum)


if __name__ == "__main__":
    main()
