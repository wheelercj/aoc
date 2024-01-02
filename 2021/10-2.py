def pairs_with(char1: str, char2: str) -> bool:
    if char1 == '(' and char2 == ')' \
            or char1 == '[' and char2 == ']' \
            or char1 == '{' and char2 == '}' \
            or char1 == '<' and char2 == '>':
        return True
    return False

def get_pair(char: str) -> str:
    match char:
        case '(':
            return ')'
        case '[':
            return ']'
        case '{':
            return '}'
        case '<':
            return '>'

def get_points(closers: str) -> int:
    points = 0
    for char in closers:
        points *= 5
        match char:
            case ')':
                points += 1
            case ']':
                points += 2
            case '}':
                points += 3
            case '>':
                points += 4
    return points

def complete_line(openers: list[str]) -> str:
    closers = []
    for char in openers:
        closers.append(get_pair(char))
    closers.reverse()
    return ''.join(closers)

def main():
    incomplete_lines = []
    stack = []
    for line in INPUT.splitlines():
        corrupted_line = False
        for char in line:
            if stack and pairs_with(stack[-1], char):
                stack.pop()
            elif stack and stack[-1] in '([{<' and char in ')]}>':
                corrupted_line = True
                break
            else:
                stack.append(char)
        if not corrupted_line:
            incomplete_lines.append(line)

    scores: list[int] = []
    stack: list[str] = []
    for i, line in enumerate(incomplete_lines):
        for char in line:
            if stack and pairs_with(stack[-1], char):
                stack.pop()
            else:
                stack.append(char)
        line_completer = complete_line(stack)
        scores.append(get_points(line_completer))
        incomplete_lines[i] += line_completer
        stack = []

    scores = sorted(scores)
    print(f'{scores[len(scores) // 2] = }')  # 2116639949

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = '''\
'''

INPUT = '''\
'''

if __name__ == '__main__':
    main()
