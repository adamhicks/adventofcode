def opcode_1(cmd, arr):
	if cmd[0] != 1:
		raise Exception("invalid opcode")

	read1 = arr[cmd[1]]
	read2 = arr[cmd[2]]
	idx = cmd[3]
	arr[idx] = read1 + read2

def opcode_2(cmd, arr):
	if cmd[0] != 2:
		raise Exception("invalid opcode")

	read1 = arr[cmd[1]]
	read2 = arr[cmd[2]]
	idx = cmd[3]
	arr[idx] = read1 * read2

def process_array(arr):
	idx = 0
	while True:
		cmd = arr[idx:idx+4]
		if cmd[0] == 1:
			opcode_1(cmd, arr)
		elif cmd[0] == 2:
			opcode_2(cmd, arr)
		elif cmd[0] == 99:
			return tuple(arr)
		else:
			raise Exception("invalid command")
		idx += 4

def test():
	arrs = [
		((1,0,0,0,99), (2,0,0,0,99)),
		((2,3,0,3,99), (2,3,0,6,99)),
		((2,4,4,5,99,0), (2,4,4,5,99,9801)),
		((1,1,1,4,99,5,6,0,99), (30,1,1,4,2,5,6,0,99)),
	]

	for t in arrs:
		inArr, outArr = t
		res = process_array(list(inArr))
		if res != outArr:
			print("invalid result {} -> {} != expected {}".format(inArr, res, outArr))

def run():
	inArr = [
		1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,
		19,5,23,2,10,23,27,2,27,13,31,1,10,31,35,1,
		35,9,39,2,39,13,43,1,43,5,47,1,47,6,51,2,6,
		51,55,1,5,55,59,2,9,59,63,2,6,63,67,1,13,67,
		71,1,9,71,75,2,13,75,79,1,79,10,83,2,83,9,
		87,1,5,87,91,2,91,6,95,2,13,95,99,1,99,5,
		103,1,103,2,107,1,107,10,0,99,2,0,14,0,
	]
	print(process_array(inArr))

def run_part2():
	target = 19690720
	inArr = (
		1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,
		19,5,23,2,10,23,27,2,27,13,31,1,10,31,35,1,
		35,9,39,2,39,13,43,1,43,5,47,1,47,6,51,2,6,
		51,55,1,5,55,59,2,9,59,63,2,6,63,67,1,13,67,
		71,1,9,71,75,2,13,75,79,1,79,10,83,2,83,9,
		87,1,5,87,91,2,91,6,95,2,13,95,99,1,99,5,
		103,1,103,2,107,1,107,10,0,99,2,0,14,0,
	)
	for x in range(100):
		for y in range(100):
			arr = list(inArr)
			arr[1] = x
			arr[2] = y

			res = process_array(arr)
			if res[0] == target:
				print(x, y)
				return


if __name__ == "__main__":
	run()
	run_part2()
