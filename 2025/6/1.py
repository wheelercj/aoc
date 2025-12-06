def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    lines: list[str] = s.splitlines()

    nums: list[list[int]] = [[int(num) for num in line.split()] for line in lines[:-1]]
    ops: list[str] = lines[-1].split()

    grand_total: int = 0
    for x in range(0, len(nums[0])):
        total: int = 0
        for y in range(0, len(nums)):
            if ops[x] == "*":
                if total == 0:
                    total = 1
                total *= nums[y][x]
            else:
                total += nums[y][x]

        grand_total += total

    print(grand_total)


if __name__ == "__main__":
    main()
