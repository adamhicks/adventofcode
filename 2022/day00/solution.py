
from typing import List
from os import path

def part1(i: List[str]) -> int:
    return len(i)

def part2(i: List[str]) -> int:
    return len(i) ** 2

def default_input() -> List[str]:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return i.readlines()

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))
