class Planet(object):
	def __init__(self, name, parent):
		self.name = name
		self.children = []
		self.orbit_count = 0
		self.parent = parent

	def __repr__(self):
		return self.name

def print_planet(p):
	print("{} -> {}[{}] -> {}".format(p.parent, p.name, p.orbit_count, ", ".join(p.children)))

def parse_orbits(orbits):
	solar = {"COM": Planet("COM", None)}
	for a, b in orbits:
		orbitee = solar.get(a)
		if orbitee is None:
			orbitee = Planet(a, None)
			solar[a] = orbitee

		orbiter = solar.get(b)
		if orbiter is None:
			orbiter = Planet(b, orbitee)
			solar[b] = orbiter
		elif orbiter.parent is None:
			orbiter.parent = orbitee

		orbitee.children.append(b)
	return solar

def process_orbits(solar):
	queue = [solar["COM"]]
	while queue:
		p = queue.pop()
		for n in p.children:
			c = solar[n]
			c.orbit_count = p.orbit_count + 1
			queue.append(c)

def create_solar(in_str):
	orbits = [
		line.split(')')
		for line in in_str.splitlines()
	]
	solar = parse_orbits(orbits)
	process_orbits(solar)
	return solar


def count_orbits(solar):
	total = 0
	for n, p in solar.iteritems():
		total += p.orbit_count
	return total

def get_path(p):
	current = p
	path = []
	while current is not None:
		path.insert(0, current)
		current = current.parent
	return path

def find_join(path_one, path_two):
	for p in reversed(path_two):
		if p in path_one:
			return p
	return None

def find_path(solar, start, end):
	p_one = get_path(solar[start])
	p_two = get_path(solar[end])

	join = find_join(p_one, p_two)
	if join is None:
		raise Exception("No join found")

	pre = p_one[p_one.index(join):]
	post = p_two[p_two.index(join)+1:]
	return list(reversed(pre)) + post
	
def transfers_required(path):
	return len(path) - 3

def test():
	in_str = """COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L"""
	solar = create_solar(in_str)
	print(count_orbits(solar))

def test_path():
	in_str = """COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN"""
	solar = create_solar(in_str)
	p = find_path(solar, "YOU", "SAN")
	print(transfers_required(p))

def run_part1():
	with open("data.txt") as f:
		solar = create_solar(f.read())
		print(count_orbits(solar))

def run_part2():
	with open("data.txt") as f:
		solar = create_solar(f.read())
		p = find_path(solar, "YOU", "SAN")
		print(transfers_required(p))

if __name__ == "__main__":
	run_part1()
	run_part2()
