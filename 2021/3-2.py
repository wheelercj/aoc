b = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''

from typing import Optional
from copy import copy


def is_ones_most_common(rows: list[str], column_number: int) -> Optional[bool]:
    ones = 0
    zeros = 0
    for row in rows:
        if row[column_number] == '1':
            ones += 1
        else:
            zeros += 1
    if ones > zeros:
        return True
    elif ones < zeros:
        return False
    else:
        return None


def remove_rows_starting_without(subrow: str, rows: list[str]) -> list[str]:
    i = 0
    while i < len(rows):
        row = rows[i]
        if not row.startswith(subrow):
            rows.remove(row)
            if len(rows) == 1:
                return rows
        else:
            i += 1
    return rows


def remove_wrong_rows(rows: list[str], is_o2: bool) -> str:
    target_row = ''
    column_number = 0
    while len(rows) > 1:
        match is_ones_most_common(rows, column_number):
            case None:
                target_row += '1' if is_o2 else '0'
                print('equal')
            case True:
                target_row += '1' if is_o2 else '0'
            case False:
                target_row += '0' if is_o2 else '1'
        rows = remove_rows_starting_without(target_row, rows)
        column_number += 1
    return rows[0]


def to_decimal(num: str) -> int:
    result = 0
    for i, n in enumerate(num[::-1]):
        result += int(n) * 2 ** i
    return result


o2_rows = b.split('\n')
co2_rows = copy(o2_rows)

o2_rating = remove_wrong_rows(o2_rows, is_o2=True)
print(f'{o2_rating = }')
co2_rating = remove_wrong_rows(co2_rows, is_o2=False)
print(f'{co2_rating = }')

o2_rating = to_decimal(o2_rating)
co2_rating = to_decimal(co2_rating)

life_support_rating = o2_rating * co2_rating
print(f'{life_support_rating = }')  # 5736383
