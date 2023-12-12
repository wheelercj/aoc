SAMPLE = '''\
A Y
B X
C Z
'''

with open("2/input.txt", "r") as file:
    INPUT = file.read()
score = 0
for opponent, must_do in (line.split() for line in INPUT.splitlines()):
    if must_do == 'Y':
        score += 3
        me = opponent
    elif must_do == 'Z':
        score += 6
        me = (
            'A' if opponent == 'C'
            else 'B' if opponent == 'A'
            else 'C'
        )
    else:
        me = (
            'A' if opponent == 'B'
            else 'B' if opponent == 'C'
            else 'C'
        )
    match me:
        case 'A':
            score += 1
        case 'B':
            score += 2
        case 'C':
            score += 3

print(score)  # 10398
