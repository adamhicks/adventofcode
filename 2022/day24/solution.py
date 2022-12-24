from collections import defaultdict, namedtuple
from itertools import chain
from typing import Set, Tuple
from os import path

Point = namedtuple("Point", "x y")
Blizzard = namedtuple("Blizzard", "point dir")

N = Point(0, -1)
E = Point(1, 0)
S = Point(0, 1)
W = Point(-1, 0)
Wait = Point(0, 0)

char_to_dir = {"^": N, ">": E, "v": S, "<": W}
dir_to_char = {v: k for k, v in char_to_dir.items()}

Options = (Wait, N, E, S, W)

def add(a: Point, b: Point) -> Point:
    return Point(a.x + b.x, a.y + b.y)

def position_at(b: Blizzard, t: int, modp: Point) -> Point:
    return Point(
        (b.point.x + b.dir.x * t) % modp.x, 
        (b.point.y + b.dir.y * t) % modp.y,
    )

def can_come_close(b: Blizzard, p: Point) -> bool:
    return (
        (b.dir.x != 0 and abs(b.point.y - p.y) <= 1)
        or
        (b.dir.y != 0 and abs(b.point.x - p.x) <= 1)
    )

def char(s: Set[Point]) -> str:
    if len(s) == 0:
        return "."
    if len(s) == 1:
        return dir_to_char[next(iter(s))]
    return str(len(s))

def print_map_at(blizz: Tuple[Blizzard], t: int, modp: Point):
    m = defaultdict(set)
    for b in blizz:
        pos = position_at(b, t, modp)
        m[pos].add(b.dir)
    
    for y in range(modp.y):
        print("".join(char(m[Point(x, y)]) for x in range(modp.x)))

def distance(a: Point, b: Point) -> int:
    return abs(a.x - b.x) + abs(a.y - b.y)

def search_route(src: Point, target: Point, t: int, blizz: Tuple[Blizzard]) -> int:
    maxp = Point(
        max(b.point.x for b in blizz) + 1,
        max(b.point.y for b in blizz) + 1,
    )
    seen = set()
    blizz_at = dict()
    walkers = [(t, src)]
    best = 100_000_000
    while len(walkers) > 0:
        t, pos = walkers.pop()
        t += 1

        if t not in blizz_at:
            blizz_at[t] = set(position_at(b, t, maxp) for b in blizz)

        if pos == target:
            best = min(best, t)
            continue

        if t + distance(target, pos) > best:
            continue

        opts = list(filter(lambda p: (
                0 <= p.x < maxp.x and
                0 <= p.y < maxp.y and
                p not in blizz_at[t]
            ),
            (add(pos, o) for o in Options),
        ))
        if pos == src:
            opts.append(pos)

        nxt = [(t, p) for p in opts]
        walkers += filter(lambda n: n not in seen, nxt)
        walkers.sort(key=lambda w: distance(w[1], target), reverse=True)
        seen.update(nxt)
    
    return best

def part1(i: Tuple[Point, Point, Tuple[Blizzard]]) -> int:
    start, end, blizz = i
    return search_route(start, end, 0, blizz)

def part2(i: Tuple[Point, Point, Tuple[Blizzard]]) -> int:
    start, end, blizz = i
    out = search_route(start, end, 0, blizz)
    back = search_route(end, Point(0, 0), out, blizz)
    end = search_route(start, end, back, blizz)
    return end

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> Tuple[Point, Point, Tuple[Blizzard]]:
    line = i.splitlines()
    start = Point(next(i-1 for i, c in enumerate(line[0]) if c == "."), -1)
    end = Point(next(i-1 for i, c in enumerate(line[-1]) if c == "."), len(line) - 3)
    blizz = tuple(chain(*(
        tuple(
            Blizzard(Point(x-1, y), char_to_dir[c])
            for x, c in enumerate(row) 
            if c in char_to_dir
        ) for y, row in enumerate(line[1:-1])
    )))
    return start, end, blizz

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()