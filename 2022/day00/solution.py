
from typing import List
from os import path

def part1(i: List[str]) -> int:
    return len(i)

def part2(i: List[str]) -> int:
    return len(i) ** 2

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return i.read()

def parse(i: str) -> List[str]:
    return i.splitlines()

def run():
    i = parse(default_input())
    print(part1(i))
    print(part2(i))

def test():
    s = """This
    is
    a
    test"""
    i = parse(s)
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()