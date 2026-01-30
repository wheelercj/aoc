import math
from dataclasses import dataclass


@dataclass(order=True, frozen=True)
class Point:
    x: int
    y: int
    z: int


class Pair:
    def __init__(self, a: Point, b: Point) -> None:
        self.a: Point = a
        self.b: Point = b
        self.distance: float = self.__get_distance(a, b)

    def __get_distance(self, a: Point, b: Point) -> float:
        return math.hypot(a.x - b.x, a.y - b.y, a.z - b.z)


class DisjointSet:
    def __init__(self, pairs: list[Pair] | None = None) -> None:
        """A collection of non-overlapping sets

        This disjoint set is implemented as a forest of parent pointer trees where the paths are
        compressed to make lookups of roots faster. The disjoint-set data structure is also known
        as union-find or merge-find. More details:
        https://en.wikipedia.org/wiki/Disjoint-set_data_structure#Representation
        """
        self.__set_count: int = 0
        self.__parents: dict[Point, Point] = dict()

        if pairs:
            for pair in pairs:
                self.union(pair.a, pair.b)

    def __len__(self) -> int:
        """Returns the number of points in the disjoint set"""
        return len(self.__parents)

    def get_set_count(self) -> int:
        """Returns the number of sets in the disjoint set"""
        return self.__set_count

    def union(self, p1: Point, p2: Point) -> None:
        """Combine the trees containing p1 and p2

        If any of the points are not in the disjoint set yet, they are added to it.
        """
        if p1 not in self.__parents:
            self.__parents[p1] = p1
            self.__set_count += 1
        if p2 not in self.__parents:
            self.__parents[p2] = p2
            self.__set_count += 1

        root1: Point = self.find(p1)
        root2: Point = self.find(p2)
        if root1 != root2:
            self.__parents[root1] = root2
            self.__set_count -= 1

    def find(self, p: Point) -> Point:
        """Finds the root of the point's parent pointer tree

        If the given point is a root, it is returned.
        """
        # compress the path for faster lookups
        if self.__parents[p] != p:  # the roots are self-referencing
            self.__parents[p] = self.find(self.__parents[p])

        return self.__parents[p]

    def split(self) -> list[list[Point]]:
        """Splits the disjoint sets of points into their own lists"""
        results: dict[Point, list[Point]] = dict()
        for point in self.__parents:
            root: Point = self.find(point)
            if root not in results:
                results[root] = []
            results[root].append(point)

        return list(results.values())
