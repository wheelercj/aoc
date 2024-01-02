# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE = ''
#         8                         7                 4           1    5     3     5     3

SAMPLE2 = ''
#          8       4    7   0      6      2     3     9      1  5       8       1  7   4

INPUT = '''\
'''


from tqdm import tqdm


def decode_patterns(
        input_patterns: list[str],
        _digit_patterns: dict[int, set[str]]) -> dict[int, set[str]]:
    """Attempts to get the segments for the digits.
    
    It is possible not all of the segments will be found, and it is 
    possible that calling this function again with its previous output 
    will find more digits. However, if the function does not change 
    digit_patterns, no more can be found. If digit_patterns is provided,
    it will be assumed to be correct and will be used to help find the 
    remaining digit patterns.
    """
    digit_patterns = dict(_digit_patterns)
    for pattern in input_patterns:
        pattern = set(pattern)
        if pattern in digit_patterns.values():
            continue
        match len(pattern):
            case 2:
                digit_patterns[1] = pattern
                continue
            case 3:
                digit_patterns[7] = pattern
                continue
            case 4:
                digit_patterns[4] = pattern
                continue
            case 7:
                digit_patterns[8] = pattern
                continue
            case 5:
                if 7 in digit_patterns and pattern > digit_patterns[7]:
                    digit_patterns[3] = pattern
                    continue
                if 1 in digit_patterns and pattern > digit_patterns[1]:
                    digit_patterns[3] = pattern
                    continue
                if 4 in digit_patterns \
                        and len(pattern - digit_patterns[4]) == 2:
                    digit_patterns[5] = pattern
                    continue
            case 6:
                if 7 in digit_patterns:
                    if pattern > digit_patterns[7]:
                        if 4 in digit_patterns:
                            if pattern > digit_patterns[4]:
                                digit_patterns[9] = pattern
                                continue
                            else:
                                digit_patterns[0] = pattern
                                continue
                    else:
                        digit_patterns[6] = pattern
                        continue

    if len(digit_patterns) == 9:
        missing_digit: int = None
        for digit in range(10):
            if digit not in digit_patterns:
                missing_digit = digit
                break
        for pattern in input_patterns:
            if set(pattern) not in digit_patterns.values():
                digit_patterns[missing_digit] = set(pattern)
                break

    return digit_patterns


def map_segments(input_patterns: list[str],
                 segments: dict[str, str]) -> dict[str, str]:
    """Maps incorrect segments to the correct ones.
    
    The correct segments are:
     aaaa
    f    b
    f    b
     gggg
    e    c
    e    c
     dddd
    """
    digit_patterns: dict[int, set[str]] = {}
    changed = True
    while changed:
        new_digit_patterns = decode_patterns(input_patterns, digit_patterns)
        if new_digit_patterns == digit_patterns:
            changed = False
        else:
            digit_patterns = new_digit_patterns

    a = None
    ade = None
    b = None
    c = None
    d = None
    de = None
    e = None
    f = None
    fg = None
    g = None

    if 8 in digit_patterns and 4 in digit_patterns:
        ade = digit_patterns[8] - digit_patterns[4]
    if 8 in digit_patterns and 9 in digit_patterns:
        e = digit_patterns[8] - digit_patterns[9]
    if 7 in digit_patterns and 1 in digit_patterns:
        a = (digit_patterns[7] - digit_patterns[1]).pop()
    elif ade and de:
        a = (ade - de).pop()
    elif 9 in digit_patterns and 4 in digit_patterns:
        a = (digit_patterns[9] - digit_patterns[4]).pop()
    if 4 in digit_patterns and 7 in digit_patterns:
        fg = digit_patterns[4] - digit_patterns[7]
    elif 9 in digit_patterns and 7 in digit_patterns:
        fg = digit_patterns[9] - digit_patterns[7]
    elif 4 in digit_patterns and 1 in digit_patterns:
        fg = digit_patterns[4] - digit_patterns[1]
    if 8 in digit_patterns and 0 in digit_patterns:
        g = (digit_patterns[8] - digit_patterns[0]).pop()
        if fg:
            f = (fg - set(g)).pop()

    if 8 in digit_patterns and 0 in digit_patterns:
        g = (digit_patterns[8] - digit_patterns[0]).pop()
    if 4 in digit_patterns and 1 in digit_patterns and g:
        f = (digit_patterns[4] - digit_patterns[1] - set(g)).pop()
    if 3 in digit_patterns and 7 in digit_patterns and g:
        d = (digit_patterns[3] - digit_patterns[7] - set(g)).pop()
    if 5 in digit_patterns and a and f and g and d:
        c = (digit_patterns[5] - set(a) - set(f) - set(g) - set(d)).pop()
    if 1 in digit_patterns and c:
        b = (digit_patterns[1] - set(c)).pop()
    if 8 in digit_patterns and 3 in digit_patterns and f:
        e = (digit_patterns[8] - digit_patterns[3] - set(f)).pop()
    
    if a:
        segments[a] = 'a'
    if b:
        segments[b] = 'b'
    if c:
        segments[c] = 'c'
    if d:
        segments[d] = 'd'
    if e:
        segments[e] = 'e'
    if f:
        segments[f] = 'f'
    if g:
        segments[g] = 'g'

    return segments


def correct_patterns(output_patterns: list[str], segments: dict[str, str]) -> list[str]:
    """Converts the patterns to the correct segments."""
    corrected_outputs: list[str] = []
    for pattern in output_patterns:
        corrected_output = []
        for segment in pattern:
            if segment in segments:
                corrected_output.append(segments[segment])
        corrected_outputs.append(''.join(corrected_output))
    return corrected_outputs


def convert_to_digits(corrected_outputs: list[str]) -> list[int]:
    """Converts corrected output segments to their corresponding digits."""
    digits: dict[str, int] = {
        'abcdef': 0,
        'bc': 1,
        'abdeg': 2,
        'abcdg': 3,
        'bcfg': 4,
        'acdfg': 5,
        'acdefg': 6,
        'abc': 7,
        'abcdefg': 8,
        'abcdfg': 9}
    return [digits[''.join(sorted(output))] for output in corrected_outputs]


def add_missing_segment(segments: dict[str, str]) -> dict[str, str]:
    """Adds the missing segment to the segments."""
    missing_value = None
    for char in 'abcdefg':
        if char not in segments.values():
            missing_value = char
            break
    for char in 'abcdefg':
        if char not in segments.keys():
            segments[char] = missing_value
            break
    return segments


entries: list[str] = INPUT.split('\n')
output_sum: int = 0
for entry in tqdm(entries):
    split_entry = entry.split('|')
    output_patterns: list[str] = split_entry[-1].strip().split()
    input_patterns: list[str] = split_entry[0].strip().split()
    input_patterns.extend(output_patterns)

    segments: dict[str, str] = {}
    while True:
        new_segments = map_segments(input_patterns, segments)
        if new_segments == segments:
            break
        segments = new_segments
    if len(segments) == 6:
        segments = add_missing_segment(segments)
    corrected_outputs: list[str] = correct_patterns(output_patterns, segments)
    digits: list[int] = convert_to_digits(corrected_outputs)
    number: int = int(''.join(map(str, digits)))
    output_sum += number

print(f'{output_sum = }')  # 994266
