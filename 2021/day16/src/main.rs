fn parse_input(input: &str) -> Vec<u8> {
    (0..input.trim().len())
        .step_by(2)
        .map(|i| u8::from_str_radix(&input[i..i+2], 16))
        .map(Result::unwrap)
        .collect()
}

fn default_input() -> Vec<u8> {
    parse_input(include_str!("input.txt"))
}

struct BitReader {
    bytes: Vec<u8>,
    byte_off: usize,
    mask: u8,
    read: usize,
}

impl BitReader {
    fn from_bytes(b: &Vec<u8>) -> Self {
        BitReader{
            bytes: b.clone(), 
            byte_off: 0, 
            mask: 1<<7, 
            read: 0,
        }
    }

    fn offset(&self) -> usize {
        self.read
    }

    fn read_bit(&mut self) -> Option<u8> {
        if self.mask == 0 {
            self.byte_off += 1;
            self.mask = 1<<7;
        }
        if self.byte_off >= self.bytes.len() {
            return None
        }
        let one = (self.bytes[self.byte_off] & self.mask) == self.mask;
        self.mask >>= 1;
        self.read += 1;
        if one {
            Some(1)
        } else {
            Some(0)
        }
    }

    fn read_u8(&mut self, count: usize) -> Option<u8> {
        if count > 8 {
            return None;
        }
        match self.read_u32(count) {
            Some(i) => Some(i as u8),
            None => None,
        }
    }

    fn read_u32(&mut self, count: usize) -> Option<u32> {
        if count > 32 {
            return None;
        }
        let mut ret = 0;
        for _ in 0..count {
            ret <<= 1;
            match self.read_bit() {
                Some(s) => ret += s as u32,
                None => return None,
            }
        }
        Some(ret)
    }
}

struct Packet {
    version: u8,
    type_id: u8,

    val: u64,
    sub: Vec<Packet>,
}

impl Packet {
    fn from_bytes(input: &Vec<u8>) -> Packet {
        read_packet(&mut BitReader::from_bytes(input))
    }

    fn literal(version: u8, type_id: u8, val: u64) -> Self {
        Packet{version, type_id, val, sub: Vec::new()}
    }

    fn operator(version: u8, type_id: u8, sub: Vec<Packet>) -> Self {
        Packet{version, type_id, val: 0, sub}
    }

    fn value(&self) -> u64 {
        match self.type_id {
            0 => self.sub.iter().map(Packet::value).sum(),
            1 => self.sub.iter().map(Packet::value).product(),
            2 => self.sub.iter().map(Packet::value).min().unwrap(),
            3 => self.sub.iter().map(Packet::value).max().unwrap(),
            4 => self.val,
            5 => {
                if self.sub[0].value() > self.sub[1].value() { 1 } else { 0 }
            },
            6 => {
                if self.sub[0].value() < self.sub[1].value() { 1 } else { 0 }
            },
            7 => {
                if self.sub[0].value() == self.sub[1].value() { 1 } else { 0 }
            },
            _ => unreachable!(),
        }
    }
}

fn read_packet(r: &mut BitReader) -> Packet {
    let version = r.read_u8(3).unwrap();
    let type_id = r.read_u8(3).unwrap();

    if type_id == 4 {
        return Packet::literal(version, type_id, read_literal(r))
    }

    let mut sub = Vec::new();
    let lt_id = r.read_u8(1).unwrap();

    if lt_id == 0 {
        let s = r.read_u32(15).unwrap() as usize;
        let end = r.offset() + s;
        while r.offset() < end {
            sub.push(read_packet(r));
        }
    } else {
        for _ in 0..r.read_u32(11).unwrap() {
            sub.push(read_packet(r));
        }
    }
    return Packet::operator(version, type_id, sub)
}

fn read_literal(r: &mut BitReader) -> u64 {
    let mut ret: u64 = 0;
    loop {
        ret <<= 4;
        let more = r.read_u8(1).unwrap() == 1;
        ret += r.read_u8(4).unwrap() as u64;
        if !more {
            break
        }
    }
    ret
}

fn part1(input: &Vec<u8>) -> usize {
    let p = Packet::from_bytes(input);
    let mut v_sum = 0;
    let mut stk = vec![p];
    while let Some(pack) = stk.pop() {
        v_sum += pack.version as usize;
        stk.extend(pack.sub);
    }
    v_sum
}

fn part2(input: &Vec<u8>) -> u64 {
    Packet::from_bytes(input).value()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test_example() {
    assert_eq!(part1(&parse_input("D2FE28")), 6);
    assert_eq!(part1(&parse_input("38006F45291200")), 9);
    assert_eq!(part1(&parse_input("8A004A801A8002F478")), 16);
    assert_eq!(part1(&parse_input("620080001611562C8802118E34")), 12);
    assert_eq!(part1(&parse_input("C0015000016115A2E0802F182340")), 23);
    assert_eq!(part1(&parse_input("A0016C880162017C3686B18A3D4780")), 31);
}
