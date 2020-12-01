from itertools import cycle, islice, repeat
from collections import namedtuple


def repeat_phase(in_it, repeat):
	if repeat == 0:
		return
	while True:
		o = next(in_it)
		for _ in range(repeat):
			yield o

def fourier(signal):
	return signal

FFT_MASK = (0, 1, 0, -1)

def get_mask(i, l):
	return tuple(
		islice(repeat_phase(cycle(FFT_MASK), i), 1, l+1)
	)

def gen_masks(i):
	return tuple(get_mask(j+1, i) for j in range(i))

Range = namedtuple("Range", "start stop mult")

def get_ranges(idx, max_i):
	fft_cycle = cycle(FFT_MASK)
	r = []

	i = 0
	while i <= max_i:
		j = i + idx + 1
		n = next(fft_cycle)
		if n != 0:
			start = i - 1
			stop = j - 1
			r.append(Range(start, min(stop, max_i), n))
		i = j
	return r

def gen_ranges(i):
	return tuple(get_ranges(j, i) for j in range(i))

def build_signal(raw_sig):
	return tuple(int(c) for c in raw_sig)

def calc(signal, mask):
	s = sum(a * b for a, b in zip(signal, mask))
	return abs(s) % 10

def calc_ranges(signal, ranges):
	s = sum(
		sum(islice(signal, r.start, r.stop)) * r.mult
		for r in ranges
	)
	return abs(s) % 10


def build_number(num_list):
	return "".join(str(i) for i in num_list)

def get_mask_method(signal):
	masks = gen_masks(len(signal))
	def fft_method(state):
		return tuple(
			calc(state, masks[i]) 
			for i in range(len(state))
		)
	return fft_method

def get_range_method(length):
	ranges = gen_ranges(length)
	def fft_method(state):
		return tuple(
			calc_ranges(state, ranges[i])
			for i in range(len(state))
		)
	return fft_method

class Fourier():
	def __init__(self, signal, fft_method):
		self.state = signal
		self.fft_method = fft_method

	def next_phase(self):
		self.state = self.fft_method(self.state)
		return self.state

	def do(self, n):
		for _ in range(n):
			self.next_phase()
		return self.state


def test_cases():
	cases = (
		("80871224585914546619083218645595", "24176176"),
		("19617804207202209144916044189917", "73745418"),
		("69317163492948606335995924319873", "52432133"),
	)

	for sig, exp in cases:
		signal = build_signal(sig)
		masks = get_range_method(len(signal))
		f = Fourier(signal, 0, masks)
		s = f.do(100)
		n = build_number(s[:8])
		if n != exp:
			print("Wrong answer")

def run():
	with open("data.txt") as f:
		sig = build_signal(f.read())

	masks = get_range_method(len(sig))
	f = Fourier(sig, 0, masks)
	s = f.do(100)
	print(build_number(s[:8]))


def expanded_fft(base_signal):
	outIdx = int("".join(str(i) for i in base_signal[:7]))

	maxIdx = len(base_signal) * 10000
	slicedSig = list(islice(cycle(base_signal), outIdx, maxIdx))

	if outIdx < maxIdx / 2:
		raise Exception("cant shortcut")

	for p in range(100):
		phaseSum = 0
		i = len(slicedSig) - 1
		while i >= 0:
			phaseSum += slicedSig[i]
			slicedSig[i] = abs(phaseSum) % 10
			i -= 1

	return build_number(slicedSig[:8])


def test_cases_part_2():
	cases = (
		("03036732577212944063491565474664", "84462026"),
		("02935109699940807407585447034323", "78725270"),
		("03081770884921959731165446850517", "53553731"),
	)

	for raw_sig, exp in cases:
		base_sig = build_signal(raw_sig)
		r = expanded_fft(base_sig)
		if r != exp:
			print("Wrong answer", r, exp)


def part_2():
	with open("data.txt") as f:
		sig = build_signal(f.read())
	print(expanded_fft(sig))

if __name__ == '__main__':
	# run()
	part_2()

