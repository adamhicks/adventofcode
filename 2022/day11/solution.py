from collections import namedtuple
from copy import deepcopy
from functools import reduce
from typing import List
from os import path

Monkey = namedtuple("Monkey", "stuff operation test")
Test = namedtuple("Test", "div t1 t2")

def part1(i: List[namedtuple]) -> int:
    monkeys = deepcopy(i)
    counts = [0 for _ in range(len(monkeys))]
    for round in range(20):
        for idx, m in enumerate(monkeys):
            counts[idx] += len(m.stuff)
            for item in m.stuff:
                item = do_op(item, m.operation)
                item //= 3

                if item % m.test.div == 0:
                    tgt = m.test.t1
                else:
                    tgt = m.test.t2

                monkeys[tgt].stuff.append(item)
            m.stuff.clear()

    first, second = sorted(counts)[-2:]
    return first*second

def part2(i: List[str]) -> int:
    monkeys = deepcopy(i)
    counts = [0 for _ in range(len(monkeys))]
    modo = reduce(lambda a,b: a*b, (m.test.div for m in monkeys))
    for round in range(10_000):
        for idx, m in enumerate(monkeys):
            counts[idx] += len(m.stuff)
            for item in m.stuff:
                item = do_op(item, m.operation)
                item %= modo

                if item % m.test.div == 0:
                    tgt = m.test.t1
                else:
                    tgt = m.test.t2

                monkeys[tgt].stuff.append(item)
            m.stuff.clear()

    first, second = sorted(counts)[-2:]
    return first*second

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def do_op(old: int, op: str) -> int:
    return eval(op)

def parse(i: str) -> List[str]:
    monkeys = []
    for m in (j.splitlines() for j in i.split("\n\n")):
        items = [int(item) for item in m[1].split(": ")[1].split(", ")]
        op = m[2].split(" = ")[1]
        test = Test(*(int(l.split()[-1]) for l in m[-3:]))
        monkeys.append(Monkey(items, op, test))
    return monkeys

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()