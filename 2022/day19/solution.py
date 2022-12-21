from collections import namedtuple
from functools import reduce
from typing import Iterable, List, Tuple
from os import path

Robot = namedtuple("Robot", "build generate")
State = namedtuple("State", "left generate resources")

def buildable(robots: Tuple[Robot], s: State) -> Iterable[Robot]:
    return (
        r for r in robots
        if all(
            have + need >= 0 
            for have, need in zip(s.resources, r.build)
        )
    )

def next_state(s: State, r: Robot) -> State:
    res = tuple(map(sum, zip(s.generate, r.build, s.resources)))

    gen = s.generate
    if r.generate is not None:
        gen = tuple(
            v+1 if idx == r.generate else v 
            for idx, v in enumerate(s.generate)
        )
    return State(s.left-1, gen, res)

def wait_to_build(s: State, r: Robot) -> State:
    res = s.resources
    for i in range(s.left):
        if all(
            have + need >= 0 
            for have, need in zip(res, r.build)
        ):
            res = tuple(map(sum, zip(res, s.generate, r.build)))
            gen = tuple(
                v+1 if idx == r.generate else v 
                for idx, v in enumerate(s.generate)
            )
            return State(s.left - (i + 1), gen, res)
        res = tuple(sum(v) for v in zip(res, s.generate))
    return None

def wait_until_end(s: State) -> State:
    res = tuple(v + (s.left * g) for v, g in zip(s.resources, s.generate))
    return State(0, s.generate, res)

def best_possible(s: State) -> int:
    got = s.resources[-1]
    exist = s.generate[-1] * s.left
    new = (s.left * (s.left + 1)) // 2
    return got + new + exist

def run_build(bp: Tuple[Robot], mins: int) -> int:
    states = [State(mins, (1, 0, 0, 0), (0, 0, 0, 0))]

    seen = set(states[0])

    needed = tuple(min(v) for v in zip(*(r.build for r in bp)))
    best = 0

    while len(states) > 0:
        s = states.pop()

        if s.left == 0:
            best = max(best, s.resources[-1])
            continue

        if best_possible(s) < best:
            continue

        poss = []
        for r in bp:
            if r.generate < 3 and s.generate[r.generate] + needed[r.generate] >= 0:
                continue
            nxt = wait_to_build(s, r)
            if nxt is not None:
                poss.append(nxt)
        if len(poss) == 0:
            poss.append(wait_until_end(s))

        poss = [s for s in poss if s not in seen]
        states += poss
        seen.update(poss)
    
    return best

def part1(i: List[Tuple[Robot]]) -> int:
    return sum((idx + 1) * run_build(bp, 24) for idx, bp in enumerate(i))

def part2(i: List[Tuple[Robot]]) -> int:
    return reduce(lambda a, b: a * b, (run_build(bp, 32) for bp in i[:3]))

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def to_int(i: str) -> int:
    try:
        return int(i)
    except ValueError:
        return None

def parse(i: str) -> List[Tuple[Robot]]:
    blueprints = []
    for bp in i.splitlines():
        num = filter(None, (to_int(s) for s in bp.split()))

        ore = Robot((-next(num), 0, 0, 0), 0)
        clay = Robot((-next(num), 0, 0, 0), 1)
        obsidian = Robot((-next(num), -next(num), 0, 0), 2)
        geode = Robot((-next(num), 0, -next(num), 0), 3)

        blueprints.append((ore, clay, obsidian, geode))
    return blueprints

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()