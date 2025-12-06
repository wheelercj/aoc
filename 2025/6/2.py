def main():
    with open("input.txt", "r", encoding="utf8") as file:
        s: str = file.read().strip()
    lines: list[str] = s.splitlines()

    nums: list[list[int]] = []
    column: list[int] = []
    for x in range(len(lines[0])):
        is_empty_column: bool = True

        digits: list[str] = []
        for line in lines[:-1]:
            ch: str = line[x]
            if ch != " ":
                is_empty_column = False
                digits.append(ch)

        if is_empty_column:
            nums.append(column)
            column = []
        else:
            column.append(int("".join(digits)))
    nums.append(column)

    ops: list[str] = lines[-1].split()

    grand_total: int = 0
    for x, column in enumerate(nums):
        if ops[x] == "*":
            product: int = 1
            for num in column:
                product *= num
            grand_total += product
        elif ops[x] == "+":
            grand_total += sum(column)
        else:
            raise ValueError(ops[x])

    print(grand_total)


if __name__ == "__main__":
    main()
