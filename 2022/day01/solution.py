from typing import List
from os import path

def part1(i: List[List[int]]) -> int:
    return max((sum(j) for j in i))

def part2(i: List[List[int]]) -> int:
    return sum(sorted((sum(j) for j in i), reverse=True)[:3])

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[List[int]]:
    return [[int(k) for k in l.splitlines()] for l in i.split("\n\n")]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))


def test():
    i = parse("""1000
2000
3000

4000

5000
6000

7000
8000
9000

10000""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()
