from more_itertools import pairwise
from collections import Counter
from tqdm import tqdm


def main():
    template, rules_s = INPUT.split('\n\n')
    rules = {}
    for rule_s in rules_s.splitlines():
        arg, result = rule_s.split(' -> ')
        rules[(arg[0], arg[1])] = result

    counts = Counter(list(pairwise(template)))

    for _ in tqdm(range(40)):
        new_counts = Counter()
        for pair in counts:
            mid = rules[pair]
            left, right = pair[0], pair[1]
            new_counts[(left, mid)] += counts[pair]
            new_counts[(mid, right)] += counts[pair]
        counts = new_counts

    char_counts = Counter()
    for pair in counts:
        char_counts[pair[0]] += counts[pair]
    char_counts[template[-1]] += 1

    print(f'{char_counts = }')
    char_counts = char_counts.most_common()
    answer = char_counts[0][1] - char_counts[-1][1]
    print(f'{answer = }')  # 4704817645083

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''

if __name__ == '__main__':
    main()
