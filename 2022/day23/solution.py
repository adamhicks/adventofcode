from collections import defaultdict, namedtuple
from itertools import count
from functools import reduce
from typing import Set
from os import path

Point = namedtuple("Point", "x y")

NW = Point(-1, -1)
N = Point(0, -1)
NE = Point(1, -1)
W = Point(-1, 0)
E = Point(1, 0)
SW = Point(-1, 1)
S = Point(0, 1)
SE = Point(1, 1)

Directions = (
    NW, N, NE, E, SE, S, SW, W
)

Scan = (
    (N, NW, NE),
    (S, SW, SE),
    (W, NW, SW),
    (E, NE, SE),
)

def add(a: Point, b: Point) -> Point:
    return Point(a.x + b.x, a.y + b.y)

def propose_move(p: Point, scan_idx: int, elves: Set[Point]) -> Point:
    neighbours = {d: add(p, d) in elves for d in Directions}
    if not any(neighbours.values()):
        return None
    for i in range(len(Scan)):
        idx = (scan_idx + i) % len(Scan)
        if any(neighbours[d] for d in Scan[idx]):
            continue
        return add(p, Scan[idx][0])
    return None

def inc(d: defaultdict, m: Point) -> defaultdict:
    d[m] += 1
    return d

def round(elves: Set[Point], scan_idx: int) -> set[Point]:
    moves = {p: propose_move(p, scan_idx, elves) for p in elves}
    if not any(moves.values()):
        return None
    count = reduce(inc, moves.values(), defaultdict(int))
    return {
        m if m is not None and count[m] == 1 else p
        for p, m in moves.items()
    }

def min_dim(elves: Set[Point]) -> Point:
    return Point(
        min(p.x for p in elves),
        min(p.y for p in elves),
    )

def max_dim(elves: Set[Point]) -> Point:
    return Point(
        max(p.x for p in elves),
        max(p.y for p in elves),
    )

def print_elves(elves: Set[Point]):
    mi, ma = min_dim(elves), max_dim(elves)
    print(mi, ma)
    for y in range(mi.y, ma.y + 1):
        print("".join("#" if Point(x, y) in elves else "." for x in range(mi.x, ma.x + 1)))
    print("")

def part1(elves: Set[Point]) -> int:
    for i in range(10):
        nxt = round(elves, i)
        if nxt is None:
            break
        elves = nxt
    mi, ma = min_dim(elves), max_dim(elves)
    return ((ma.x - mi.x) + 1) * ((ma.y - mi.y) + 1) - len(elves)

def part2(elves: Set[Point]) -> int:
    for i in count():
        nxt = round(elves, i)
        if nxt is None:
            return i+1
        elves = nxt

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> Set[Point]:
    return reduce(lambda s, n: s | n, (
        set(Point(x, y) for x, c in enumerate(row) if c == "#")
        for y, row in enumerate(i.splitlines())
    ), set())

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()