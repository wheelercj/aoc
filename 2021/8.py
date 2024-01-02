digits = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''

lines: list[str] = digits.split('\n')
outputs: list[str] = [line.split('|')[1] for line in lines]

unique_digit_count = 0
for output in outputs:
    for digit in output.split():
        if len(digit) in (2, 3, 4, 7):
            unique_digit_count += 1

print(f'{unique_digit_count = }')  # 543
