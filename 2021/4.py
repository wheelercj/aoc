nums = [
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
]

boards = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''


def convert_to_int(board: str) -> list[list[int]]:
    rows = board.split('\n')
    return [[int(x) for x in row.strip().replace('  ', ' ').split(' ')] for row in rows]


def find_winning_board(boards: list[list[list[int]]], nums: list[int]) -> list[list[int]]:
    for board in boards:
        for row in board:
            is_winner = True
            for r in row:
                if r not in nums:
                    is_winner = False
                    break
            if is_winner:
                return board
        for col in range(len(board[0])):
            is_winner = True
            for row in board:
                if row[col] not in nums:
                    is_winner = False
                    break
            if is_winner:
                return board


boards: list[str] = boards.split('\n\n')
boards: list[list[list[int]]] = [convert_to_int(board) for board in boards]

j = 0
for i in range(len(nums)):
    j = i
    winning_board = find_winning_board(boards, nums[:i+1])
    if winning_board:
        break

winning_nums: list[int] = nums[:j+1]
print(f'{winning_nums = }')
from more_itertools import flatten
unmarked_nums: list[int] = [x for x in flatten(winning_board) if x not in winning_nums]
print(f'{[x for x in flatten(winning_board)] = }')
print(f'{unmarked_nums = }')
unmarked_sum: int = sum(unmarked_nums)
print(f'{unmarked_sum = }')
score: int = unmarked_sum * nums[j]
print(f'{score = }')  # 34506
