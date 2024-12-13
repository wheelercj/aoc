with open('input.txt', 'r', encoding='utf8') as file:
    reports = file.readlines()

safe_report_count = 0
for report in reports:
    safe = True

    increasing = False
    decreasing = False
    level_strs: list[str] = report.strip().split(' ')
    levels: list[int] = [int(x.strip()) for x in level_strs]
    for i, _ in enumerate(levels):
        if i == len(levels) - 1:
            break
        if levels[i] < levels[i+1]:
            increasing = True
        elif levels[i] > levels[i+1]:
            decreasing = True
        if increasing and decreasing:
            safe = False
            break
        diff = abs(levels[i] - levels[i+1])
        if diff == 0 or diff > 3:
            safe = False
            break

    if safe:
        safe_report_count += 1

print(safe_report_count)
