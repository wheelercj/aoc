# https://adventofcode.com/2021/day/1#part2

depths = (
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
)

from itertools import pairwise


def triplewise(iterable):
    "Return overlapping triplets from an iterable"
    # triplewise('ABCDEFG') -> ABC BCD CDE DEF EFG
    # source: https://docs.python.org/3/library/itertools.html?highlight=pairwise#itertools.pairwise
    for (a, b), (_, c) in pairwise(pairwise(iterable)):
        yield a, b, c


increased_count = 0
previous = depths[0] + depths[1] + depths[2]
for d1, d2, d3 in triplewise(depths[1:]):
    sum = d1 + d2 + d3
    if sum > previous:
        increased_count += 1
    previous = sum

print(increased_count)  # 1262
