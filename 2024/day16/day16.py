import heapq
from collections import namedtuple

test_input = """#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
"""

Point = namedtuple("Point", ["x", "y"])

North = Point(0, -1)
East = Point(1, 0)
South = Point(0, 1)
West = Point(-1, 0)

dirs = (North, East, South, West)

def clockwise(p):
    return Point(-p.y, p.x)

def counter_clockwise(p):
    return Point(p.y, -p.x)

def neighbours(p):
    return [add(p, d) for d in dirs]

def add(a, b):
    return Point(a.x + b.x, a.y + b.y)

def parse_input(s):
    return s.splitlines()

Score = namedtuple("Score", ["score", "pos", "dir", "path"])

def run_part1(i):
    score, _ = find_cheapest_path(i)
    print(score)

def find_cheapest_path(i):
    start = next(Point(x, y) for y, row in enumerate(i) for x, c in enumerate(row) if c == "S")
    end = next(Point(x, y) for y, row in enumerate(i) for x, c in enumerate(row) if c == "E")

    paths = [Score(0, start, East, tuple(start,))]
    lows = set()
    low_score = None
    done = set()
    while paths:
        low = heapq.heappop(paths)
        done.add((low.pos, low.dir))

        if low.pos == end:
            lows.update(low.path)
            low_score = low.score

        if low_score is not None and low.score > low_score:
            continue

        np = add(low.pos, low.dir)

        if i[np.y][np.x] in ".E" and (np, low.dir) not in done:
            heapq.heappush(paths, Score(low.score+1, np, low.dir, low.path + (np,)))

        for d in (clockwise(low.dir), counter_clockwise(low.dir)):
            np = add(low.pos, d)
            if i[np.y][np.x] in ".E" and (np, d) not in done:
                heapq.heappush(paths, Score(low.score+1001, np, d, low.path + (np,)))

    return low_score, len(lows) - 1

def run_part2(i):
    _, nodes = find_cheapest_path(i)
    print(nodes)

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