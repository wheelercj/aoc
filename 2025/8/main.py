import sys

from disjoint_set import DisjointSet
from disjoint_set import Pair
from disjoint_set import Point


def main():
    points: list[Point] = []
    for line in sys.stdin:
        x, y, z = (int(num) for num in line.split(","))
        points.append(Point(x, y, z))
    print(f"{len(points) = }")

    part_1_connection_count: int = 10
    if len(points) > 50:
        part_1_connection_count = 1000

    pairs: list[Pair] = create_pairs(points)
    print(f"{len(pairs) = }")

    # sort the pairs by increasing distance
    pairs.sort(key=lambda pair: pair.distance)

    circuits: list[list[Point]] = create_circuits(pairs, len(points), part_1_connection_count)
    print(f"{len(circuits) = }")

    # sort the circuits by decreasing number of points
    circuits.sort(key=len, reverse=True)

    s, t, u = len(circuits[0]), len(circuits[1]), len(circuits[2])
    print(f"Part 1: {s} * {t} * {u} = {s * t * u}")


def create_pairs(points: list[Point]) -> list[Pair]:
    """Creates all possible unique pairs"""
    pairs: list[Pair] = []
    j: int = 1
    while j < len(points):
        i: int = 0
        while i < j:
            pairs.append(Pair(points[i], points[j]))
            i += 1
        j += 1

    return pairs


def create_circuits(
    pairs: list[Pair], point_count: int, part_1_connection_count: int
) -> list[list[Point]]:
    """Combines connected points into circuits

    In this program, "circuit" refers to connected components, not necessarily a path that starts
    and ends at the same point.
    """
    disjoint_set = DisjointSet()
    i: int = 0

    while i < part_1_connection_count:
        disjoint_set.union(pairs[i].a, pairs[i].b)
        i += 1

    circuits: list[list[Point]] = disjoint_set.split()

    while len(disjoint_set) < point_count:
        disjoint_set.union(pairs[i].a, pairs[i].b)
        i += 1
    assert disjoint_set.get_set_count() == 1

    last: Pair = pairs[i - 1]
    print(f"Part 2: {last.a.x} * {last.b.x} = {last.a.x * last.b.x}")

    # make sure no point is used more than once
    points_in_circuits: list[Point] = []
    for circuit in circuits:
        points_in_circuits.extend(circuit)
    assert len(points_in_circuits) == len(set(points_in_circuits))

    return circuits


if __name__ == "__main__":
    main()
