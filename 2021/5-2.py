vent_lines_input: str = '''
    # The inputs here are not present in the public repo as requested
    # https://adventofcode.com/2021/about
'''

def parse_input(input_lines: str) -> list[list[list[int]]]:
    '''Returns the start and end points of each line of vents.'''
    start_and_end_points: list[list[list[int]]] = []
    for line in input_lines.split('\n'):
        if line == '':
            continue
        points: list[str] = line.split(' -> ')
        points: list[list[int]] = [[int(x) for x in point.split(',')] for point in points]
        start_and_end_points.append(points)
    return start_and_end_points

def remove_diagonals(vent_lines: list[list[list[int]]]) -> list[list[list[int]]]:
    """Removes all pairs of points that are diagonals."""
    valid_lines: list[list[list[int]]] = []
    for line in vent_lines:
        if line[0][0] == line[1][0] or line[0][1] == line[1][1]:
            valid_lines.append(line)
    return valid_lines

def get_all_points(vent_lines: list[list[list[int]]]) -> list[list[int]]:
    """Returns all the points with vents, sorted by x, then y.
    
    Duplicates indicate multiple vents at one location.
    """
    points: list[list[int]] = []
    for line in vent_lines:
        points.extend(get_points_between(line[0], line[1]))
    return sorted(points)

def get_points_between(point1: list[int], point2: list[int]) -> list[list[int]]:
    """Get all points between two points.
    
    There may be duplicates.
    """
    points: list[list[int]] = []
    if point1[0] == point2[0]:
        for y in range(min(point1[1], point2[1]), max(point1[1], point2[1]) + 1):
            points.append([point1[0], y])
    elif point1[1] == point2[1]:
        for x in range(min(point1[0], point2[0]), max(point1[0], point2[0]) + 1):
            points.append([x, point1[1]])
    else:
        # This is a diagonal line.
        if point1[0] < point2[0]:
            x_increment = 1
        else:
            x_increment = -1
        if point1[1] < point2[1]:
            y_increment = 1
        else:
            y_increment = -1
        x = point1[0]
        y = point1[1]
        while x != point2[0] or y != point2[1]:
            points.append([x, y])
            x += x_increment
            y += y_increment
        points.append([x, y])
    return points

def count_dangerous_points(sorted_points: list[list[int]]) -> int:
    """Counts the number of points with dangerous vents."""
    count: int = 0
    previous_point: list[int] = None
    previous_point2: list[int] = None
    for point in sorted_points:
        if point == previous_point == previous_point2:
            continue
        if point == previous_point:
            count += 1
            print(f'\r{count = }', end='')
        previous_point2 = previous_point
        previous_point = point
    print('\r', end='')
    return count

vent_lines: list[list[list[int]]] = parse_input(vent_lines_input)
print(f'{vent_lines = }')
vent_points: list[list[int]] = get_all_points(vent_lines)
print(f'{vent_points = }')
dangerous_points_count: int = count_dangerous_points(vent_points)
print(f'{dangerous_points_count = }')  # 22335
