from functools import reduce
from itertools import product
from typing import List
from os import path

def clear(line: List[int]) -> bool:
    return max(line[1:], default=-1) < line[0]

def visible(x: int, y: int, forest: List[List[int]], forest_cols: List[List[int]]) -> bool:
    return any(clear(d) for d in [
        forest[y][x::-1],
        forest[y][x:],
        forest_cols[x][y::-1],
        forest_cols[x][y:]
    ])

def sight(line: List[int]) -> int:
    return next(
        (idx+1 for idx, t in enumerate(line[1:]) if t >= line[0]),
        len(line)-1,
    )

def score(x: int, y: int, forest: List[List[int]], forest_cols: List[List[int]]) -> int:
    return reduce(lambda a, b: a*b, (sight(d) for d in [
        forest[y][x::-1],
        forest[y][x:],
        forest_cols[x][y::-1],
        forest_cols[x][y:],
    ]))


def part1(i: List[List[int]]) -> int:
    cols = list(zip(*i))
    return len([
        (x, y)
        for x, y in product(range(len(cols)), range(len(i)))
        if visible(x, y, i, cols)
    ])

def part2(i: List[List[int]]) -> int:
    cols = list(zip(*i))
    return max(score(x, y, i, cols) for x, y in product(range(len(cols)), range(len(i))))

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[List[int]]:
    return [[int(v) for v in j] for j in i.splitlines()]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""30373
25512
65332
33549
35390""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()