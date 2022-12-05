from copy import deepcopy
from typing import List, Tuple
from os import path

def part1(i: Tuple[List[List[str]], List[Tuple[int, int, int]]]) -> str:
    stks, inst = i
    stks = deepcopy(stks)
    for ins in inst:
        amt = ins[0]
        src, tgt = ins[1]-1, ins[2]-1
        for j in range(amt):
            v = stks[src].pop()
            stks[tgt].append(v)
    return ''.join(c[-1] for c in stks)

def part2(i: List[str]) -> int:
    stks, inst = i
    stks = deepcopy(stks)
    for ins in inst:
        amt = ins[0]
        src, tgt = ins[1]-1, ins[2]-1
        from_stk = stks[src]
        rem, move = from_stk[:-amt], from_stk[-amt:]

        stks[src] = rem
        stks[tgt] += move

    return ''.join(c[-1] for c in stks)

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> Tuple[List[List[str]], List[Tuple[int, int, int]]]:
    stacks, instruct = i.split("\n\n")
    stacks = stacks.splitlines()
    indexes = stacks.pop()
    num = (len(indexes)+1) // 4

    stacks.reverse()

    stks = []
    for i in range(num):
        idx = (i * 4) + 1
        stks.append([s[idx] for s in stacks if s[idx] != " "])

    inst = [
        tuple(
            int(v) for (idx, v) 
            in enumerate(insline.split()) if idx % 2 == 1
        )
        for insline in instruct.splitlines()
    ]
    return stks, inst

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()