with open("3/input.txt", "r") as file:
    INPUT = file.read()
alphabet = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
priorities_sum = 0
for line in INPUT.splitlines():
    first, second = line[len(line) // 2:], line[:len(line) // 2]
    for letter in first:
        if letter in second:
            priorities_sum += alphabet.index(letter) + 1
            break

print(priorities_sum)  # 7727
