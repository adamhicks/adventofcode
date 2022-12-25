from typing import List
from os import path

def from_snafu(s: str) -> int:
    val = 0
    exp = 1
    for i, c in enumerate(reversed(s)):
        if c == "-":
            v = -1
        elif c == "=":
            v = -2
        else:
            v = int(c)
        val += v * exp
        exp *= 5
    return val

def to_snafu(i: int) -> str:
    exp = 0
    while 2 * (5**exp) < i:
        exp += 1

    opts = ((2, "2"), (1, "1"), (0, "0"), (-1, "-"), (-2, "="))

    s = ""
    v = 0
    for j in range(0, exp+1):
        e = 5 ** (exp-j)
        o = min(opts, key=lambda o: abs(i - (v + e * o[0])))
        s += o[1]
        v += o[0] * e
    return s

def part1(i: List[str]) -> str:
    v = sum(from_snafu(s) for s in i)
    return to_snafu(v)

def part2(i: List[str]) -> int:
    return len(i) ** 2

def default_input() -> str:
    fn = path.join(path.dirname(__file__), "input.txt")
    with open(fn) as i:
        return parse(i.read())

def parse(i: str) -> List[str]:
    return i.splitlines()

def run():
    i = default_input()
    print(part1(i))

def test():
    test_snafu()
    i = parse("""1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122""")
    print(part1(i))

def test_snafu():
    tc = (
        (1, "1"),
        (2, "2"),
        (3, "1="),
        (4, "1-"),
        (5, "10"),
        (6, "11"),
        (7, "12"),
        (8, "2="),
        (9, "2-"),
        (10, "20"),
        (15, "1=0"),
        (20, "1-0"),
        (2022, "1=11-2"),
        (12345, "1-0---0"),
        (314159265, "1121-1110-1=0"),
    )
    for i, exp in tc:
        s = to_snafu(i)
        if s != exp:
            print(f"{i} -> {s} != {exp}")
            return
        print(f"{i} -> {s} ✔")

if __name__ == "__main__":
    test()