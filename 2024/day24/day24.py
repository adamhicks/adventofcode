import random
from collections import namedtuple
from itertools import chain

test_input = """x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
"""

Gate = namedtuple("Gate", ["in_1", "op", "in_2", "equals", "out"])

def parse_input(s):
    first, second = s.split("\n\n")
    vals = {g: v == "1" for g, v in (l.split(": ") for l in first.splitlines())}
    gates = [Gate(*l.split(" ")) for l in second.splitlines()]
    return vals, gates

def build_decimal(vals):
    i = 0
    for v in vals:
        i <<= 1
        if v:
            i += 1
    return i

def convert_vals(vals):
    ins = reversed(sorted((n, v) for n, v in vals.items()))
    return build_decimal(v for _, v in ins)

def run_gates(vals, gates):
    while True:
        ng = next((
            g for g in gates if
            g.in_1 in vals and
            g.in_2 in vals and
            g.out not in vals
        ), None)
        if ng is None:
            break
        v_and = vals[ng.in_1] and vals[ng.in_2]
        v_or = vals[ng.in_1] or vals[ng.in_2]
        match ng.op:
            case "AND":
                vals[ng.out] = v_and
            case "OR":
                vals[ng.out] = v_or
            case "XOR":
                vals[ng.out] = v_or and not v_and
            case _:
                raise ValueError("invalid op", ng.op)

def run_part1(i):
    vals, gates = i

    run_gates(vals, gates)

    zs = reversed(sorted((k, v) for k, v in vals.items() if k.startswith("z")))
    d = build_decimal(v for k, v in zs)
    print(d)

def get_bit(z, b):
    return (z & (1 << b)) != 0

def vbit(name, bit):
    return f"{name}{bit:0>2}"

def find_faulty_bit(gates):
    max_bits = sum(1 for g in gates if g.out.startswith("z"))

    faults = []

    for _ in range(20):
        x_vals = {vbit("x", i): random.randint(0, 1) == 0 for i in range(max_bits)}
        y_vals = {vbit("y", i): random.randint(0, 1) == 0 for i in range(max_bits)}
        x = convert_vals(x_vals)
        y = convert_vals(y_vals)
        exp_z = x + y
        vals = x_vals
        vals.update(y_vals)

        run_gates(vals, gates)
        for i in range(max_bits):
            z_var = vbit("z", i)
            if vals.get(z_var, None) != get_bit(exp_z, i):
                faults.append(i)
                break

    return min(faults)

def fix_faulty_bit(b, gates):
    all_gates = set(gates)
    good = set(involved(b-1, gates))
    sus = set(involved(b, gates)) - good
    rest = set(gates) - good

    for g in sus:
        for h in rest:
            if g == h:
                continue
            new_gates = all_gates.copy()
            new_gates.remove(g)
            new_gates.remove(h)
            new_gates.add(g._replace(out=h.out))
            new_gates.add(h._replace(out=g.out))

            nb = find_faulty_bit(list(new_gates))
            if nb > b:
                return g, h
    return None

def run_part2(i):
    _, gates = i

    faults = []
    while True:
        b = find_faulty_bit(gates)
        if b is None:
            break
        g, h = fix_faulty_bit(b, gates)
        faults += g, h
        if len(faults) == 8:
            break
        gates.remove(g)
        gates.remove(h)
        gates.append(g._replace(out=h.out))
        gates.append(h._replace(out=g.out))

    print(",".join(sorted(g.out for g in faults)))

def involved(bit, gates):
    nodes = [vbit("z", i) for i in range(bit+1)]
    while nodes:
        n = nodes.pop()
        gs = [g for g in gates if g.out == n]
        yield from gs
        nodes.extend(chain.from_iterable((g.in_1, g.in_2) for g in gs))

### Generated by start script

def test_part1():
    run_part1(parse_input(test_input))

def part1():
    run_part1(parse_input(open("input.txt").read()))

def test_part2():
    # test input doesn't solve in the same way as real input
    pass

def part2():
    run_part2(parse_input(open("input.txt").read()))

def main():
    print("=== running part 1 test ===")
    test_part1()
    print("=== running part 1 ===")
    part1()
    print("=== running part 2 test ===")
    test_part2()
    print("=== running part 2 ===")
    part2()
    print("=== ===")

if __name__ == "__main__":
    main()
