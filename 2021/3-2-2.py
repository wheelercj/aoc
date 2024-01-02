b = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''

from copy import copy


def transpose(strings: list[str]) -> list[str]:
    return [''.join(tuple_) for tuple_ in zip(*strings)]


def get_rating(rows: list[str], is_o2: bool) -> str:
    rating_prefix = ''
    column_number = 0
    while len(rows) > 1:
        columns = transpose(rows)
        column = columns[column_number]
        if column.count('0') > column.count('1'):
            rating_prefix += '0' if is_o2 else '1'
        else:
            rating_prefix += '1' if is_o2 else '0'
        rows = [row for row in rows if row.startswith(rating_prefix)]
        column_number += 1
    return rows[0]


o2_rows = b.split()
co2_rows = copy(o2_rows)

o2_rating = get_rating(o2_rows, is_o2=True)
co2_rating = get_rating(co2_rows, is_o2=False)

life_support_rating = int(o2_rating, base=2) * int(co2_rating, base=2)
print(f'{life_support_rating = }')  # 5736383
