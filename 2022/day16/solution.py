from collections import namedtuple
from itertools import accumulate, chain, combinations
from typing import Dict, List, Tuple
from os import path
from re import findall, search

Node = namedtuple("Node", "name flow to")

def pairwise(l: List) -> List[Tuple]:
    for i in range(1, len(l)):
        yield (l[i-1], l[i])

def distances(g: Dict[str, List[str]], start: str) -> Dict[str, int]:
    nodes = [start]
    dist = dict()
    d = 0
    while len(nodes) > 0:
        dist.update((n, d) for n in nodes)
        nodes = [
            m for m in chain(*(g[n] for n in nodes))
            if m not in dist
        ]
        d += 1
    return dist

def pressure(path: Tuple[str], distance: Dict[str, Dict[str, int]], flow: Dict[str, int], max_t: int) -> int:
    open_times = accumulate(
        pairwise(path),
        func=lambda sum, arc: sum + distance[arc[0]][arc[1]] + 1,
        initial=0,
    )
    return sum((max_t-t) * flow[v] for (v, t) in zip(path, open_times))

def paths(valves: List[str], d: Dict[str, Dict[str, int]], t: int) -> List[Tuple[str]]:
    stk = [(("AA",), t, set(valves))]
    while len(stk) > 0:
        path, time_left, unvisited = stk.pop()
        yield path

        for n in unvisited:
            t = time_left - d[path[-1]][n]
            if t <= 0:
                continue
            p = path + (n,)
            s = unvisited.copy()
            s.remove(n)

            stk.append((p, t, s))

def part1(i: List[Node]) -> int:
    g = {n.name: n.to for n in i}
    f = {n.name: n.flow for n in i}
    valves = [n.name for n in i if n.flow > 0]
    vd = {v: distances(g, v) for v in valves}
    vd["AA"] = distances(g, "AA")

    max_t = 30

    return max(
        pressure(p, vd, f, max_t)
        for p in paths(valves, vd, max_t)
    )

def part2(i: List[Node]) -> int:
    g = {n.name: n.to for n in i}
    f = {n.name: n.flow for n in i}
    valves = [n.name for n in i if n.flow > 0]
    vd = {v: distances(g, v) for v in valves}
    vd["AA"] = distances(g, "AA")

    t = 26

    all_paths = sorted(
        paths(valves, vd, t), 
        key=lambda p: pressure(p, vd, f, t), 
        reverse=True,
    )

    low = next(
        idx for idx, p in enumerate(all_paths) 
        if set(p) & set(all_paths[0]) == {"AA"}
    )

    return max(
        pressure(p1, vd, f, t) + pressure(p2, vd, f, t) for
        (p1, p2) in combinations(all_paths[:low+1], 2)
        if set(p1) & set(p2) == {"AA"}
    )

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[Node]:
    nodes = []
    for l in i.splitlines():
        f = search(r"rate=([0-9]+)", l)
        names = findall(r"([A-Z]{2})", l)
        nodes.append(Node(names[0], int(f.group(1)), names[1:]))
    return nodes

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()