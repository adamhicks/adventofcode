def mod_fuel(mass):
	return int(float(mass) / 3) - 2

def sum_fuel(fuel):
	total = fuel
	while True:
		fuel_fuel = mod_fuel(fuel)
		if fuel_fuel <= 0:
			return total
		total += fuel_fuel
		fuel = fuel_fuel

def test():
	pairs = [
		(12, 2),
		(14, 2),
		(1969, 654),
		(100756, 33583),
	]
	for p in pairs:
		m, exp = p
		res = mod_fuel(m)
		if res != exp:
			print("error {} != {}".format(p, res))

def test_part2():
	pairs = [
		(14, 2),
		(1969, 966),
		(100756, 50346),
	]
	for p in pairs:
		m, exp = p
		res = sum_fuel(mod_fuel(m))
		if res != exp:
			print("error {} != {}".format(p, res))

def run():
	with open("data.txt") as mass_file:
		total = sum(
			sum_fuel(mod_fuel(int(l.strip())))
			for l in mass_file.readlines()
		)
		print(total)


if __name__ == "__main__":
	run()
