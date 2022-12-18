from collections import namedtuple
from itertools import chain, combinations, product
from typing import List, Set
from os import path

Point = namedtuple("Point", "x y z")

def is_neighbour(a, b: Point) -> bool:
    d = (abs(b.x - a.x), abs(b.y - a.y), abs(b.z - a.z))
    return max(d) == 1 and sum(d) == 1

def neighbours(a: Point) -> List[Point]:
    return [
        Point(a.x - 1, a.y, a.z),
        Point(a.x + 1, a.y, a.z),
        Point(a.x, a.y - 1, a.z),
        Point(a.x, a.y + 1, a.z),
        Point(a.x, a.y, a.z - 1),
        Point(a.x, a.y, a.z + 1),
    ]

def faces(l: List[Point]) -> int:
    return len(l) * 6 - sum(
        2 for (a, b) in combinations(l, 2) 
        if is_neighbour(a, b)
    )

def part1(i: List[Point]) -> int:
    return faces(i)

def fill(max: int, present: Set[Point]) -> Set[Point]:
    v = set()
    stk = [Point(0, 0, 0)]
    while len(stk) > 0:
        n = stk.pop()
        for nxt in neighbours(n):
            if nxt in present or nxt in v:
                continue
            if any(not(0 <= c <= max) for c in nxt):
                continue
            v.add(nxt)
            stk.append(nxt)
    return v

def part2(i: List[Point]) -> int:
    m = max(chain(*i)) + 1
    
    internal = set(Point(*p) for p in product(*((range(m),) * 3)))
    lava = set(i)

    internal -= lava
    internal -= fill(m, lava)

    return faces(lava) - faces(list(internal))

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[Point]:
    return [Point(*(int(v) for v in l.split(","))) for l in i.splitlines()]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()