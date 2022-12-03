from typing import List
from os import path
from functools import reduce

def score(c):
    if c.isupper():
        return (ord(c) - ord('A')) + 27
    else:
        return (ord(c) - ord('a')) + 1

def part1(i: List[str]) -> int:
    sum = 0
    for l in i:
        half = int(len(l)/2)
        first, second = l[:half], l[half:]
        c = set(first) & set(second)
        sum += score(c.pop())
    return sum

def split(l, size):
  for i in range(0, len(l), size):
    yield l[i:i + size]

def part2(i: List[str]) -> int:
    groups = [[frozenset(i) for i in c] for c in split(i, 3)]
    common = [reduce(lambda a, b: a & b, g) for g in groups]
    return sum(score(next(iter(a))) for a in common)

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[str]:
    return i.splitlines()

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()