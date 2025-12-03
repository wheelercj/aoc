def main():
    with open("input.txt", encoding="utf8") as file:
        input_s: str = file.read().strip()

    sum: int = 0
    for bank in input_s.splitlines():
        digits: list[str] = []
        prev_max_i: int = -1
        for _ in range(12):
            max: int = -1
            max_i: int = -1

            i: int = prev_max_i + 1
            while i < len(bank) - (11 - len(digits)):
                num: int = int(bank[i])
                if num > max:
                    max = num
                    max_i = i
                i += 1
            assert max >= 0, max

            prev_max_i = max_i
            digits.append(str(max))

        sum += int("".join(digits))

    print(sum)


if __name__ == "__main__":
    main()
