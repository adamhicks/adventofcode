from collections import defaultdict

test_input = """3   4
4   3
2   5
1   3
3   9
3   3
"""

def parse_input(s):
    a, b = [], []
    for l in s.splitlines():
        parts = l.split(" ")
        a.append(int(parts[0]))
        b.append(int(parts[-1]))
    return a, b

def run_part1(i):
    a, b = i
    a.sort()
    b.sort()
    s = sum(abs(j - i) for i, j in zip(a, b))
    print(s)

def run_part2(i):
    a, b = i
    freq = defaultdict(int)
    for i in b:
        freq[i] += 1

    s = sum(i * freq[i] for i in a)
    print(s)

### Generated by start script

def test_part1():
    run_part1(parse_input(test_input))

def part1():
    run_part1(parse_input(open("input.txt").read()))

def test_part2():
    run_part2(parse_input(test_input))

def part2():
    run_part2(parse_input(open("input.txt").read()))

def main():
    print("=== running part 1 test ===")
    test_part1()
    print("=== running part 1 ===")
    part1()
    print("=== running part 2 test ===")
    test_part2()
    print("=== running part 2 ===")
    part2()
    print("=== ===")

if __name__ == "__main__":
    main()
