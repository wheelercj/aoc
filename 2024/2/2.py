def main():
    with open('input.txt', 'r', encoding='utf8') as file:
        reports = file.readlines()

    safe_report_count = 0
    for report in reports:
        level_strs: list[str] = report.strip().split(' ')
        levels: list[int] = [int(x.strip()) for x in level_strs]

        if check_safety(levels, len(levels)):
            safe_report_count += 1

    print(safe_report_count)

def check_safety(levels: list[int], original_length: int) -> bool:
    bad_levels: list[int] = []

    increasing_at = -1
    decreasing_at = -1
    for i, _ in enumerate(levels):
        if i + 1 == len(levels):
            break
        if levels[i] < levels[i+1]:
            increasing_at = i
        elif levels[i] > levels[i+1]:
            decreasing_at = i
        if increasing_at >= 0 and decreasing_at >= 0:
            bad_levels.append(increasing_at)
            bad_levels.append(increasing_at + 1)
            bad_levels.append(decreasing_at)
            bad_levels.append(decreasing_at + 1)
            break
        diff = abs(levels[i] - levels[i+1])
        if diff == 0 or diff > 3:
            bad_levels.append(i)
            bad_levels.append(i + 1)
            break

    if not bad_levels:
        return True
    if len(levels) < original_length and bad_levels:
        return False

    for bad_level in bad_levels:
        if check_safety(levels[:bad_level] + levels[bad_level+1:], original_length):
            return True

    return False

main()
