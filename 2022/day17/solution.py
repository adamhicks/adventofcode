from collections import defaultdict, namedtuple
from typing import List, Set, Tuple
from os import path

Point = namedtuple("Point", "x y")

class Repeater:
    def __init__(self, it):
        self.l = list(it)

    def __iter__(self):
        self.idx = 0
        self.repeat = 0
        return self

    def __next__(self):
        o = self.l[self.idx]
        self.idx += 1
        if self.idx == len(self.l):
            self.idx = 0
            self.repeat += 1
        return o
    
    def emitted(self):
        return self.idx + self.repeat * len(self.l)
    
def get_rocks() -> List[Tuple[Point]]:
    return [
        (Point(0, 0), Point(1, 0), Point(2, 0), Point(3, 0)),
        (Point(0, 1), Point(1, 0), Point(1, 1), Point(1, 2), Point(2, 1)),
        (Point(0, 0), Point(1, 0), Point(2, 0), Point(2, 1), Point(2, 2)),
        (Point(0, 0), Point(0, 1), Point(0, 2), Point(0, 3)),
        (Point(0, 0), Point(0, 1), Point(1, 0), Point(1, 1)),
    ]

def wind_direction(s: chr) -> Point:
    if s == "<":
        return Point(-1, 0)
    elif s == ">":
        return Point(1, 0)
    raise ValueError(f"unknown wind {s}")

def add(a, b: Point) -> Point:
    return Point(a.x + b.x, a.y + b.y)

def move(r: Tuple[Point], w: Point) -> Tuple[Point]:
    return tuple(add(p, w) for p in r)

def hit_walls(r: Tuple[Point], width: int) -> bool:
    return any(p.x < 0 or p.x >= width for p in r)

def hit_other_rock(r: Tuple[Point], others: Set[Point]) -> bool:
    return bool(set(r) & others)

def hit_floor(r: Tuple[Point]) -> bool:
    return lowest(r) < 0

def start_rock(r: Tuple[Point], highest: int) -> Tuple[Point]:
    return move(r, Point(2, highest+4))

def highest(blocked: Set[Point]) -> int:
    return max((p.y for p in blocked), default=-1)

def lowest(blocked: Set[Point]) -> int:
    return min((p.y for p in blocked), default=-1)

def print_column(rock: Tuple[Point], blocked: Set[Point], width, height: int):
    r = set(rock)
    h = max(highest(blocked), highest(r)) + 3
    for y in range(h, max(h-height, -1), -1):
        print("".join(
            '#' if Point(x, y) in r
            else '@' if Point(x, y) in blocked
            else '.' 
            for x in range(0, width)
        ))
    print("")

down = Point(0, -1)

def drop(rock: Tuple[Point], wind: List[str], width: int, blocked: Set[Point]) -> Tuple[Point]:
    r = start_rock(rock, highest(blocked))
    while True:
        w = wind_direction(next(wind))

        nxt = move(r, w)
        if not(hit_walls(nxt, width) or hit_other_rock(nxt, blocked)):
            r = nxt

        nxt = move(r, down)
        if hit_other_rock(nxt, blocked) or hit_floor(nxt):
            return r
        r = nxt

def part1(i: str) -> int:
    rock = iter(Repeater(get_rocks()))
    wind = iter(Repeater(i))
    width = 7
    blocked = set()

    for i in range(2022):
        r = drop(next(rock), wind, width, blocked)
        blocked.update(r)
        blocked = top(blocked, width)

    return highest(blocked)+1

def top(blocked: Set[Point], width: int) -> Set[Point]:
    h = highest(blocked)
    x_in = set(range(width))
    res = set()
    for d in range(h+1):
        row = [p for p in (Point(x, h-d) for x in range(width))]
        res.update(p for p in row if p in blocked)

        x_empty = set(p.x for p in row if p not in blocked)
        x_in &= x_empty
        fill = {x for x in x_empty - x_in if x - 1 in x_in or x + 1 in x_in}
        x_in |= fill

        if len(x_in) == 0:
            break

    return res

def signature(blocked: Set[Point]) -> int:
    low = lowest(blocked)
    return hash(tuple(sorted(
        Point(p.x, p.y - low) for p in blocked
    )))

def part2(i: List[str]) -> int:
    rock = iter(Repeater(get_rocks()))
    wind = iter(Repeater(i))
    width = 7

    seen = defaultdict(dict)

    blocked = set()
    i = 0
    tgt = 1_000_000_000_000

    while True:
        r = drop(next(rock), wind, width, blocked)
        blocked.update(r)
        i += 1

        t = top(blocked, width)
        key = (rock.idx, wind.idx)
        sig = signature(t)
        if sig in seen[key]:
            start, h_start = seen[key][sig]
            end, h_end = rock.emitted(), highest(t)
            cycle_len = end - start

            skip = (tgt - end) // cycle_len
            elev = (h_end - h_start) * skip

            i += skip * cycle_len
            blocked = {Point(p.x, p.y + elev) for p in t}
            break
        else:
            seen[key][sig] = (rock.emitted(), highest(t))
        blocked = t

    while i < tgt:
        r = drop(next(rock), wind, width, blocked)
        blocked.update(r)
        i += 1

    return highest(blocked)+1

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> str:
    return i.splitlines()[0]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse(""">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()