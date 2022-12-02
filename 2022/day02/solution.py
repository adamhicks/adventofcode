from typing import List
from os import path

def beats(a: int) -> int:
    return (a + 1) % 3

def loses_to(a: int) -> int:
    return (a - 1) % 3

def part1(i: List[List[str]]) -> int:
    score = 0
    for round in i:
        them = ord(round[0]) - ord('A')
        us = ord(round[1]) - ord('X')
        score += us + 1
        if us == beats(them):
            score += 6
        elif us == them:
            score += 3
    return score

def target(result: int, them: int) -> int:
    if result == 0:
        return them
    elif result == 1:
        return (them + 1) % 3
    elif result == 2:
        return 

def part2(i: List[str]) -> int:
    score = 0
    for round in i:
        them = ord(round[0]) - ord('A')
        result = round[1]
        if result == "X":
            us = loses_to(them)
        elif result == "Y":
            us = them
            score += 3
        elif result == "Z":
            us = beats(them)
            score += 6
        score += us + 1
    return score

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[List[str]]:
    return [j.split() for j in i.splitlines()]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""A Y
B X
C Z""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()