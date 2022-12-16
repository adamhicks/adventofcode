from collections import namedtuple
from itertools import chain, combinations, product
from typing import List, Set, Tuple
from os import path
from re import finditer

Point = namedtuple("Point", "x y")

def distance(a, b: Point) -> int:
    return abs(b.x-a.x) + abs(b.y-a.y)

def project(a: Point, d: int, y: int) -> Tuple[int, int]:
    width = d - abs(a.y - y)
    if width <= 0:
        return None
    return (a.x-width, a.x+width)

def excluded(sensors: List[Tuple[Point, Point, int]], y: int) -> Set[int]:
    return set(
        chain(*(range(a, b+1)
        for a, b in
        filter(None, (project(s, d, y) for s, _, d in sensors))
        ))
    )

def part1(i: List[List[Point]], y: int) -> int:
    sensors = [((s, b, distance(s, b))) for s, b in i]
    return len(excluded(sensors, y)) - 1

def in_range(p: Point, max: int):
    return p.x >= 0 and p.x <= max and p.y >= 0 and p.y <= max

def borderlines(p: Point, d: int) -> List[Tuple[Point, Point]]:
    yield (Point(p.x - d, p.y), Point(p.x, p.y + d))
    yield (Point(p.x, p.y + d), Point(p.x + d, p.y))
    yield (Point(p.x - d, p.y), Point(p.x, p.y - d))
    yield (Point(p.x, p.y - d), Point(p.x + d, p.y))

def mag(i: int) -> int:
    return 1 if i > 0 else -1 if i < 0 else 0

def gradient(line: Tuple[Point, Point]) -> int:
    return mag(line[1].y - line[0].y)

def between(a, i, b: int) -> bool:
    if b < a:
        a, b = b, a
    return a <= i <= b

def on_line(p: Point, l: Tuple[Point, Point]) -> bool:
    return between(l[0].x, p.x, l[1].x) and between(l[0].y, p.y, l[1].y)

def point_intersect(l1, l2: Tuple[Point, Point]) -> Point:
    g1, g2 = gradient(l1), gradient(l2)
    if g1 == g2:
        return None
    if g1 == -1:
        l1, l2 = l2, l1

    xi1 = l1[0].x - l1[0].y
    xi2 = l2[0].x + l2[0].y
    
    d = (xi2 - xi1) // 2
    midx = xi1 + d
    p = Point(midx, d)
    if on_line(p, l1) and on_line(p, l2):
        return p
    return None

def border_intersect(p1: Point, d1: int, p2: Point, d2: int) -> List[Point]:
    return filter(None, (
        point_intersect(l1, l2) 
        for l1, l2 in product(borderlines(p1, d1), borderlines(p2, d2))
    ))

def part2(i: List[List[Point]], max: int) -> int:
    possibles = set(
        p for p in chain(*(
            border_intersect(a[0], distance(a[0], a[1])+1, b[0], distance(b[0], b[1])+1)
            for a, b in combinations(i, 2)
        )) if in_range(p, max)
    )
    for s, b in i:
        possibles = set(p for p in possibles if distance(s, p) > distance(s, b))

    if len(possibles) != 1:
        raise Exception("didn't find unique pixel")
    
    x, y = possibles.pop()
    return x * 4_000_000 + y

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[List[Point]]:
    return [
        [Point(int(m.group(1)), int(m.group(2))) 
        for m in finditer(r"x=([-0-9]+), y=([-0-9]+)", line)] 
        for line in i.splitlines()
    ]

def run():
    i = default_input()
    print(part1(i, 2_000_000))
    print(part2(i, 4_000_000))

def test():
    i = parse("""Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3""")
    print(part1(i, 10))
    print(part2(i, 20))

if __name__ == "__main__":
    test()