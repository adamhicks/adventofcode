from collections import namedtuple
import heapq


Coord = namedtuple("Coord", "x y")
Coord.__repr__ = lambda c: "(%d,%d)" % c

def neighbours(c):
	return [
		Coord(c.x, c.y-1),
		Coord(c.x+1, c.y),
		Coord(c.x, c.y+1),
		Coord(c.x-1, c.y),
	]

NodeState = namedtuple("NodeState", "keys d")
Route = namedtuple("Route", "d keys")

def is_accessible(maze, c):
	if c.x < 0 or c.y < 0:
		return False
	if c.x >= len(maze[0]) or c.y >= len(maze):
		return False
	return maze[c.y][c.x] != '#'

def get_routes(maze, s):
	cur_routes = {s: Route(0, set())}
	ends = [s]

	while ends:
		c = ends.pop(0)
		r = cur_routes[c]

		for n in neighbours(c):
			if n in cur_routes:
				continue
			if not is_accessible(maze, n):
				continue

			ch = maze[c.y][c.x]
			keys_needed = set(r.keys)
			if ch.isupper():
				keys_needed.add(ch.lower())

			cur_routes[n] = Route(r.d + 1, keys_needed)
			ends.append(n)

	return {maze[c.y][c.x]: r for c, r in cur_routes.items() if maze[c.y][c.x].islower() and r.d > 0}

def convert_maze_str(maze_str):
	m = maze_str.splitlines()
	max_y, max_x = len(m), len(m[0])

	start = None
	keys = {}
	for y in range(max_y):
		for x in range(max_x):
			if m[y][x] == "@":
				start = Coord(x, y)
			elif m[y][x].islower():
				keys[m[y][x]] = Coord(x, y)

	if start is None:
		raise Exception("no start")

	maze = tuple(tuple(c if c != "@" else "." for c in l) for l in m)
	return maze, start, keys

def build_graph(maze, start, keys):
	g = {"@": get_routes(maze, start)}
	for k, c in keys.items():
		g[k] = get_routes(maze, c)
	return g

def print_graph(g):
	for k, routes in g.items():
		print("starting at %s" % k)
		for t, r in routes.items():
			keys = ""
			if r.keys:
				keys = "[%s]" % ", ".join(r.keys)
			print("-> %s = %d %s" % (t, r.d, keys))

def find_shortest_traverse(graph, keys):
	nodes = {"@": 0}
	while True:
		d, n = nodes[0]



def test_example():
	maze_str = """#########
#b.A.@.a#
#########"""
	
	maze, start, keys = convert_maze_str(maze_str)
	g = build_graph(maze, start, keys)
	print_graph(g)


def test_cases():
	cases = (
		# ("########################\n#...............b.C.D.f#\n#.######################\n#.....@.a.B.c.d.A.e.F.g#\n########################", 132),
		("#################\n#i.G..c...e..H.p#\n########.########\n#j.A..b...f..D.o#\n########@########\n#k.E..a...g..B.n#\n########.########\n#l.F..d...h..C.m#\n#################", 136),
	)
	for m_str, expD in cases:
		m, s, k = convert_maze_str(m_str)
		g = build_graph(m, s, k)
		print_graph(g)
		# s = shortest_key_path(m, s, n)
		# print(str(s.d) + " " + " ".join(s.keys))

if __name__ == '__main__':
	test_cases()
