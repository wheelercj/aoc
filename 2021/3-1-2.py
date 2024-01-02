b = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''

gamma_rate = ''
epsilon_rate = ''
for column in zip(*b.split()):
    if column.count('0') > column.count('1'):
        gamma_rate += '0'
        epsilon_rate += '1'
    else:
        gamma_rate += '1'
        epsilon_rate += '0'
power_consumption = int(gamma_rate, base=2) * int(epsilon_rate, base=2)
print(f'{power_consumption = }')  # 3277364
