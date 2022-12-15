from collections import namedtuple
from itertools import chain
from typing import Dict, Tuple, List
from os import path

Point = namedtuple("Point", "x y")

def pairwise(i):
    for idx in range(1, len(i)):
        yield (i[idx-1], i[idx])

def mag(v):
    return 1 if v >= 1 else -1 if v <= -1 else 0

def add(a, b: Point):
    return Point(a.x + b.x, a.y + b.y)

def line(a, b: Point):
    vec = Point(mag(b.x-a.x), mag(b.y-a.y))
    while a != b:
        yield a
        a = add(a, vec)
    yield a

def drop_sand(src: Point, cave: Dict, max_depth: int) -> Point:
    vecs = [Point(0, 1), Point(-1, 1), Point(1, 1)]
    sand = src
    while True:
        if sand[1] == max_depth:
            return sand
        nxt = next(
            (p for p in (add(sand, v) for v in vecs) if p not in cave),
            None,
        )
        if nxt is None:
            break
        sand = nxt
    return sand

def make_cave(i: List[List[Point]]) -> Dict[Point, str]:
    lines = chain(*(pairwise(l) for l in i))
    points = chain(*(line(src, tgt) for src, tgt in lines))
    return {p: "#" for p in points}

def part1(i: List[List[Point]]) -> int:
    cave = make_cave(i)
    max_depth = max(p.y for p in cave) + 1
    count = 0

    while True:
        sand = drop_sand(Point(500, 0), cave, max_depth)
        if sand.y == max_depth:
            break
        cave[sand] = "o"
        count += 1

    return count

def part2(i: List[List[Point]]) -> int:
    cave = make_cave(i)
    max_depth = max(p.y for p in cave) + 1
    count = 0

    while True:
        sand = drop_sand(Point(500, 0), cave, max_depth)
        cave[sand] = "o"
        count += 1
        if sand.y == 0:
            break

    return count

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[List[Point]]:
    paths = [
        [Point(*(int(v) for v in p.split(",")))
        for p in l.split(" -> ")]
        for l in i.splitlines()
    ]
    return paths

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()