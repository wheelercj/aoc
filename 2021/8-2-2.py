# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = ''
#         8                         7                 4           1    5     3     5     3

SAMPLE2 = ''
#          8       4    7   0      6      2     3     9      1  5       8       1  7   4

INPUT = '''\
'''


from tqdm import tqdm


def format_patterns(patterns: str) -> list[str]:
    return [''.join(sorted(p)) for p in patterns.strip().split()]


output_sum = 0
for row in tqdm(INPUT.splitlines()):
    input_patterns, output_patterns = row.split('|')
    input_patterns: list[str] = format_patterns(input_patterns)
    output_patterns: list[str] = format_patterns(output_patterns)

    input_patterns = sorted(input_patterns, key=len)
    digit_patterns: dict[set[str]] = {}
    digit_patterns[1] = set(input_patterns[0])
    digit_patterns[7] = set(input_patterns[1])
    digit_patterns[4] = set(input_patterns[2])
    digit_patterns[8] = set(input_patterns[-1])

    fives: list[set[str]] = []
    sixes: list[set[str]] = []
    for p in input_patterns:
        if len(p) == 5:
            fives.append(set(p))
        elif len(p) == 6:
            sixes.append(set(p))

    for p in fives:
        if digit_patterns[1] < p:
            digit_patterns[3] = p
            fives.remove(p)
            break
    if len(fives[0] - digit_patterns[4]) == 2:
        digit_patterns[5] = fives[0]
        digit_patterns[2] = fives[1]
    else:
        digit_patterns[2] = fives[0]
        digit_patterns[5] = fives[1]

    for p in sixes:
        if digit_patterns[4] < p:
            digit_patterns[9] = p
            sixes.remove(p)
            break
    if digit_patterns[7] < sixes[0]:
        digit_patterns[0] = sixes[0]
        digit_patterns[6] = sixes[1]
    else:
        digit_patterns[6] = sixes[0]
        digit_patterns[0] = sixes[1]

    digit_patterns = {''.join(sorted(v)): k for k, v in digit_patterns.items()}
    number = ''
    for op in output_patterns:
        number += str(digit_patterns[op])
    output_sum += int(number)


print(f'{output_sum = }')  # 994266
