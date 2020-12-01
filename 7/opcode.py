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

def get_param(idx, val, modes, arr):
	r = val
	if get_param_mode(modes, idx) == 0:
		r = arr[val]
		log("[{}] using value {} from {}".format(idx, r, val))
	else:
		log("[{}] using value {}".format(idx, r))
	return r

def test_get_param():
	idx = 0
	val = 4
	modes = [0, 1]
	arr = [1002, 4, 3, 4, 33]
	print(get_param(idx, val, modes, arr))

def add(cmd, modes, arr):
	read1 = get_param(0, cmd[1], modes, arr)
	read2 = get_param(1, cmd[2], modes, arr)
	# idx = get_param(2, cmd[3], modes, arr)

	arr[cmd[3]] = read1 + read2

def multiply(cmd, modes, arr):
	read1 = get_param(0, cmd[1], modes, arr)
	read2 = get_param(1, cmd[2], modes, arr)
	# idx = get_param(2, cmd[3], modes, arr)

	arr[cmd[3]] = read1 * read2

def write_val(inVal, cmd, modes, arr):
	# outIdx = get_param(0, cmd[1], modes, arr)
	arr[cmd[1]] = inVal

def read_val(cmd, modes, arr):
	return arr[cmd[1]]
	# return get_param(0, cmd[1], modes, arr)

def jump_if(cond, cmd, modes, arr):
	read1 = get_param(0, cmd[1], modes, arr)
	read2 = get_param(1, cmd[2], modes, arr)

	if cond(read1):
		return read2
	else:
		return None

jump_if_true = partial(jump_if, lambda a: a != 0)
jump_if_false = partial(jump_if, lambda a: a == 0)

def cmp(cond, cmd, modes, arr):
	read1 = get_param(0, cmd[1], modes, arr)
	read2 = get_param(1, cmd[2], modes, arr)
	# read3 = get_param(2, cmd[3], modes, arr)
	read3 = cmd[3]

	if cond(read1, read2):
		arr[read3] = 1
	else:
		arr[read3] = 0

less_than = partial(cmp, lambda a, b: a < b)
equals = partial(cmp, lambda a, b: a == b)


def print_command(cmd):
	log("command: " + ", ".join(str(c) for c in cmd))

INIT = 0
WAITING_FOR_INPUT = 2
TERMINATED = 3

class Processor():
	def __init__(self, memory):
		self.memory = list(memory)
		self.state = INIT
		self.output_list = []
		self.idx = 0
		self.input_list = []

	def feed_input(self, inVal):
		self.input_list.append(inVal)

	def get_input(self):
		if not self.input_list:
			self.state = WAITING_FOR_INPUT
			return None
		return self.input_list.pop(0)

	def pop_output(self):
		if not self.output_list:
			raise Exception("No output!")
		return self.output_list.pop(0)

	def push_output(self, outVal):
		self.output_list.append(outVal)

	def process(self):
		while True:
			idx = self.idx

			op = self.memory[idx]
			op_code = opcode(op)
			for i, val in enumerate(self.memory):
				log("{} = {}".format(i, val))
			log("instruction index " + str(idx))

			modes = param_modes(op)

			if op_code == 1:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				add(cmd, modes, self.memory)
				idx += 4
			elif op_code == 2:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				multiply(cmd, modes, self.memory)
				idx += 4
			elif op_code == 3:
				cmd = self.memory[idx:idx+2]
				print_command(cmd)

				inVal = self.get_input()
				if inVal is None:
					return

				write_val(inVal, cmd, modes, self.memory)
				idx += 2
			elif op_code == 4:
				cmd = self.memory[idx:idx+2]
				print_command(cmd)
				outVal = read_val(cmd, modes, self.memory)
				self.push_output(outVal)
				idx += 2
			elif op_code == 5:
				cmd = self.memory[idx:idx+3]
				print_command(cmd)
				retVal = jump_if_true(cmd, modes, self.memory)
				if retVal is not None:
					idx = retVal
				else:
					idx += 3
			elif op_code == 6:
				cmd = self.memory[idx:idx+3]
				print_command(cmd)
				retVal = jump_if_false(cmd, modes, self.memory)
				if retVal is not None:
					idx = retVal
				else:
					idx += 3
			elif op_code == 7:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				less_than(cmd, modes, self.memory)
				idx += 4
			elif op_code == 8:
				cmd = self.memory[idx:idx+4]
				print_command(cmd)
				equals(cmd, modes, self.memory)
				idx += 4
			elif op_code == 99:
				self.state = TERMINATED
				return
			else:
				raise Exception("invalid command")

			self.idx = idx

def run_amps(amp_code, phases):
	ex = 0
	for p in phases:
		proc = Processor(amp_code)
		proc.feed_input(p)
		proc.feed_input(ex)
		proc.process()
		ex = proc.pop_output()

	return ex

def run_feedback_amp(amp_code, phases):
	amps = []
	for p in phases:
		proc = Processor(amp_code)
		proc.feed_input(p)
		amps.append(proc)
	
	curVal = 0
	while True:
		for proc in amps:
			proc.feed_input(curVal)
			proc.process()
			curVal = proc.pop_output()

		last = amps[-1]
		if last.state == TERMINATED:
			return curVal

def find_biggest_phase(amp_code, amp_f, start, end):
	maxVal, maxPhases = 0, ()
	for p in permutations(range(start, end)):
		v = amp_f(amp_code, p)
		if v > maxVal:
			maxVal = v
			maxPhases = p
	return maxVal, maxPhases

def test_run_amps():
	arrs = [
		((3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0), (4,3,2,1,0), 43210),
		((3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0), (0,1,2,3,4), 54321),
		((3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0), (1,0,4,3,2), 65210),
	]

	for t in arrs:
		inArr, phases, outVal = t

		maxVal, maxPhases = find_biggest_phase(inArr, run_amps, 0, 5)
		if maxVal != outVal:
			print("invalid result {} != expected {}".format(maxVal, outVal))


def test_example():
	arr = (3, 16, 101, 1, 16, 16, 4, 16, 108, 8, 16, 17, 1006, 17, 0, 99, 0, 0)
	proc = Processor(arr)

	print("feeding 0")
	proc.feed_input(0)

	while proc.state != TERMINATED:
		proc.process()
		if proc.state == WAITING_FOR_INPUT:
			o = proc.pop_output()
			print("feeding " + str(o))
			proc.feed_input(o)

	print("finished " + str(proc.pop_output()))

def test_run_feedback():
	arrs = [
	    ((3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5), (9,8,7,6,5), 139629729),
		((3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4, 53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10), (9,7,8,5,6), 18216),
	]

	for t in arrs:
		inArr, phases, outVal = t

		maxVal, _ = find_biggest_phase(inArr, run_feedback_amp, 5, 10)
		if maxVal != outVal:
			print("invalid result {} != expected {}".format(maxVal, outVal))


def run():
	with open("data.txt") as inFile:
		arr = tuple(int(v) for v in inFile.read().split(','))
		maxVal, maxPhases = find_biggest_phase(arr, run_amps, 0, 5)
		print(maxVal)

def run_feedback():
	with open("data.txt") as inFile:
		arr = tuple(int(v) for v in inFile.read().split(','))
		maxVal, maxPhases = find_biggest_phase(arr, run_feedback_amp, 5, 10)
		print(maxVal)

if __name__ == "__main__":
	run()
	run_feedback()
