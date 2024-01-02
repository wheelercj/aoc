b = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''


def is_ones_most_common(rows: list[str], column_number: int) -> bool:
    ones = 0
    zeros = 0
    for row in rows:
        if int(row[column_number]):
            ones += 1
        else:
            zeros += 1
    if ones > zeros:
        return True
    return False


gamma_rate_binary = ''
epsilon_rate_binary = ''

rows = b.split('\n')
for column_number, _ in enumerate(rows[0]):
    if is_ones_most_common(rows, column_number):
        gamma_rate_binary += '1'
        epsilon_rate_binary += '0'
    else:
        gamma_rate_binary += '0'
        epsilon_rate_binary += '1'

gamma_rate_binary = ''.join(gamma_rate_binary)
epsilon_rate_binary = ''.join(epsilon_rate_binary)


def to_decimal(num: str) -> int:
    result = 0
    for i, n in enumerate(num[::-1]):
        result += int(n) * 2 ** i
    return result


gamma_rate = to_decimal(gamma_rate_binary)
epsilon_rate = to_decimal(epsilon_rate_binary)

power_consumption = gamma_rate * epsilon_rate
print(f'{power_consumption = }')  # 3277364
