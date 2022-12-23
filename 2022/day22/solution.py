from collections import namedtuple
from itertools import product
from typing import Dict, Iterable, List, Set, Tuple
from os import path
from re import split

Point = namedtuple("Point", "x y")

Left = Point(-1, 0)
Right = Point(1, 0)
Up = Point(0, -1)
Down = Point(0, 1)

Directions = [Right, Down, Left, Up]

def add(p1: Point, p2: Point) -> Point:
    return Point(p1.x + p2.x, p1.y + p2.y)

def sub(p1: Point, p2: Point) -> Point:
    return Point(p1.x - p2.x, p1.y - p2.y)

def neg(p: Point) -> Point:
    return Point(-p.x, -p.y)

def turn(p: Point, clockwise: bool):
    idx = Directions.index(p)
    if clockwise:
        idx += 1
    else:
        idx -= 1
    return Directions[idx % len(Directions)]

def wrap(frm: Point, d: Point, mapper: Dict[Point, str]) -> Point:
    if d == Left:
        return Point(max(p.x for p in mapper if p.y == frm.y), frm.y)
    elif d == Right:
        return Point(min(p.x for p in mapper if p.y == frm.y), frm.y)
    elif d == Up:
        return Point(frm.x, max(p.y for p in mapper if p.x == frm.x))
    elif d == Down:
        return Point(frm.x, min(p.y for p in mapper if p.x == frm.x))
    raise ValueError("invalid direction", d)

def travel(p: Point, d: Point, i: int, mapper: Dict[Point, str]) -> Point:
    for _ in range(i):
        nxt = add(p, d)
        if nxt not in mapper:
            nxt = wrap(nxt, d, mapper)
        if mapper[nxt] == "#":
            return p
        p = nxt
    return p

def part1(i: Tuple[Dict[Point,str], List]) -> int:
    mapper, ins = i

    loc = Point(min(p.x for p in mapper if p.y == 0), 0)
    dir = Point(1, 0)

    for instruct in ins:
        if instruct in ("L", "R"):
            dir = turn(dir, instruct == "R")
        else:
            loc = travel(loc, dir, instruct, mapper)
    
    px = loc.x + 1
    py = loc.y + 1
    pd = Directions.index(dir)

    return 1000*py + 4*px + pd

def neighbours(p: Point) -> Iterable[Point]:
    return (
        add(p, Point(*m))
        for m in product(*((-1, 0, 1),) * 2)
    )

def get_inward(p: Point, dir: Point, mapper: Set[Point]) -> Point:
    right = turn(dir, True)
    left = turn(dir, False)
    r_in = add(p, right) in mapper
    l_in = add(p, left) in mapper

    if r_in and not l_in:
        return right
    elif l_in and not r_in:
        return left
    else:
        return None

def find_next(p: Point, dir: Point, mapper: Set[Point]) -> Tuple[Point, Point, Point]:
    nxt = add(p, dir)
    inward = get_inward(nxt, dir, mapper)
    if inward is not None:
        return nxt, dir, inward

    right = turn(dir, True)
    r_in = get_inward(p, right, mapper)
    if r_in is not None and add(p, right) in mapper:
        return p, right, r_in

    left = turn(dir, False)
    l_in = get_inward(p, left, mapper)
    if l_in is not None and add(p, left) in mapper:
        return p, left, l_in
    
    inw = get_inward(p, dir, mapper)
    d = neg(inw)
    i = dir
    nxt = add(add(p, i), d)
    return nxt, d, i

def walk(start: Point, dir: Point, mapper: Set[Point]) -> Iterable[Tuple[Point, Point, Point]]:
    p = start
    inward = get_inward(p, dir, mapper)
    while True:
        yield p, dir, inward
        p, dir, inward = find_next(p, dir, mapper)

Walker = namedtuple("Walker", "position direction inward")

def until_separate(w: Iterable[Tuple[Walker, Walker]]) -> Iterable[Tuple[Walker, Walker]]:
    d1, d2 = None, None
    for w1, w2 in w:
        nd1, nd2 = w1[1], w2[1]
        if d1 is not None and d2 is not None:
            if d1 != nd1 and d2 != nd2:
                break
        d1, d2 = nd1, nd2
        yield w1, w2

def perimeter(mapper: Dict[Point, str]) -> Iterable[Iterable[Tuple[Point, Point]]]:
    points = set(mapper)
    inner_corners = [
        p for p in mapper
        if len(points & set(neighbours(p))) == 8
    ]

    for c in inner_corners:
        missing = next(p for p in neighbours(c) if p not in mapper)

        a, b = Point(c.x, missing.y), Point(missing.x, c.y)
        ad, bd = sub(a, c), sub(b, c)
        yield until_separate(zip(walk(a, ad, mapper), walk(b, bd, mapper)))

def travel2(
    p: Point, d: Point, i: int,
    mapper: Dict[Point, str],
    move: Dict[Tuple[Point, Point], Tuple[Point, Point]],
) -> Point:
    for _ in range(i):
        if (p, d) in move:
            nxt, nd = move[(p, d)]
        else:
            nxt = add(p, d)
            nd = d
        if mapper[nxt] == "#":
            return p, d
        p = nxt
        d = nd
    return p, d

def part2(i: Tuple[Dict[Point,str], List]) -> int:
    mapper, ins = i

    move = dict()
    for z in perimeter(mapper):
        for a, b in z:
            a_p, _, a_in = a
            a_out = neg(a_in)
            b_p, _, b_in = b
            b_out = neg(b_in)
            move[(a_p, a_out)] = (b_p, b_in)
            move[(b_p, b_out)] = (a_p, a_in)

    loc = Point(min(p.x for p in mapper if p.y == 0), 0)
    dir = Point(1, 0)

    for instruct in ins:
        if instruct in ("L", "R"):
            dir = turn(dir, instruct == "R")
        else:
            loc, dir = travel2(loc, dir, instruct, mapper, move)

    px = loc.x + 1
    py = loc.y + 1
    pd = Directions.index(dir)

    return 1000*py + 4*px + pd

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> Tuple[Dict[Point,str], List]:
    mapper, instruct = i.split("\n\n")

    m = dict()
    for y, row in enumerate(mapper.splitlines()):
        for x, c in enumerate(row):
            if c != " ":
                m[Point(x, y)] = c
    
    ins = [v if v in ("L", "R") else int(v) for v in split(r"([LR])", instruct)]
    return m, ins

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()