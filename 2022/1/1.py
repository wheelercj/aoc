with open ("1/input.txt", "r") as file:
    INPUT = file.read()
elves: list[str] = INPUT.split("\n\n")
elves: list[list[int]] = [[int(x) for x in elf.splitlines()] for elf in elves]
calories: list[int] = [sum(elf) for elf in elves]
calories.sort()
print(calories[-1])  # 69836
