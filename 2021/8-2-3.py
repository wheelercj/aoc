# some ideas here from https://www.youtube.com/watch?v=WDFh2jdUYlw
from tqdm import tqdm


def format_patterns(patterns: str) -> list[str]:
    return [''.join(sorted(p)) for p in patterns.strip().split()]


def main():
    output_sum = 0
    for row in tqdm(INPUT.splitlines()):
        input_patterns, output_patterns = row.split('|')
        input_patterns: list[str] = format_patterns(input_patterns)
        output_patterns: list[str] = format_patterns(output_patterns)

        digit_patterns: dict[set[str]] = {}
        digit_patterns[1], = (p for p in input_patterns if len(p) == 2)
        digit_patterns[7], = (p for p in input_patterns if len(p) == 3)
        digit_patterns[4], = (p for p in input_patterns if len(p) == 4)
        digit_patterns[8], = (p for p in input_patterns if len(p) == 7)
        
        len5 = {p for p in input_patterns if len(p) == 5}
        len6 = {p for p in input_patterns if len(p) == 6}

        digit_patterns[6], = (p for p in len6
                                if len(set(p) & set(digit_patterns[1])) == 1)
        digit_patterns[9], = (p for p in len6
                                if len(set(p) & set(digit_patterns[4])) == 4)
        digit_patterns[0], = len6 - {digit_patterns[9], digit_patterns[6]}

        digit_patterns[3], = (p for p in len5
                                if len(set(p) & set(digit_patterns[1])) == 2)
        digit_patterns[5], = (p for p in len5
                                if len(set(p) & set(digit_patterns[6])) == 5)
        digit_patterns[2], = len5 - {digit_patterns[5], digit_patterns[3]}

        digit_patterns = {v: k for k, v in digit_patterns.items()}
        number = ''
        for op in output_patterns:
            number += str(digit_patterns[op])
        output_sum += int(number)


    print(f'{output_sum = }')  # 994266


# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = ''
#         8                         7                 4           1    5     3     5     3

SAMPLE2 = ''
#          8       4    7   0      6      2     3     9      1  5       8       1  7   4

INPUT = '''\
'''


if __name__ == '__main__':
    main()
