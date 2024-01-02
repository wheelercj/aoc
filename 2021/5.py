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
    """Returns all the points with vents.
    
    Duplicates indicate multiple vents at one location.
    """
    points: list[list[int]] = []
    for line in vent_lines:
        points.extend(get_points_between(line[0], line[1]))
    return points

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
        raise ValueError('Points are not on the same line.')
    return points

def convert_to_map(points: list[list[int]]) -> str:
    """Converts points to a map."""
    map_: list[str] = []
    previous_point: list[int] = None
    for y in range(0, 1000):
        for x in range(0, 1000):
            point: list[int] = [x, y]
            if point == previous_point:
                continue
            previous_point = point
            if point in points:
                map_.append(str(points.count(point)))
            else:
                map_.append('.')
        map_.append('\n')
    return ''.join(map_)

def count_dangerous_points(sorted_points: list[list[int]]) -> int:
    """Counts the number of dangerous points in the map."""
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
vent_lines: list[list[list[int]]] = remove_diagonals(vent_lines)
print(f'{vent_lines = }')
assert isinstance(vent_lines[0][0][0], int)

vent_points: list[list[int]] = get_all_points(vent_lines)
print(f'{vent_points = }')
assert isinstance(vent_points[0][0], int)

sorted_points: list[list[int]] = sorted(vent_points)
print(f'{sorted_points = }')

dangerous_points_count: int = count_dangerous_points(sorted_points)
print(f'{dangerous_points_count = }')  # 6397

# map_: str = convert_to_map(sorted_points)
# print(f'{map_ = }')
