SAMPLE = '''\
A Y
B X
C Z
'''

with open("2/input.txt", "r") as file:
    INPUT = file.read()
score = 0
for opponent, me in (line.split() for line in INPUT.splitlines()):
    me = 'A' if me == 'X' else 'B' if me == 'Y' else 'C'
    if me == 'A':
        score += 1
    elif me == 'B':
        score += 2
    else:
        score += 3
    
    if opponent == me:
        score += 3
    elif (
        opponent == 'A' and me == 'B'
        or opponent == 'B' and me == 'C'
        or opponent == 'C' and me == 'A'
    ):
        score += 6

print(score)  # 13009
