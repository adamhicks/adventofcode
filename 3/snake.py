from collections import namedtuple

Node = namedtuple("Node", "x y")
Line = namedtuple("Line", "nodes")
Meeting = namedtuple("Meeting", "node dist1 dist2")

def is_horizontal(start, end):
	return start.y == end.y

def is_vertical(start, end):
	return start.x == end.x

def x_marks(vert_start, vert_end, horiz_start, horiz_end):
	if horiz_start.x > horiz_end.x:
		horiz_start, horiz_end = horiz_end, horiz_start
	if vert_start.y > vert_end.y:
		vert_start, vert_end = vert_end, vert_start

	if not (vert_start.y < horiz_start.y <= vert_end.y):
		return None
	if not (horiz_start.x < vert_start.x <= horiz_end.x):
		return None
	
	return Node(x=vert_start.x, y=horiz_start.y)

def crossing_point(a_start, a_end, b_start, b_end):
	if is_horizontal(a_start, a_end) and is_vertical(b_start, b_end):
		return x_marks(b_start, b_end, a_start, a_end)

	if is_vertical(a_start, a_end) and is_horizontal(b_start, b_end):
		return x_marks(a_start, a_end, b_start, b_end)

	return None

def meeting_points(line1, line2):
	meets = []
	line1_dist = 0
	for idx1 in range(len(line1) - 1):
		a_start, a_end = line1[idx1], line1[idx1+1]
		line2_dist = 0

		for idx2 in range(len(line2) - 1):
			b_start, b_end = line2[idx2], line2[idx2+1]
			m = crossing_point(a_start, a_end, b_start, b_end)
			if m is not None:
				dist1 = line1_dist + cardinal(a_start, m)
				dist2 = line2_dist + cardinal(b_start, m)
				meets.append(Meeting(m, dist1, dist2))

			line2_dist += cardinal(b_start, b_end)

		line1_dist += cardinal(a_start, a_end)
	return meets

def construct_nodes(cmdLine):
	current = Node(0, 0)
	line = [current]
	for c in cmdLine.split(","):
		direction = c[0]
		mag = int(c[1:])

		if direction == "L":
			current = Node(current.x - mag, current.y)
		elif direction == "R":
			current = Node(current.x + mag, current.y)
		elif direction == "U":
			current = Node(current.x, current.y + mag)
		elif direction == "D":
			current = Node(current.x, current.y - mag)
		else:
			raise Exception("unknown direction")
		line.append(current)
	return line

def cardinal(node1, node2):
	x = abs(node2.x - node1.x)
	y = abs(node2.y - node1.y)
	return x + y

def min_cardinal(commands):
	line1 = construct_nodes(commands[0])
	line2 = construct_nodes(commands[1])
	distance = min(
		cardinal(m.node, Node(0, 0)) for m in meeting_points(line1, line2)
	)
	return distance

def min_signal(commands):
	line1 = construct_nodes(commands[0])
	line2 = construct_nodes(commands[1])
	distance = min(
		m.dist1 + m.dist2 for m in meeting_points(line1, line2)
	)
	return distance

def test_min_cardinal():
	lines = [
		(("R8,U5,L5,D3", "U7,R6,D4,L4"), 6),
		(("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"), 159),
		(("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"), 135),
	]
	for cmds, exp in lines:
		res = min_cardinal(cmds)
		if res != exp:
			print("{} = {}, expected {}".format(cmds, res, exp))

def test_min_signal():
	lines = [
		(("R8,U5,L5,D3", "U7,R6,D4,L4"), 30),
		(("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"), 610),
		(("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"), 410),
	]
	for cmds, exp in lines:
		res = min_signal(cmds)
		if res != exp:
			print("{} = {}, expected {}".format(cmds, res, exp))


def run_cross(func):
	with open("data.txt") as snake_file:
		m = func(snake_file.readlines())
		print(m)

if __name__ == "__main__":
	run_cross(min_cardinal)
	run_cross(min_signal)