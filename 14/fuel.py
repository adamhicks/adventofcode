from math import ceil
from collections import defaultdict, namedtuple

Reagent = namedtuple("Reagent", "name amount")
Reaction = namedtuple("Reaction", "req prod")

def parse_component(s):
	i, chem = s.split(" ")
	return Reagent(chem, int(i))

def parse_reaction(s):
	in_half, out_half = s.split(" => ")
	in_chems = tuple(parse_component(c.strip()) for c in in_half.split(","))
	out_chem = next(parse_component(c.strip()) for c in out_half.split(","))
	return Reaction(in_chems, out_chem)

def log(s):
	if False:
		print(s)

def get_required(reactions, need_chem, need_amt):
	r_lookup = {}
	for r in reactions:
		r_lookup[r.prod.name] = r

	req_chem = [Reagent(need_chem, need_amt)]
	pool = defaultdict(int)
	ore_amount = 0

	while req_chem:
		chem = req_chem.pop(0)
		log("Need " + str(chem.amount) + " " + chem.name)
		if chem.name == "ORE":
			ore_amount += chem.amount
			continue
		
		pool_amt = pool[chem.name]
		if pool_amt > 0:
			log("Using " + str(pool_amt) + " pooled " + chem.name)
			if pool_amt > chem.amount:
				pool[chem.name] -= chem.amount
				continue
			chem = Reagent(chem.name, chem.amount - pool_amt)
			del pool[chem.name]

		react = r_lookup[chem.name]

		prod_amt = react.prod.amount
		react_mul = ((chem.amount-1) // prod_amt) + 1

		log("Reacted " + chem.name + " " + str(react_mul) + " times")

		surplus = (react_mul * prod_amt) - chem.amount
		if surplus > 0:
			log("Generated " + str(surplus) + " extra " + chem.name)
		pool[chem.name] += surplus

		for in_chem in react.req:
			req_chem.append(Reagent(in_chem.name, in_chem.amount * react_mul))

	return ore_amount


def test_example():
	cases = (
		("10 ORE => 10 A\n1 ORE => 1 B\n7 A, 1 B => 1 C\n7 A, 1 C => 1 D\n7 A, 1 D => 1 E\n7 A, 1 E => 1 FUEL", 31),
		("157 ORE => 5 NZVS\n165 ORE => 6 DCFZ\n44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL\n12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ\n179 ORE => 7 PSHF\n177 ORE => 5 HKGWZ\n7 DCFZ, 7 PSHF => 2 XJWVT\n165 ORE => 2 GPVTF\n3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT", 13312),
		("171 ORE => 8 CNZTR\n7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL\n114 ORE => 4 BHXH\n14 VRPVC => 6 BMBT\n6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL\n6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT\n15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW\n13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW\n5 BMBT => 4 WPTQ\n189 ORE => 9 KTJDG\n1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP\n12 VRPVC, 27 CNZTR => 2 XDBXC\n15 KTJDG, 12 BHXH => 5 XCVML\n3 BHXH, 2 VRPVC => 7 MZWV\n121 ORE => 7 VRPVC\n7 XCVML => 6 RJRHP\n5 BHXH, 4 VRPVC => 5 LTCX", 2210736)
	)

	for in_str, expOre in cases:
		reactions = [parse_reaction(l) for l in in_str.splitlines()]
		ore = get_required(reactions, "FUEL", 1)

		if ore != expOre:
			print("Wrong ore amount!", ore, expOre)

def run():
	with open("reactions.txt") as r:
		reactions = [parse_reaction(l) for l in r.readlines()]
		ore = get_required(reactions, "FUEL", 1)
		print(ore)

def run_part_2():
	max_ore = 1000000000000
	max_fuel = 1000000000000
	min_fuel = 0

	with open("reactions.txt") as r:
		reactions = [parse_reaction(l) for l in r.readlines()]

	while max_fuel != min_fuel:
		mid = (max_fuel + min_fuel) // 2
		req_ore = get_required(reactions, "FUEL", mid)
		print(str(min_fuel) + "-" + str(max_fuel) + " " + str(mid) + " FUEL needs " + str(req_ore) + " ORE")

		if req_ore > max_ore:
			max_fuel = mid
		else:
			min_fuel = mid
		
		if abs(max_fuel - min_fuel) == 1:
			max_fuel = min_fuel



if __name__ == "__main__":
	run()
	run_part_2()
