from collections import namedtuple
from itertools import chain

test_input = """RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
"""

Point = namedtuple("Point", ["x", "y"])

Dirs = (Point(0, -1), Point(1, 0), Point(0, 1), Point(-1, 0))

def add(a, b):
    return Point(a.x+b.x, a.y+b.y)

def clockwise(p):
    return Point(-p.y, p.x)

def counter_clockwise(p):
    return Point(p.y, -p.x)

def neighbours(p):
    return (add(p, d) for d in Dirs)

def parse_input(s):
    return {
        Point(x, y):c
        for y, row in enumerate(s.splitlines())
        for x, c in enumerate(row)
    }

def run_part1(i):
    s = sum(len(r) * sum(len(p) for p in r.values()) for r in find_regions(i))
    print(s)

def find_regions(i):
    todo = list(i.keys())
    done = set()
    regions = []
    c = 0
    while todo:
        n = todo.pop()
        if n in done:
            continue
        r = fill_region(n, i)
        regions.append(r)
        done.update(r.keys())
    return regions

def fill_region(p, i):
    region = {}
    stk = [p]
    while stk:
        n = stk.pop()
        c = i[n]
        region[n] = set()
        for d in Dirs:
            n2 = add(n, d)
            if i.get(n2) == c:
                if n2 not in region:
                    stk.append(n2)
            else:
                region[n].add(d)
    return region

def run_part2(i):
    s = sum(len(r) * sides(r) for r in find_regions(i))
    print(s)

def sides(r):
    s = chain((n, d) for n, ds in r.items() for d in ds)
    s = sorted(s)
    edges = dict()
    for n, d in s:
        n1 = add(n, clockwise(d))
        n2 = add(n, counter_clockwise(d))
        s = 1
        if (n1, d) in edges or (n2, d) in edges:
            s = 0
        edges[(n, d)] = s
    return sum(edges.values())


### Generated by start script

def test_part1():
    run_part1(parse_input(test_input))

def part1():
    run_part1(parse_input(open("input.txt").read()))

def test_part2():
    run_part2(parse_input(test_input))

def part2():
    run_part2(parse_input(open("input.txt").read()))

def main():
    print("=== running part 1 test ===")
    test_part1()
    print("=== running part 1 ===")
    part1()
    print("=== running part 2 test ===")
    test_part2()
    print("=== running part 2 ===")
    part2()
    print("=== ===")

if __name__ == "__main__":
    main()
