from typing import List
from os import path

def unique_index(s: str, l: int) -> int:
    for i in range(len(s)-1-l):
        if len(set(s[i:i+l])) == l:
            return l+i
    raise ValueError("no header found")

def part1(i: str) -> int:
    return unique_index(i, 4)

def part2(i: str) -> int:
    return unique_index(i, 14)

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> str:
    return i.strip()

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = [
        parse("""mjqjpqmgbljsphdztnvjfqwrcgsmlb"""),
        parse("""bvwbjplbgvbhsrlpgdmjqwftvncz"""),
        parse("""nppdvjthqldpwncqszvftbrmjlhg"""),
        parse("""nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"""),
        parse("""zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"""),
    ]
    for j in i:
        print(j, part1(j))
    for j in i:
        print(j, part2(j))

if __name__ == "__main__":
    test()