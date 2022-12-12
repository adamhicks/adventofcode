from itertools import chain
from typing import List, Tuple
from os import path

def find_start(i: List[str]) -> Tuple[int, int]:
    for y, row in enumerate(i):
        for x, c in enumerate(row):
            if c == "S":
                return (x, y)
    raise ValueError("no start node")

def neighbours(p: Tuple[int, int]) -> List[Tuple[int, int]]:
    yield p[0], p[1]-1
    yield p[0], p[1]+1
    yield p[0]-1, p[1]
    yield p[0]+1, p[1]

def in_bounds(a, bounds: Tuple[int, int]) -> bool:
    return a[0] >= 0 and a[0] < bounds[0] and a[1] >= 0 and a[1] < bounds[1]

def fetch(i: List[str], p: Tuple[int, int]) -> str:
    return i[p[1]][p[0]]

def height_of(i: List[str], p: Tuple[int, int]) -> int:
    c = fetch(i, p)
    if c == "S":
        return "a"
    elif c == "E":
        return "z"
    else:
        return c

def height_diff(i: List[str], p1, p2: Tuple[int, int]) -> int:
    pc1 = height_of(i, p1)
    pc2 = height_of(i, p2)
    return ord(pc2) - ord(pc1)

def shortest_route(i: List[str], start: List[Tuple[int, int]]) -> Tuple:
    bounds = len(i[0]), len(i)
    routes = start
    visited = set()
    while len(routes) > 0:
        nxt = []
        for r in routes:
            p = r[-1]
            if fetch(i, p) == "E":
                return r
            for n in neighbours(p):
                if n in visited:
                    continue
                if in_bounds(n, bounds) and height_diff(i, p, n) <= 1:
                    visited.add(n)
                    nxt.append((r + (n,)))
        routes = nxt
    raise Exception("no route found")

def part1(i: List[str]) -> int:
    return len(shortest_route(i, [(find_start(i),)])) - 1

def part2(i: List[str]) -> int:
    all_as = [
        [(x, y) for x, v in enumerate(row) if v == "a"]
        for y, row in enumerate(i)
    ]
    start = [(p,) for p in chain(*all_as)]
    return len(shortest_route(i, start)) - 1

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[str]:
    return i.splitlines()

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()