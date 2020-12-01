from functools import partial
from itertools import permutations

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

def run_until_complete(prog):
	p = Processor(prog)
	while p.state != TERMINATED:
		p.process()
	return read_all(p)

def test_relative():
	test_cases = (
		((109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99), (109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99)),
		((1102,34915192,34915192,7,4,7,99,0), (1219070632396864,)),
		((104,1125899906842624,99), (1125899906842624,)),
	)

	for prog, expOut in test_cases:
		out = run_until_complete(prog)
		out = tuple(out)
		print(out)
		if out != expOut:
			print("Output != expected")

def run():
	with open("data.txt") as f:
		in_arr = tuple(int(c) for c in f.read().split(","))
		p = Processor(in_arr)
		p.feed_input(1)
		p.process()
		print(read_all(p))

def run_part2():
	with open("data.txt") as f:
		in_arr = tuple(int(c) for c in f.read().split(","))
		p = Processor(in_arr)
		p.feed_input(2)
		p.process()
		print(read_all(p))

if __name__ == "__main__":
	run()
	run_part2()
