from tqdm import tqdm

def pairs_with(char1: str, char2: str) -> bool:
    if char1 == '(' and char2 == ')' \
            or char1 == '[' and char2 == ']' \
            or char1 == '{' and char2 == '}' \
            or char1 == '<' and char2 == '>':
        return True
    return False

def get_points(char: str) -> int:
    match char:
        case ')':
            return 3
        case ']':
            return 57
        case '}':
            return 1197
        case '>':
            return 25137

def main():
    points = 0
    stack = []
    for line in tqdm(INPUT.splitlines()):
        for char in line:
            if stack and pairs_with(stack[-1], char):
                stack.pop()
            elif stack and stack[-1] in '([{<' and char in ')]}>':
                points += get_points(char)
                break
            else:
                stack.append(char)

    print(f'{points = }')  # 374061

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''

if __name__ == '__main__':
    main()
