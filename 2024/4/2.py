# M S    M M    S M    S S
#  A      A      A      A
# M S    S S    S M    M M

with open('input.txt', 'r', encoding='utf8') as file:
    lines = file.readlines()

x_mas_count = 0
for y, line in enumerate(lines):
    for x, ch in enumerate(line):
        if ch == 'A' and 0 < x < len(line) - 1 and 0 < y < len(lines) - 1:
            if lines[y-1][x-1] == 'M' and lines[y+1][x+1] == 'S':
                if lines[y-1][x+1] == 'M' and lines[y+1][x-1] == 'S':
                    x_mas_count += 1
                if lines[y-1][x+1] == 'S' and lines[y+1][x-1] == 'M':
                    x_mas_count += 1
            if lines[y-1][x-1] == 'S' and lines[y+1][x+1] == 'M':
                if lines[y-1][x+1] == 'M' and lines[y+1][x-1] == 'S':
                    x_mas_count += 1
                if lines[y-1][x+1] == 'S' and lines[y+1][x-1] == 'M':
                    x_mas_count += 1

print(x_mas_count)
