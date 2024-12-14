import re

with open('input.txt', 'r', encoding='utf8') as file:
    lines = file.readlines()

pattern = re.compile(r'mul\(\d{1,3},\d{1,3}\)')

sum = 0
for line in lines:
    matches = pattern.findall(line)
    for match in matches:
        op1, op2 = match[4:-1].split(',')
        sum += int(op1) * int(op2)

print(sum)
