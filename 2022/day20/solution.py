from typing import List
from os import path

class Node:
    def __init__(self, v: int) -> None:
        self.v = v
        self.left = None
        self.right = None

def pairwise(l: List):
    for i in range(1, len(l)):
        yield (l[i-1], l[i])

def join(li: List[Node]):
    for l, r in pairwise(li):
        l.right = r
        r.left = l
    li[0].left = li[-1]
    li[-1].right = li[0]

def pop(i: Node):
    l, r = i.left, i.right
    l.right = r
    r.left = l

def insert_after(i: Node, l: Node):
    r = l.right
    l.right = i
    r.left = i
    i.left = l
    i.right = r

def to_list(i: Node) -> List[int]:
    r = [i.v]
    s = i.right
    while s != i:
        r.append(s.v)
        s = s.right
    return r

def from_list(l: List[int]) -> List[Node]:
    r = [Node(v) for v in l]
    join(r)
    return r

def shift(n: Node, c: int, size: int) -> Node:
    if c < 0:
        c = size + c
    r = n.left
    for _ in range(c % size):
        r = r.right
    return r

def mix(l: List[Node]):
    for n in l:
        pop(n)
        at = shift(n, n.v, len(l)-1)
        insert_after(n, at)

def part1(i: List[int]) -> int:
    l = from_list(i)
    mix(l)
    f = to_list(l[0])
    zero = f.index(0)
    return sum(f[(zero + i) % len(f)] for i in [1000, 2000, 3000])

def part2(i: List[str]) -> int:
    l = from_list(v * 811589153 for v in i)
    for _ in range(10):
        mix(l)
    f = to_list(l[0])
    zero = f.index(0)
    return sum(f[(zero + i) % len(f)] for i in [1000, 2000, 3000])

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[int]:
    return [int(v) for v in i.splitlines()]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""1
2
-3
3
-2
0
4""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()
