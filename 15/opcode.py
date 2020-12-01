 #!/usr/bin/python
 # -*- coding: utf-8 -*-

from functools import partial
from collections import namedtuple, defaultdict
import curses
import time

def log(s):
	if False:
		print(s)

def opcode(cmdVal):
	return cmdVal % 100

def param_modes(cmdVal):
	cmdVal = cmdVal / 100
	modes = []
	while cmdVal > 0:
		modes.append(cmdVal%10)
		cmdVal = cmdVal / 10
	return modes

def test_param_modes():
	print(param_modes(1002))


def get_param_mode(modes, idx):
	if idx >= len(modes):
		return 0
	return modes[idx]

def test_get_mode():
	print(get_param_mode([0, 1], 0))

def get_address(idx, val, modes, base):
	r = val
	p_mode = get_param_mode(modes, idx)
	if p_mode == 2:
		r += base
	return r

def get_param(idx, val, modes, arr, base):
	r = val
	p_mode = get_param_mode(modes, idx)
	if p_mode == 0:
		r = arr[val]
		log("[{}] using value {} from {}".format(idx, r, val))
	elif p_mode == 1:
		log("[{}] using value {}".format(idx, r))
	elif p_mode == 2:
		r = arr[val+base]
		log("[{}] using value {} from {}+{}".format(idx, r, val, base))
	else:
		raise Exception("Unknown parameter mode")
	return r

def test_get_param():
	idx = 0
	val = 4
	modes = [0, 1]
	arr = [1002, 4, 3, 4, 33]
	print(get_param(idx, val, modes, arr, base))

def add(cmd, modes, arr, base):
	read1 = get_param(0, cmd[1], modes, arr, base)
	read2 = get_param(1, cmd[2], modes, arr, base)
	idx = get_address(2, cmd[3], modes, base)

	arr[idx] = read1 + read2

def multiply(cmd, modes, arr, base):
	read1 = get_param(0, cmd[1], modes, arr, base)
	read2 = get_param(1, cmd[2], modes, arr, base)
	idx = get_address(2, cmd[3], modes, base)

	arr[idx] = read1 * read2

def write_val(inVal, cmd, modes, arr, base):
	idx = get_address(0, cmd[1], modes, base)
	arr[idx] = inVal

def read_val(cmd, modes, arr, base):
	return get_param(0, cmd[1], modes, arr, base)

def jump_if(cond, cmd, modes, arr, base):
	read1 = get_param(0, cmd[1], modes, arr, base)
	read2 = get_param(1, cmd[2], modes, arr, base)

	if cond(read1):
		return read2
	else:
		return None

jump_if_true = partial(jump_if, lambda a: a != 0)
jump_if_false = partial(jump_if, lambda a: a == 0)

def cmp(cond, cmd, modes, arr, base):
	read1 = get_param(0, cmd[1], modes, arr, base)
	read2 = get_param(1, cmd[2], modes, arr, base)
	idx = get_address(2, cmd[3], modes, base)

	if cond(read1, read2):
		arr[idx] = 1
	else:
		arr[idx] = 0

less_than = partial(cmp, lambda a, b: a < b)
equals = partial(cmp, lambda a, b: a == b)

def alter_base(cmd, modes, arr, base):
	return base + get_param(0, cmd[1], modes, arr, base)

def print_command(cmd):
	log("command: " + ", ".join(str(c) for c in cmd))

INIT = 0
WAITING_FOR_INPUT = 2
TERMINATED = 3

class Processor():
	def __init__(self, memory):
		self.memory = list(memory) + [0 for _ in range(4096)] # 4K mem
		self.state = INIT
		self.output_list = []
		self.idx = 0
		self.input_list = []
		self.relative_base = 0

	def feed_input(self, inVal):
		self.input_list.append(inVal)

	def get_input(self):
		if not self.input_list:
			self.state = WAITING_FOR_INPUT
			return None
		return self.input_list.pop(0)

	def pop_output(self):
		if not self.output_list:
			return None
		return self.output_list.pop(0)

	def push_output(self, outVal):
		self.output_list.append(outVal)

	def process(self):
		while True:
			idx = self.idx

			op = self.memory[idx]
			op_code = opcode(op)
			log("instruction index " + str(idx))

			modes = param_modes(op)

			if op_code == 1:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				add(cmd, modes, self.memory, self.relative_base)
				idx += 4
			elif op_code == 2:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				multiply(cmd, modes, self.memory, self.relative_base)
				idx += 4
			elif op_code == 3:
				cmd = self.memory[idx:idx+2]
				print_command(cmd)

				inVal = self.get_input()
				if inVal is None:
					return

				write_val(inVal, cmd, modes, self.memory, self.relative_base)
				idx += 2
			elif op_code == 4:
				cmd = self.memory[idx:idx+2]
				print_command(cmd)
				outVal = read_val(cmd, modes, self.memory, self.relative_base)
				self.push_output(outVal)
				idx += 2
			elif op_code == 5:
				cmd = self.memory[idx:idx+3]
				print_command(cmd)
				retVal = jump_if_true(cmd, modes, self.memory, self.relative_base)
				if retVal is not None:
					idx = retVal
				else:
					idx += 3
			elif op_code == 6:
				cmd = self.memory[idx:idx+3]
				print_command(cmd)
				retVal = jump_if_false(cmd, modes, self.memory, self.relative_base)
				if retVal is not None:
					idx = retVal
				else:
					idx += 3
			elif op_code == 7:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				less_than(cmd, modes, self.memory, self.relative_base)
				idx += 4
			elif op_code == 8:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				equals(cmd, modes, self.memory, self.relative_base)
				idx += 4
			elif op_code == 9:
				cmd = self.memory[idx:idx+2]
				print_command(cmd)
				self.relative_base = alter_base(cmd, modes, self.memory, self.relative_base)
				idx += 2
			elif op_code == 99:
				self.state = TERMINATED
				return
			else:
				raise Exception("invalid command")

			self.idx = idx

def read_all(p):
	ret = []
	while True:
		o = p.pop_output()
		if o is None:
			break
		ret.append(o)
	return ret

# End of computer

DIR_NORTH = 1
DIR_SOUTH = 2
DIR_WEST = 3
DIR_EAST = 4

RES_BLOCKED = 0
RES_MOVED = 1
RES_GOAL = 2

Coord = namedtuple("Coord", "x y")

def new_coord(c, dir):
	if dir == DIR_NORTH:
		return Coord(c.x, c.y - 1)
	elif dir == DIR_SOUTH:
		return Coord(c.x, c.y + 1)
	elif dir == DIR_WEST:
		return Coord(c.x - 1, c.y)
	elif dir == DIR_EAST:
		return Coord(c.x + 1, c.y)
	raise Exception("invalid dir")

def inverse(dir):
	if dir == DIR_NORTH:
		return DIR_SOUTH
	elif dir == DIR_SOUTH:
		return DIR_NORTH
	elif dir == DIR_WEST:
		return DIR_EAST
	elif dir == DIR_EAST:
		return DIR_WEST
	raise Exception("invalid dir")

def next_dir(dir):
	if dir == DIR_NORTH:
		return DIR_EAST
	elif dir == DIR_EAST:
		return DIR_SOUTH
	elif dir == DIR_SOUTH:
		return DIR_WEST
	elif dir == DIR_WEST:
		return None
	raise Exception("invalid dir")

def str_dir(dir):
	if dir == DIR_NORTH:
		return "N"
	elif dir == DIR_EAST:
		return "E"
	elif dir == DIR_SOUTH:
		return "S"
	elif dir == DIR_WEST:
		return "W"
	raise Exception("invalid dir")


class OxygenRobot():
	def __init__(self, program):
		self.trace = open("trace.log", "w")

		self.p = program
		z = Coord(0, 0)
		self.grid = {z: "S"}
		self.bot_position = z

	def go(self, direction):
		self.p.feed_input(direction)
		self.p.process()
		res = self.p.pop_output()

		g = new_coord(self.bot_position, direction)
		if res == RES_BLOCKED:
			self.grid[g] = "#"
		elif res == RES_MOVED:
			self.grid[g] = "."
		elif res == RES_GOAL:
			self.grid[g] = "O"

		if res != RES_BLOCKED:
			self.bot_position = g


		self.log("tried to go %s got %s" % (str_dir(direction), self.grid[g]))

		return res

	def draw_ch(self, win, co, ch):
		max_y, max_x = win.getmaxyx()
		y_off = max_y // 2
		x_off = max_x // 2
		y = co.y + y_off
		x = co.x + x_off

		if not(0 < x < max_x):
			return

		if not(0 < y < max_y):
			return

		win.addch(y, x, ord(ch))

	def render(self, win):
		win.clear()

		for c, t in self.grid.items():
			self.draw_ch(win, c, t)

		self.draw_ch(win, self.bot_position, "D")

	def log(self, s):
		self.trace.write(s + "\n")

	def explore(self, d):
		c = new_coord(self.bot_position, d)
		if c in self.grid:
			return False

		r = self.go(d)
		return r != RES_BLOCKED

	def path_forward(self, last_d):
		next_d = DIR_NORTH

		while next_d is not None:
			if next_d == inverse(last_d):
				next_d = next_dir(next_d)
				continue

			if self.explore(next_d):
				return next_d

			next_d = next_dir(next_d)

		return None


	def auto_explore(self, w):
		w.nodelay(1)

		path = [self.path_forward(DIR_NORTH)]

		while path:
			self.render(w)
			w.refresh()

			if w.getch() == ord('q'):
				break

			s = " ".join(str_dir(p) for p in path)
			self.log("@[%d,%d] from - %s" % (self.bot_position.x, self.bot_position.y, s))

			last_d = path[-1]
			d = self.path_forward(last_d)
			if d is not None:
				path.append(d)
			else:
				self.go(inverse(last_d))
				path.pop()
			# time.sleep(0.01)

		w.nodelay(0)

	def fill_oxygen(self, w):
		w.nodelay(1)

		start_o = None
		for k, v in self.grid.items():
			if v == "O":
				start_o = k
				break

		if start_o is None:
			raise Exception("No starting oxygen")

		frontier = [start_o]
		tick = 0

		while frontier:
			self.render(w)
			w.refresh()
			if w.getch() == ord('q'):
				break

			tick += 1
			new_frontier = []

			for f in frontier:
				d = DIR_NORTH
				while d is not None:
					neigh = new_coord(f, d)
					if self.grid[neigh] == ".":
						self.grid[neigh] = "O"
						new_frontier.append(neigh)
					d = next_dir(d)

			frontier = new_frontier

		self.log("filled with oxygen in %d minutes" % (tick - 1))
		w.nodelay(0)


	def play_game(self, w):
		curses.curs_set(0)
		w.clear()

		while self.p.state != TERMINATED:
			self.render(w)
			w.refresh()

			c = w.getch()
			if c == curses.KEY_UP:
				self.go(DIR_NORTH)
			elif c == curses.KEY_DOWN:
				self.go(DIR_SOUTH)
			elif c == curses.KEY_LEFT:
				self.go(DIR_WEST)
			elif c == curses.KEY_RIGHT:
				self.go(DIR_EAST)
			elif c == ord('q'):
				break
			elif c == ord('a'):
				self.auto_explore(w)
			elif c == ord('o'):
				self.fill_oxygen(w)
			else:
				continue

		self.trace.close()


def run():
	with open("program.txt") as f:
		in_arr = list(int(c) for c in f.read().split(","))
		p = Processor(in_arr)

		bot = OxygenRobot(p)
		curses.wrapper(bot.play_game)


if __name__ == "__main__":
	run()
