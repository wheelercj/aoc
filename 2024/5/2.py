with open('input.txt', 'r', encoding='utf8') as file:
    content = file.read().strip()

rules_s, pages_s = content.split('\n\n')
rules_lines = rules_s.strip().split('\n')
pages_lines = pages_s.strip().split('\n')

page_groups: list[list[int]] = [
    [int(num) for num in line.split(',')] for line in pages_lines
]

rules: dict[int, set[int]] = dict()
for rule_s in rules_lines:
    left_s, right_s = rule_s.split('|')
    left, right = int(left_s), int(right_s)
    if left in rules:
        rules[left].add(right)
    else:
        rules[left] = { right }

sum: int = 0
for page_group in page_groups:
    correct_order: bool = True

    i: int = 0
    while i < len(page_group) - 1:
        page1: int = page_group[i]
        swapped: bool = False

        j: int = i + 1
        while j < len(page_group):
            page2: int = page_group[j]
            if page2 in rules and page1 in rules[page2]:
                correct_order = False
                # swap them
                page_group[i] = page2
                page_group[j] = page1
                swapped = True
                # check if the new page_group[i] should be swapped with another page
                break
            j += 1
        if swapped:
            continue
        i += 1
    if not correct_order:
        sum += page_group[len(page_group)//2]

print(sum)
