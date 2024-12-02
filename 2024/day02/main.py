from itertools import count


def parse_input(ins):
    return [[int(i) for i in l.split(" ")] for l in ins.splitlines()]

def run_part1(i):
    print(sum(safe(r) for r in i))

def safe(r):
    diffs = [r[i+1]-r[i] for i in range(len(r)-1)]
    inc = diffs[0] > 0
    for d in diffs:
        if d == 0:
            return False
        this_inc = d > 0
        if this_inc != inc:
            return False
        if abs(d) > 3:
            return False
    return True


def safe_damp(r):
    if safe(r):
        return True
    for i in range(len(r)):
        l = r[:i] + r[i+1:]
        if safe(l):
            return True
    return False

def run_part2(i):
    print(sum(safe_damp(r) for r in i))

testInput = """7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
"""

def test_part1():
    run_part1(parse_input(testInput))

def part1():
    run_part1(parse_input(open("input.txt").read()))

def test_part2():
    run_part2(parse_input(testInput))

def part2():
    run_part2(parse_input(open("input.txt").read()))

def main():
    test_part1()
    part1()
    test_part2()
    part2()

if __name__ == "__main__":
    main()
