from typing import List, Tuple
from os import path

def contains(a: Tuple[int], b: Tuple[int]) -> bool:
    return a[0] <= b[0] and a[1] >= b[1]

def part1(i: List[Tuple[Tuple[int]]]) -> int:
    contained = 0
    for pair in i:
        left, right = pair
        if contains(left, right) or contains(right, left):
            contained += 1
    return contained

def overlaps(a: Tuple[int], b: Tuple[int]) -> bool:
    if a[0] > b[0]:
        a, b = b, a
    return b[0] <= a[1]

def part2(i: List[Tuple[Tuple[int]]]) -> int:
    overs = 0
    for pair in i:
        left, right = pair
        if overlaps(left, right):
            overs += 1
    return overs

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[Tuple[Tuple[int]]]:
    return [
        tuple(
            tuple(int(v) for v in iv.split("-"))
            for iv in l.split(",")
        )
        for l in i.splitlines()
    ]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()