from pathlib import Path


def main():
    input_s: str = Path("input.txt").read_text(encoding="utf8").strip()

    sum: int = 0
    for range_ in input_s.split(","):
        first_id, last_id = range_.split("-")
        for id in range(int(first_id), int(last_id) + 1):
            id_s = str(id)

            denominator: int = 2
            while denominator <= len(id_s):
                is_evenly_divisible: bool = len(id_s) % denominator == 0
                if not is_evenly_divisible:
                    denominator += 1
                    continue

                substr_len: int = len(id_s) // denominator

                is_silly: bool = True
                substr: str = id_s[:substr_len]
                for j in range(substr_len, len(id_s), substr_len):
                    if substr != id_s[j : j + substr_len]:
                        is_silly = False
                        break

                if is_silly:
                    sum += id
                    break

                denominator += 1

    print(sum)


if __name__ == "__main__":
    main()
