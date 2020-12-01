from functools import partial
from collections import namedtuple, defaultdict


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

Coord = namedtuple("Coord", "x y")

DIR_UP = 0
DIR_RIGHT = 1
DIR_DOWN = 2
DIR_LEFT = 3

COLOUR_BLACK = 0
COLOUR_WHITE = 1

class PaintRobot():
	def __init__(self, proc, start_colour):
		self.p = proc
		self.direction = DIR_UP
		self.position = Coord(0, 0)

		self.painted = defaultdict(int)
		self.colour = defaultdict(int)

		self.colour[self.position] = start_colour

	def move_forward(self):
		if self.direction == DIR_UP:
			self.position = Coord(self.position.x, self.position.y - 1)
		elif self.direction == DIR_RIGHT:
			self.position = Coord(self.position.x + 1, self.position.y)
		elif self.direction == DIR_DOWN:
			self.position = Coord(self.position.x, self.position.y + 1)
		elif self.direction == DIR_LEFT:
			self.position = Coord(self.position.x - 1, self.position.y)
		else:
			raise Exception("invalid direction")

	def turn_left(self):
		if self.direction == DIR_UP:
			self.direction = DIR_LEFT
		else:
			self.direction = self.direction - 1

	def turn_right(self):
		self.direction = (self.direction + 1) % 4

	def run_bot(self):
		while self.p.state != TERMINATED:
			self.p.feed_input(self.colour[self.position])
			self.p.process()

			c = self.p.pop_output()
			if c is None:
				raise Exception("Bad colour")

			self.colour[self.position] = c
			self.painted[self.position] += 1

			d = self.p.pop_output()
			if d is None:
				raise Exception("No direction!")
			if d == 0:
				self.turn_left()
			elif d == 1:
				self.turn_right()
			else:
				raise Exception("Bad direction")


			self.move_forward()


def run():
	with open("program.txt") as f:
		in_arr = tuple(int(c) for c in f.read().split(","))
		p = Processor(in_arr)

		bot = PaintRobot(p, COLOUR_BLACK)
		bot.run_bot()

		print(len(bot.painted))

def render(image):
	for row in image:
		line = ""
		for cell in row:
			if cell == COLOUR_BLACK:
				line += u"\u001b[40m" + " "
			elif cell == COLOUR_WHITE:
				line += u"\u001b[47m" + " "
		line += u"\u001b[0m"
		print(line)

def run_part2():
	with open("program.txt") as f:
		in_arr = tuple(int(c) for c in f.read().split(","))
		p = Processor(in_arr)

		bot = PaintRobot(p, COLOUR_WHITE)
		bot.run_bot()

		max_x, max_y = 0, 0
		for p in bot.painted:
			if p.x > max_x:
				max_x = p.x
			if p.y > max_y:
				max_y = p.y

		img = [[COLOUR_BLACK for _ in range(max_x + 1)] for _ in range(max_y + 1)]
		for cell, colour in bot.colour.items():
			img[cell.y][cell.x] = colour

		render(img)
		# print(len(bot.painted))


if __name__ == "__main__":
	run()
	run_part2()
