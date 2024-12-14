import re

with open('input.txt', 'r', encoding='utf8') as file:
    lines = file.readlines()

pattern = re.compile(r"(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))")

sum = 0
enabled = True
for line in lines:
    matches = pattern.findall(line)
    for match in matches:
        if match.startswith('mul'):
            if not enabled:
                continue
            op1, op2 = match[4:-1].split(',')
            sum += int(op1) * int(op2)
        elif match.startswith("don't"):
            enabled = False
        elif match.startswith('do'):
            enabled = True
        else:
            raise NotImplementedError

print(sum)
