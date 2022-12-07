from typing import List, NamedTuple
from os import path

class Node(NamedTuple):
    id: int
    name: str
    is_dir: bool
    children: dict[str, NamedTuple]
    size: int
    parent: NamedTuple

def id_gen():
    id = 0
    while True:
        yield id
        id += 1

def build_tree(i: List[str]) -> dict:
    id = id_gen()
    root = Node(next(id), "", True, {}, 0, None)
    current = root
    for l in i:
        if l.startswith("$ cd"):
            d = l.split()[2]
            if d == "..":
                current = current.parent
            elif d == "/":
                current = root
            else:
                current = current.children[d]
        elif l == "$ ls":
            pass
        else:
            s, item = l.split()
            this_id = next(id)
            if s == "dir":
                current.children[item] = Node(this_id, item, True, {}, 0, current)
            else:
                current.children[item] = Node(this_id, item, False, {}, int(s), current)
    return root

def print_tree(r: Node, ind: int, sizes: dict):
    s = (ind * "  ") + r.name
    if r.is_dir:
        s += f"/ ({sizes[r.id]})"
        if sizes[r.id] <= 100000:
            s += " ***"
    else:
        s += f" ({r.size})"
    print(s)
    for c in r.children.values():
        print_tree(c, ind+1, sizes)

def calc_sizes(r: Node, nodes: dict) -> int:
    if not r.is_dir:
        return r.size

    s = sum(calc_sizes(n, nodes) for n in r.children.values())
    nodes[r.id] = s
    return s

def part1(i: List[str]) -> int:
    r = build_tree(i)
    sums = {}
    calc_sizes(r, sums)
    return sum(v for _, v in sums.items() if v <= 100000)

def part2(i: List[str]) -> int:
    r = build_tree(i)
    sums = {}
    used = calc_sizes(r, sums)
    need = 30000000 - (70000000 - used)
    return min(v for _, v in sums.items() if v > need)

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
    i = parse("""$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()