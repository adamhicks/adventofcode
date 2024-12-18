from collections import namedtuple

test_input = """##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
"""

def parse_input(s):
    grid, moves = s.split("\n\n")
    return grid.splitlines(), moves.replace("\n", "")

Point = namedtuple("Point", ["x", "y"])

def add(a, b):
    return Point(a.x + b.x, a.y + b.y)

def run_part1(i):
    grid, moves = i
    things = {Point(x, y): c for y, row in enumerate(grid) for x, c in enumerate(row)}

    for m in moves:
        b = find_bot(things)
        d = move_dir(m)

        s = find_space(things, b, d)
        if s is None:
            continue
        things[s] = "O"
        things[b] = "."
        things[add(b, d)] = "@"

    s = sum(100 * p.y + p.x for p, c in things.items() if c == "O")
    print(s)

def find_bot(things):
    return next(p for p, c in things.items() if c == "@")

def move_dir(c):
    if c == "^":
        return Point(0, -1)
    elif c == "v":
        return Point(0, 1)
    elif c == ">":
        return Point(1, 0)
    elif c == "<":
        return Point(-1, 0)
    raise ValueError("invalid move", c)

def find_space(things, p, d):
    while things.get(p) in ("O", "@"):
        p = add(p, d)
    if things.get(p) == ".":
        return p
    else:
        return None

def run_part2(i):
    grid, moves = i
    things = dict()
    for y, row in enumerate(grid):
        for x, c in enumerate(row):
            a, b = Point(2*x, y), Point(2*x+1, y)
            if c == "#":
                things[a], things[b] = "#", "#"
            elif c == "O":
                things[a], things[b] = "[", "]"
            elif c == ".":
                things[a], things[b] = ".", "."
            elif c == "@":
                things[a], things[b] = "@", "."
    for m in moves:
        b = find_bot(things)
        d = move_dir(m)
        boxes = push_boxes(things, b, d)
        if boxes is not None:
            todo = dict()
            for p in boxes:
                todo[add(p, d)] = things.get(p)
                things[p] = "."
            for p, c in todo.items():
                things[p] = c
            things[b] = "."
            things[add(b, d)] = "@"

    s = sum(100 * p.y + p.x for p, c in things.items() if c == "[")
    print(s)

def push_boxes(things, b, d):
    face = set()
    face.add(b)
    boxes = set()
    while True:
        face = {add(p, d) for p in face}
        if all(things.get(p) == "." for p in face):
            return boxes
        if any(things.get(p) == "#" for p in face):
            return None
        nf = set()
        for n in face:
            c = things.get(n)
            if c == "[":
                nf.add(n)
                if d.x == 0:
                    nf.add(n._replace(x=n.x+1))
            elif c == "]":
                nf.add(n)
                if d.x == 0:
                    nf.add(n._replace(x=n.x-1))
        boxes.update(nf)
        face = nf

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
