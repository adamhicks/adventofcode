from itertools import chain, count
from typing import Dict, List, Tuple
from os import path

def depth(n: str, dep: Dict[str, str]) -> int:
    for i in count():
        n = dep.get(n, None)
        if n is None:
            return i

def eval_eq(v: str, eq: Tuple, vals: Dict[str, int]) -> Tuple[str, int]:
    if len(eq) == 1:
        return v, int(eq[0])

    l = vals.get(eq[0])
    r = vals.get(eq[2])

    if l is None or r is None:
        return None, None

    if eq[1] == "+":
        return v, l + r
    elif eq[1] == "-":
        return v, l - r
    elif eq[1] == "*":
        return v, l * r
    elif eq[1] == "/":
        return v, l // r

def solve_eq(v: str, eq: Tuple, vals: Dict[str, int]) -> Tuple[str, int]:
    vv = vals[v]
    l, op, r = eq
    lv = vals.get(l)
    rv = vals.get(r)

    if op == "+":
        if lv is None:
            return l, vv - rv
        else:
            return r, vv - lv
    elif op == "-":
        if lv is None:
            return l, vv + rv
        else:
            return r, lv - vv
    elif op == "*":
        if lv is None:
            return l, vv // rv
        else:
            return r, vv // lv
    elif op == "/":
        if lv is None:
            return l, vv * rv
        else:
            return r, lv // vv

def sort_by_depth(i: List[Tuple]) -> List[Tuple]:
    dep = dict(
        chain(*(
        ((exp[0], v), (exp[2], v))
        for v, exp in i if len(exp) == 3
        ))
    )
    depths = {k: depth(k, dep) for k, _ in i}
    return sorted(i, key=lambda v: depths[v[0]], reverse=True)

def part1(i: List[str]) -> int:
    values = dict()
    for v, eq in sort_by_depth(i):
        c, val = eval_eq(v, eq, values)
        values[c] = val
    return values["root"]

def part2(i: List[str]) -> int:
    values = dict()
    left = []
    for v, eq in sort_by_depth(i):
        if v == "humn":
            continue
        if v == "root":
            if eq[0] in values:
                values[eq[2]] = values[eq[0]]
            else:
                values[eq[0]] = values[eq[2]]
            break
        c, val = eval_eq(v, eq, values)
        if c is None:
            left.append((v, eq))
        else:
            values[c] = val

    for v, eq in reversed(left):
        c, val = solve_eq(v, eq,  values)
        values[c] = val

    return values["humn"]

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[str]:
    return [
        (v, exp.split(" "))
        for v, exp in
        (l.split(": ") for l in i.splitlines())
    ]

def run():
    i = default_input()
    print(part1(i))
    print(part2(i))

def test():
    i = parse("""root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32""")
    print(part1(i))
    print(part2(i))

if __name__ == "__main__":
    test()