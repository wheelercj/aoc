# S  S  S
#  A A A
#   MMM
# SAMXMAS
#   MMM
#  A A A
# S  S  S

with open('input.txt', 'r', encoding='utf8') as file:
    lines = file.readlines()

xmas_count = 0
for y, line in enumerate(lines):
    for x, ch in enumerate(line):
        if ch == 'X':
            if y - 3 >= 0:
                if x - 3 >= 0:
                    if lines[y-1][x-1] == 'M' and lines[y-2][x-2] == 'A' and lines[y-3][x-3] == 'S':
                        xmas_count += 1
                if lines[y-1][x] == 'M' and lines[y-2][x] == 'A' and lines[y-3][x] == 'S':
                    xmas_count += 1
                if x + 3 < len(line):
                    if lines[y-1][x+1] == 'M' and lines[y-2][x+2] == 'A' and lines[y-3][x+3] == 'S':
                        xmas_count += 1
            if x + 3 < len(line):
                if lines[y][x+1] == 'M' and lines[y][x+2] == 'A' and lines[y][x+3] == 'S':
                    xmas_count += 1
                if y + 3 < len(lines):
                    if lines[y+1][x+1] == 'M' and lines[y+2][x+2] == 'A' and lines[y+3][x+3] == 'S':
                        xmas_count += 1
            if y + 3 < len(lines):
                if lines[y+1][x] == 'M' and lines[y+2][x] == 'A' and lines[y+3][x] == 'S':
                    xmas_count += 1
                if x - 3 >= 0:
                    if lines[y+1][x-1] == 'M' and lines[y+2][x-2] == 'A' and lines[y+3][x-3] == 'S':
                        xmas_count += 1
            if x - 3 >= 0:
                if lines[y][x-1] == 'M' and lines[y][x-2] == 'A' and lines[y][x-3] == 'S':
                    xmas_count += 1

print(xmas_count)
