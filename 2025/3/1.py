def main():
    with open("input.txt", encoding="utf8") as file:
        input_s: str = file.read().strip()

    sum: int = 0
    for bank in input_s.splitlines():
        max: str = "-1"
        max_i: int = -1
        for i, num in enumerate(bank[:-1]):
            if int(num) > int(max):
                max = num
                max_i = i

        second: str = "-1"
        for num in bank[max_i + 1 :]:
            if int(num) > int(second):
                second = num

        sum += int(max + second)

    print(sum)


if __name__ == "__main__":
    main()
