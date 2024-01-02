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


from typing import Optional
def find_winning_board(boards: list[list[list[int]]], nums: list[int]) -> Optional[list[list[int]]]:
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

marked_nums: list[int] = [nums[0]]
i = 0
while boards:
    winning_board: Optional[list[list[int]]] = find_winning_board(boards, marked_nums)
    if winning_board:
        boards.remove(winning_board)
    else:
        i += 1
        marked_nums.append(nums[i])

print(f'{marked_nums = }')
print(f'{winning_board = }')
from more_itertools import flatten
unmarked_nums: list[int] = [x for x in flatten(winning_board) if x not in marked_nums]
print(f'{unmarked_nums = }')
score: int = sum(unmarked_nums) * marked_nums[-1]
print(f'{score = }')  # 7686
