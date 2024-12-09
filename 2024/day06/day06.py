from collections import defaultdict, namedtuple
from copy import deepcopy

test_input = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
"""

Point = namedtuple("Point", ["x", "y"])

def in_bounds(p, grid):
    return 0 <= p.x < len(grid[0]) and 0 <= p.y < len(grid)

def add(a, b):
    return Point(a.x + b.x, a.y + b.y)

def clockwise(p):
    return Point(-p.y, p.x)

def parse_input(s):
    return s.splitlines()

def run_part1(i):
    all_visit = set()
    walk(i, find_start(i), Point(0, -1), defaultdict(set), lambda pos, np, d, visit: all_visit.add(pos))
    print(len(all_visit)+1)

def find_start(grid):
    return next(
        Point(x, y)
        for y in range(len(grid))
        for x in range(len(grid[0]))
        if grid[y][x] == "^"
    )

def walk(grid, pos, d, visited, f):
    visit = deepcopy(visited)
    while True:
        if d in visit[pos]:
            return True
        visit[pos].add(d)
        np = add(pos, d)
        if not in_bounds(np, grid):
            return False
        if grid[np.y][np.x] == "#":
            d = clockwise(d)
        elif f(pos, np, d, visit):
            d = clockwise(d)
        else:
            pos = np

def run_part2(i):
    looper = set()
    def on_pos(pos, np, d, visit):
        if i[np.y][np.x] == "." and len(visit[np]) == 0:
            if walk(i, pos, clockwise(d), visit, lambda pos2, np2, d2, visit2: np2 == np):
                looper.add(np)
        return False

    walk(i, find_start(i), Point(0, -1), defaultdict(set), on_pos)
    print(len(looper))

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