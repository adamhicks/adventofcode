from functools import partial

def log(s):
	if True:
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
	return get_param(0, cmd[1], modes, arr)

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

def process_array(arr, idVal):
	write_id = partial(write_val, idVal)

	idx = 0
	exit_code = 0

	while True:
		log(str(idx))
		op = arr[idx]
		op_code = opcode(op)
		for i, val in enumerate(arr):
			log("{} = {}".format(i, val))

		modes = param_modes(op)

		if op_code == 1:
			cmd = arr[idx:idx+4]
			print_command(cmd)
			add(cmd, modes, arr)
			idx += 4
		elif op_code == 2:
			cmd = arr[idx:idx+4]
			print_command(cmd)
			multiply(cmd, modes, arr)
			idx += 4
		elif op_code == 3:
			cmd = arr[idx:idx+2]
			print_command(cmd)
			write_id(cmd, modes, arr)
			idx += 2
		elif op_code == 4:
			cmd = arr[idx:idx+2]
			print_command(cmd)
			exit_code = read_val(cmd, modes, arr)
			idx += 2
		elif op_code == 5:
			cmd = arr[idx:idx+3]
			print_command(cmd)
			retVal = jump_if_true(cmd, modes, arr)
			if retVal is not None:
				idx = retVal
			else:
				idx += 3
		elif op_code == 6:
			cmd = arr[idx:idx+3]
			print_command(cmd)
			retVal = jump_if_false(cmd, modes, arr)
			if retVal is not None:
				idx = retVal
			else:
				idx += 3
		elif op_code == 7:
			cmd = arr[idx:idx+4]
			print_command(cmd)
			less_than(cmd, modes, arr)
			idx += 4
		elif op_code == 8:
			cmd = arr[idx:idx+4]
			print_command(cmd)
			equals(cmd, modes, arr)
			idx += 4
		elif op_code == 99:
			return exit_code, tuple(arr)
		else:
			raise Exception("invalid command")

def test():
	arrs = [
		# ((1002,4,3,4,33), (1002,4,3,4,99)),
		((1105, 0, 0, 1101, 1, 1, 0, 99), (2, 0, 0, 1101, 1, 1, 0, 99)), 
		((1105, 1, 7, 1101, 1, 1, 0, 99), (1105, 1, 7, 1101, 1, 1, 0, 99)),
		((1106, 0, 7, 1101, 1, 1, 0, 99), (1106, 0, 7, 1101, 1, 1, 0, 99)), 
		((1106, 1, 7, 1101, 1, 1, 0, 99), (2, 1, 7, 1101, 1, 1, 0, 99)),
		((11107, 1, 3, 0, 99), (1, 1, 3, 0, 99)),
		((11107, 4, 3, 0, 99), (0, 4, 3, 0, 99)),
		((11108, 1, 3, 0, 99), (0, 1, 3, 0, 99)),
		((11108, 4, 4, 0, 99), (1, 4, 4, 0, 99)),
	]

	for t in arrs:
		idVal = 1
		inArr, outArr = t
		exit_code, res = process_array(list(inArr), 1)
		if res != outArr:
			print("invalid result {} -> {} != expected {}".format(inArr, res, outArr))

def test_output():
	arrs = [
		((3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9), 0, 0),
		((3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9), 1, 1),
		((3,3,1105,-1,9,1101,0,0,12,4,12,99,1), 0, 0),
		((3,3,1105,-1,9,1101,0,0,12,4,12,99,1), 1, 1),
	]
	for arr, inID, expOut in arrs:
		exit_code, _ = process_array(list(arr), inID)
		print(exit_code, expOut)

def example(inID):
	arr = (
		3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
		1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
	)
	exit_code, res = process_array(list(arr), inID)
	print(exit_code)

def run(inID):
	with open("data.txt") as inFile:
		arr = tuple(int(v) for v in inFile.read().split(','))
		exit_code, res = process_array(list(arr), inID)
		print(exit_code)

if __name__ == "__main__":
	run(5)