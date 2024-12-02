from collections import defaultdict

def parse_input(ins):
    a, b = [], []
    for l in ins.splitlines():
        parts = l.split(" ")
        a.append(int(parts[0]))
        b.append(int(parts[-1]))
    return a, b

def run_part1(a, b):
    a.sort()
    b.sort()
    s = sum(abs(j - i) for i, j in zip(a, b))
    print(s)

def run_part2(a, b):
    freq = defaultdict(int)
    for i in b:
        freq[i] += 1

    s = sum(i * freq[i] for i in a)
    print(s)

testInput = """3   4
4   3
2   5
1   3
3   9
3   3
"""

def test_part1():
    run_part1(*parse_input(testInput))

def part1():
    run_part1(*parse_input(open("input.txt").read()))

def test_part2():
    run_part2(*parse_input(testInput))

def part2():
    run_part2(*parse_input(open("input.txt").read()))

def main():
    test_part1()
    part1()
    test_part2()
    part2()

if __name__ == "__main__":
    main()
