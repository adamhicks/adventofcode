from collections import namedtuple
import math

Coord = namedtuple("Coord", "x y")


def in_bounds(c, max_x, max_y):
	return 0 <= c.x < max_x and 0 <= c.y < max_y

def square(depth):
	bottom_left = Coord(-depth, -depth)
	top_right = Coord(depth, depth)

	for x in range(bottom_left.x, top_right.x+1):
		yield Coord(x, bottom_left.y)
		yield Coord(x, top_right.y)

	for y in range(bottom_left.y+1, top_right.y):
		yield Coord(bottom_left.x, y)
		yield Coord(top_right.x, y)


def gen_radius(start, depth, max_x, max_y):
	for c in square(depth):
		actual_c = Coord(c.x + start.x, c.y + start.y)
		if not in_bounds(actual_c, max_x, max_y):
			continue
		yield actual_c


def lcm(a, b):
	while b != 0:
		a, b = b, a % b
	return a

def get_vector(start, target):
	c = Coord(target.x - start.x, target.y - start.y)
	m = abs(lcm(c.x, c.y))
	v = Coord(c.x / m, c.y / m)
	return v


def apply_vector(c, vector):
	return Coord(c.x + vector.x, c.y + vector.y)


def gen_line(start, vector, max_x, max_y):
	c = start
	while True:
		c = apply_vector(c, vector)
		if not in_bounds(c, max_x, max_y):
			break
		yield c


def test_gen_line():
	v = Coord(-1,-2)
	s = Coord(2,2)

	line = list(gen_line(s, v, 5, 5))
	print(line)


def gen_lines(max_x, max_y):
	zero = Coord(0, 0)
	lines = set()
	for y in range(-max_y, max_y+1):
		for x in range(-max_x, max_x+1):
			target = Coord(x, y)
			if target == zero:
				continue
			lines.add(get_vector(zero, target))
	return tuple(lines)


def calc_angle(opp, adj):
	return math.degrees(math.atan(float(opp)/float(adj)))

def get_angle(vector):
	if vector.x == 0:
		if vector.y > 0:
			return 180
		else:
			return 0
	elif vector.y == 0:
		if vector.x > 0:
			return 90
		else:
			return 270

	if vector.x > 0: # < 180
		if vector.y > 0: # > 90
			return 180 - calc_angle(vector.x, vector.y)
		else: # < 90
			return calc_angle(vector.x, -vector.y)
	else: # > 180
		if vector.y > 0: # < 270
			return 180 + calc_angle(-vector.x, vector.y)
		else: # > 270
			return 360 - calc_angle(-vector.x, -vector.y)
	raise Exception("error with angles")

def gen_vectors(max_x, max_y):
	return sorted(gen_lines(max_x, max_y), key=get_angle)

def is_asteroid(c, field):
	return field[c.y][c.x] == "#"


def print_field(field):
	for row in field:
		print("".join(row))


def lay_field(field):
	return [list(r) for r in field]

def mark_blocked(start, field):
	if not is_asteroid(start, field):
		raise Exception("not asteroid")

	field[start.y][start.x] = "S"

	height = len(field)
	width = len(field[0])

	max_dim = max(start.x + 1, width - start.x, start.y + 1, height - start.y)

	for depth in range(1, max_dim):
		for c in gen_radius(start, depth, width, height):
			if is_asteroid(c, field):
				v = get_vector(start, c)
				for b in gen_line(c, v, width, height):
					if is_asteroid(b, field):
						field[b.y][b.x] = str(depth)


def count_asteroids(field):
	total = 0
	for row in field:
		for c in row:
			if c == "#":
				total += 1
	return total


def get_all_positions(field, max_x, max_y):
	blocked = []

	for y in range(max_y):
		for x in range(max_x):

			c = Coord(x, y)

			if is_asteroid(c, field):
				this_field = lay_field(field)
				mark_blocked(c, this_field)
				blocked.append((c, this_field,))

	return blocked

def find_best_position(field, max_x, max_y):
	all_pos = ((count_asteroids(b), c) for c, b in get_all_positions(field, max_x, max_y))
	return max(all_pos, key=lambda p: p[0])


def test():
	cases = (
		((".#..#.....#####....#...##", 5, 5), 8),
		(("......#.#.#..#.#......#######..#.#.###...#..#.......#....#.##..#....#..##.#..#####...#..#..#....####", 10, 10), 33),
	)

	for case, expMax in cases:
		f, w, h = case

		rows = []
		for y in range(h):
			idx = y*w
			rows.append(tuple(c for c in f[idx:idx+w]))

		field = tuple(rows)
		print(find_best_position(field, w, h))


def zap_rocks(start, field, max_x, max_y):
	v = gen_vectors(max_x, max_y)
	warzone = lay_field(field)
	i = 1

	while True:
		pre = i
		for l in v:
			for c in gen_line(start, l, max_x, max_y):
				if is_asteroid(c, warzone):
					warzone[c.y][c.x] = str(i)
					i += 1
					break
		if pre == i:
			break
	return warzone

def find_n(warzone, n, max_x, max_y):
	c = str(n)
	for y in range(max_y):
		for x in range(max_x):
			if warzone[y][x] == c:
				return Coord(x, y)
	return None

def test_example():
	in_field = """.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##"""
	field = tuple(
		tuple(c for c in r) for r in in_field.splitlines()
	)
	w, h = len(field[0]), len(field)
	ast, pos = find_best_position(field, w, h)
	print(pos)
	z = zap_rocks(pos, field, w, h)
	print(find_n(z, 200, w, h))

def run():
	with open("field.txt") as f:
		field = tuple(
			tuple(c for c in r.strip()) for r in f.readlines()
		)
		w, h = len(field[0]), len(field)
		print(find_best_position(field, w, h))

def run_lasers():
	with open("field.txt") as f:
		field = tuple(
			tuple(c for c in r.strip()) for r in f.readlines()
		)
		w, h = len(field[0]), len(field)
		_, pos = find_best_position(field, w, h)
		z = zap_rocks(pos, field, w, h)
		print(find_n(z, 200, w, h))


if __name__ == "__main__":
	run()
	run_lasers()
