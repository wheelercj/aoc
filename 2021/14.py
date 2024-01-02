from collections import Counter


def main():
    template, rules_s = INPUT.split('\n\n')
    rules = {}
    for rule_s in rules_s.splitlines():
        arg, result = rule_s.split(' -> ')
        rules[arg] = result

    for _ in range(10):
        new_template = []
        for j, _ in enumerate(template):
            pair = template[j:j+2]
            if len(pair) == 1:
                new_template[-1] += pair
            elif pair in rules:
                new_template.append(pair[0] + rules[pair])
        template = ''.join(new_template)

    print(f'{len(template) = }')
    counts = Counter(template).most_common()
    print(f'{counts[0] = }')
    print(f'{counts[-1] = }')
    print(f'{counts[0][1] - counts[-1][1] = }')  # 4517

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''

if __name__ == '__main__':
    main()
