from typing import List, Tuple
from os import path

def closest(h, t: int) -> int:
    if t > h:
        return h+1
    else:
        return h-1

def catchup(h, t: Tuple[int, int]) -> Tuple[int, int]:
    x_dis = abs(h[0]-t[0])
    y_dis = abs(h[1]-t[1])
    dis = max(x_dis, y_dis)
    if dis <= 1:
        return t
    if x_dis < y_dis:
        return (h[0], closest(h[1], t[1]))
    elif y_dis < x_dis:
        return (closest(h[0], t[0]), h[1])
    else:
        return (closest(h[0], t[0]), closest(h[1], t[1]))

def move(p: Tuple[int, int], d: str, v: int) -> Tuple[int, int]:
    if d == "U":
        return (p[0], p[1]+v)
    elif d == "D":
        return (p[0], p[1]-v)
    elif d == "R":
        return (p[0]+v, p[1])
    elif d == "L":
        return (p[0]-v, p[1])
    raise ValueError(f"unknown direction '{d}'")

def print_rope(x, y: int, rope: List[Tuple[int, int]]):
    rows = [["." for _ in range(x)] for _ in range(y)]
    for idx, i in enumerate(rope):
        rx, ry = i
        if rx < 0 or rx >= x or ry < 0 or ry >= y:
            continue
        if rows[ry][rx] == ".":
            rows[ry][rx] = str(idx)
    for r in reversed(rows):
        print("".join(r))
    print()

def move_rope(rope: List[Tuple[int, int]], d: str):
    rope[0] = move(rope[0], d, 1)
    for j in range(1, len(rope)):
        rope[j] = catchup(rope[j-1], rope[j])

def part1(i: List[Tuple[str, int]]) -> int:
    rope = [(0, 0), (0, 0)]
    trace = set()
    trace.add(rope[-1])
    for ins in i:
        for _ in range(ins[1]):
            move_rope(rope, ins[0])
            trace.add(rope[-1])
    return len(trace)

def part2(i: List[Tuple[str, int]]) -> int:
    rope = [(0, 0) for i in range(10)]
    trace = set()
    trace.add(rope[-1])
    for ins in i:
        for _ in range(ins[1]):
            move_rope(rope, ins[0])
            trace.add(rope[-1])
    return len(trace)

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[Tuple[str, int]]:
    a = [l.split() for l in i.splitlines()]
    return [(c, int(v)) for c, v in a]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():

    s = ["""R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2""", """R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20"""]

    i1 = parse(s[0])
    i2 = parse(s[1])

    print(part1(i1))
    print(part2(i1))

    print(part2(i2))

if __name__ == "__main__":
    test()