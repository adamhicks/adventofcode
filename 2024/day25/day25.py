test_input = """#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
"""

def parse_input(s):
    patterns = [p.splitlines() for p in s.split("\n\n")]
    locks = [p for p in patterns if all(c == "#" for c in p[0])]
    keys = [p for p in patterns if all(c == "#" for c in p[-1])]
    return locks, keys

def encode(p):
    return tuple(
        sum(
            1 if p[row][col] == "#" else 0
            for row in range(len(p))
        ) - 1 for col in range(len(p[0]))
    )

def fits(lock, key):
    return all(l + k < 6 for l, k in zip(lock, key))

def run_part1(i):
    locks, keys = i

    locks_enc = [encode(l) for l in locks]
    keys_enc = [encode(k) for k in keys]

    s = sum(1 for l in locks_enc for k in keys_enc if fits(l, k))
    print(s)

def run_part2(i):
    print("Yay!")

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