from collections import namedtuple
import re
import fractions


Coord = namedtuple("Coord", "x y z")

Coord.__str__ = lambda c: "(" + str(c.x)+", " +str(c.y) +", " + str(c.z)+ ")"

ZERO_V = Coord(0, 0, 0)

Planet = namedtuple("Planet", "position velocity")

def move_planet(p):
	new_pos = apply_vector(p.position, p.velocity)
	return Planet(new_pos, p.velocity)
Planet.move = move_planet

def gravitate(p, o):
	new_x = p.velocity.x + cmp(o.position.x, p.position.x)
	new_y = p.velocity.y + cmp(o.position.y, p.position.y)
	new_z = p.velocity.z + cmp(o.position.z, p.position.z)
	v = Coord(new_x, new_y, new_z)
	return Planet(p.position, v)
Planet.gravitate = gravitate

def energy(p):
	pot = sum(abs(i) for i in p.position)
	kin = sum(abs(i) for i in p.velocity)
	return pot * kin
Planet.energy = energy

Planet.__str__ = lambda p: "p=" + str(p.position) + ", v=" + str(p.velocity) + ", e=" + str(p.energy())


def apply_vector(pos, v):
	return Coord(pos.x + v.x, pos.y + v.y, pos.z + v.z)

def apply_gravity(planets, idx):
	p = planets[idx]
	x, y, z = p.velocity

	for i, o in enumerate(planets):
		if i == idx:
			continue
		x += cmp(o.position.x, p.position.x)
		y += cmp(o.position.y, p.position.y)
		z += cmp(o.position.z, p.position.z)

	return Planet(p.position, Coord(x, y, z))

def step_planets(planets):
	return tuple(
		apply_gravity(planets, i).move()
		for i in range(len(planets))
	)

def print_planets(planets):
	for p in planets:
		print(p)

def total_energy(planets):
	return sum(p.energy() for p in planets)

def run_n_steps(planets, n):
	print("after 0 steps")
	print_planets(planets)

	for i in range(n):
		planets = step_planets(planets)
		# print("after " + str(i) + " steps")
		# print_planets(planets)

	print("")

	print_planets(planets)
	print(total_energy(planets))

def run_until_match(planets, log_n=None, max_n=None):
	initial = planets
	i = 0
	while True:
		planets = step_planets(planets)
		i += 1

		if log_n is not None and i % log_n == 0:
			print("ran " + str(i) + " steps...")
		if max_n is not None and i == max_n:
			return None

		if planets == initial:
			return i

def test_example():
	planets = (
		Planet(Coord(-1, 0, 2), ZERO_V),
		Planet(Coord(2, -10, -7), ZERO_V),
		Planet(Coord(4, -8, 8), ZERO_V),
		Planet(Coord(3, 5, -1), ZERO_V),
	)
	run_n_steps(planets, 10)

def test_example_2():
	planets = (
		Planet(Coord(-8, -10, 0), ZERO_V),
		Planet(Coord(5, 5, 10), ZERO_V),
		Planet(Coord(2, -7, 3), ZERO_V),
		Planet(Coord(9, -8, -3), ZERO_V),
	)
	run_n_steps(planets, 100)

NUM_PARSER = re.compile(r"(-?\d+)")

def parse_pos(line):
	p = re.findall(NUM_PARSER, line)
	return Coord(*(int(i) for i in p))

def run():
	with open("data.txt") as m:
		planets = tuple(
			Planet(parse_pos(l), ZERO_V) for l in m.readlines()
		)
		run_n_steps(planets, 1000)

def test_part_2_example_1():
	planets = (
		Planet(Coord(-1, 0, 2), ZERO_V),
		Planet(Coord(2, -10, -7), ZERO_V),
		Planet(Coord(4, -8, 8), ZERO_V),
		Planet(Coord(3, 5, -1), ZERO_V),
	)
	run_until_match(planets)

def test_part_2_example_2():
	planets = (
		Planet(Coord(-8, -10, 0), ZERO_V),
		Planet(Coord(5, 5, 10), ZERO_V),
		Planet(Coord(2, -7, 3), ZERO_V),
		Planet(Coord(9, -8, -3), ZERO_V),
	)
	run_until_match(planets, 10000)

def run_part_2():
	with open("data.txt") as m:
		planets = tuple(
			Planet(parse_pos(l), ZERO_V) for l in m.readlines()
		)
		p_x = tuple(
			Planet(Coord(p.position.x, 0, 0), ZERO_V) for p in planets
		)
		p_y = tuple(
			Planet(Coord(0, p.position.y, 0), ZERO_V) for p in planets
		)
		p_z = tuple(
			Planet(Coord(0, 0, p.position.z), ZERO_V) for p in planets
		)
		mul_x = run_until_match(p_x)
		mul_y = run_until_match(p_y)
		mul_z = run_until_match(p_z)

		def _lcm(a, b):
			return a * b // fractions.gcd(a, b)

		answer = _lcm(_lcm(mul_x, mul_y), mul_z)
		print(answer)

def profile():
	planets = (
		Planet(Coord(-8, -10, 0), ZERO_V),
		Planet(Coord(5, 5, 10), ZERO_V),
		Planet(Coord(2, -7, 3), ZERO_V),
		Planet(Coord(9, -8, -3), ZERO_V),
	)

	import cProfile, pstats, StringIO
	pr = cProfile.Profile()
	pr.enable()

	run_until_match(planets, log_n=1000, max_n=100000)

	pr.disable()
	s = StringIO.StringIO()
	sortby = 'cumtime'
	ps = pstats.Stats(pr, stream=s).sort_stats(sortby)
	ps.print_stats()
	print s.getvalue()


if __name__ == "__main__":
	run()
	run_part_2()