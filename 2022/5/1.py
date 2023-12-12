with open("5/input.txt", "r") as file:
    INPUT = file.read()
boxes_input, instruction_input = INPUT.split("\n\n")
box_lines = boxes_input.splitlines()
column_count = len(box_lines[-1].replace(" ", ""))
box_lines = box_lines[:-1]
columns = [[] for _ in range(column_count)]
for i, line in enumerate(reversed(box_lines)):
    for _ in range(column_count):
        line = line.replace("]    ", "] [ ]")
    line = line.replace("   ", "[ ]")
    b = line.split("] [")
    b[0] = b[0][1]
    b[-1] = b[-1][:-1]
    for j, b_ in enumerate(b):
        if b_ != " ":
            columns[j].append(b_)

instruction_lines = instruction_input.splitlines()
instructions = []
for line in instruction_lines:
    line = line.replace("move ", "").replace(" from", "").replace(" to", "")
    instructions.append(line.split(' '))

for inst in instructions:
    inst[0] = int(inst[0])
    inst[1] = int(inst[1])
    inst[2] = int(inst[2])
    for _ in range(inst[0]):
        columns[inst[2] - 1].append(columns[inst[1] - 1].pop())

for col in columns:
    print(end=col[-1])
# VCTFTJQCG
