from functools import cmp_to_key
from itertools import chain, zip_longest
from typing import List
from os import path

def cmp_int(i1, i2) -> int:
    return 1 if i1 > i2 else -1 if i1 < i2 else 0

def compare(l1, l2) -> int:
    l1_list = isinstance(l1, list)
    l2_list = isinstance(l2, list)

    if not l1_list and not l2_list:
        return cmp_int(l1, l2)
    
    if not l1_list:
        l1 = [l1]
    if not l2_list:
        l2 = [l2]
    
    for e1, e2 in zip_longest(l1, l2):
        if e1 is None:
            return -1
        if e2 is None:
            return 1
        c = compare(e1, e2)
        if c != 0:
            return c
    return 0

def part1(i: List) -> int:
    return sum(idx+1 for idx, p in enumerate(i) if compare(*p) <= 0)

def part2(i: List) -> int:
    divs = [[[2]], [[6]]]
    l = list(chain(*i, divs))
    l.sort(key=cmp_to_key(compare))
    return (l.index(divs[0])+1) * (l.index(divs[1])+1)

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List:
    return [[eval(l) for l in c.splitlines()] for c in i.split("\n\n")]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()