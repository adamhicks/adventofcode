import re

test_input = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
"""

def parse_input(s):
    return [tuple(map(int, re.split(r"[: ]+", l))) for l in s.splitlines()]

def run_part1(i):
    s = sum(eq[0] for eq in i if can_make(eq[0], eq[1:]))
    print(s)

def can_make(ans, nums, concat=False):
    stk = [(nums[0], nums[1:])]
    while stk:
        total, left = stk.pop()
        if total == ans and not left:
            return True
        if not left or total > ans:
            continue
        n = left[0]
        stk.append((total * n, left[1:]))
        stk.append((total + n, left[1:]))
        if concat:
            i = int(str(total) + str(n))
            stk.append((i, left[1:]))
    return False

def run_part2(i):
    s = sum(eq[0] for eq in i if can_make(eq[0], eq[1:], True))
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