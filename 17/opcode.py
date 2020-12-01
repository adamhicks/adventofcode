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

Coord = namedtuple("Coord", "x y")
Coord.__str__ = lambda c: "%d, %d" % (c)

def gen_neighbours(c):
	yield Coord(c.x - 1, c.y)
	yield Coord(c.x + 1, c.y)
	yield Coord(c.x, c.y - 1)
	yield Coord(c.x, c.y + 1)

def find(s, t):
	for y, row in enumerate(s):
		for x, c in enumerate(row):
			if c == t:
				return Coord(x, y)
	return None

DIR_NORTH = 0
DIR_EAST = 1
DIR_SOUTH = 2
DIR_WEST = 3

TURN_LEFT = "L"
TURN_RIGHT = "R"

def new_coord(c, d):
	if d == DIR_NORTH:
		return Coord(c.x, c.y-1)
	if d == DIR_EAST:
		return Coord(c.x+1, c.y)
	if d == DIR_SOUTH:
		return Coord(c.x, c.y+1)
	if d == DIR_WEST:
		return Coord(c.x-1, c.y)
	raise Exception("invalid direction")

def turn(current, turn):
	new_dir = current
	if turn == TURN_RIGHT:
		new_dir += 1
	elif turn == TURN_LEFT:
		new_dir -= 1
	else:
		raise Exception("bad turn")
	return new_dir % 4

def can_move(s, c):
	max_x, max_y = len(s[0]), len(s)
	if not(0 <= c.x < max_x):
		return False 
	if not(0 <= c.y < max_y):
		return False
	return s[c.y][c.x] == "#"


def solve_scaffold(s):
	max_x, max_y = len(s[0]), len(s)
	start = find(s, '^')
	if start is None:
		raise Exception("no start")
	d = DIR_NORTH
	l = 0
	c = start
	move_cmds = []
	while True:
		n = new_coord(c, d)
		if can_move(s, n):
			l += 1
			c = n
			continue
		if l > 0:
			move_cmds.append(l)
			l = 0

		left_d = turn(d, TURN_LEFT)
		left_n = new_coord(c, left_d)
		if can_move(s, left_n):
			move_cmds.append(TURN_LEFT)
			d = left_d
			continue

		right_d = turn(d, TURN_RIGHT)
		right_n = new_coord(c, right_d)
		if can_move(s, right_n):
			move_cmds.append(TURN_RIGHT)
			d = right_d
			continue
		break

	full_cmd = ",".join(str(c) for c in move_cmds) + ","
	print(full_cmd[:-1])
	ans = decompose_path(full_cmd)
	if ans is None:
		raise Exception("no answer")
	cmd_a, cmd_b, cmd_c = ans

	print("A = " + cmd_a[:-1])
	print("B = " + cmd_b[:-1])
	print("C = " + cmd_c[:-1])
	print(recompose_path(full_cmd, cmd_a, cmd_b, cmd_c)[:-1])

def decompose_path(cmd):
	for a_i in range(1, 20):
		cmd_a = cmd[:a_i]
		cmd_sub_a = cmd_without(cmd, cmd_a)

		for b_i in range(1, 20):
			cmd_b = cmd_sub_a[:b_i]
			cmd_sub_b = cmd_without(cmd_sub_a, cmd_b)

			for c_i in range(1, 20):
				cmd_c = cmd_sub_b[:c_i]

				final_left = cmd_without(cmd_sub_b, cmd_c)
				if final_left == "":
					return cmd_a, cmd_b, cmd_c

def recompose_path(cmd, a,b,c):
	return cmd.replace(a, "A,").replace(b, "B,").replace(c, "C,")

def cmd_without(cmd, sub):
	return cmd.replace(sub, "")

def dissect_cmd(cmd):
	sub_cmds = ["", "", ""]
	max_s = 20


class ScaffoldRobot():
	def __init__(self, program):
		self.p = program

	def run(self):
		while self.p.state != TERMINATED:
			self.p.process()

			s = ""
			for c in read_all(self.p):
				if c >= 256:
					print c
					return
				s += chr(c)
			print s
			s_in = raw_input()
			for c in s_in:
				self.p.feed_input(ord(c))
			self.p.feed_input(ord('\n'))

	def load_scaffold(self):
		self.p.process()
		s = "".join(chr(c) for c in read_all(self.p))
		self.scaffold = [[c for c in l] for l in s.strip().splitlines()]

	def print_scaffold(self):
		for l in self.scaffold:
			print("".join(l))

	def get_intersections(self):
		max_y = len(self.scaffold)
		max_x = len(self.scaffold[0])

		cross = []
		for y in range(1, max_y - 1):
			for x in range(1, max_x - 1):
				if self.scaffold[y][x] != '#':
					continue
				c = Coord(x, y)
				if all(self.scaffold[n.y][n.x] == '#' for n in gen_neighbours(c)):
					self.scaffold[y][x] = 'O'
					cross.append(c)
		return cross


def run():
	with open("scaffold.txt") as s:
		in_arr = list(int(c) for c in s.read().split(","))
		# in_arr[0] += 1

		p = Processor(in_arr)
		bot = ScaffoldRobot(p)
		bot.load_scaffold()
		# bot.print_scaffold()
		solve_scaffold(bot.scaffold)
		# curses.wrapper(bot.play_game)
		# bot.run()

# Solution to part 2
# A,B,A,C,B,A,B,C,C,B
# L,12,L,12,R,4
# R,10,R,6,R,4,R,4
# R,6,L,12,L,12

if __name__ == "__main__":
	run()
