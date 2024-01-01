def solve(inp: str) -> int:
    DIFF_CHARS_REQ = 4
    diff_chars = []
    for i, ch in enumerate(inp):
        if ch in diff_chars:
            j = diff_chars.index(ch)
            diff_chars = diff_chars[j + 1:]
            diff_chars.append(ch)
        else:
            diff_chars.append(ch)
        if len(diff_chars) == DIFF_CHARS_REQ:
            return i + 1


# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2022/about
result = solve("redacted")
assert result == 7, result
result = solve("redacted")
assert result == 5, result
result = solve("redacted")
assert result == 6, result
result = solve("redacted")
assert result == 10, result
result = solve("redacted")
assert result == 11, result

with open("6/input.txt", "r") as file:
    INPUT = file.read()
print(solve(INPUT))  # 1848
