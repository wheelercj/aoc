# https://adventofcode.com/2021/day/1

depths = (
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
)

increased_count = 0
previous = depths[0]
for d in depths[1:]:
    if d > previous:
        increased_count += 1
    previous = d

print(increased_count)  # 1292
