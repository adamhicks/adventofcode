fn parse_input(input: &str) -> Vec<(String, i64)> {
    input
        .lines()
        .map(|line| {
            let (cmd, value) = line.split_once(char::is_whitespace).unwrap();
            (String::from(cmd), value.parse().unwrap())
        })
        .collect()
}


fn default_input() -> Vec<(String, i64)> {
    parse_input(include_str!("input.txt"))
}

fn part1() {
    let mut h = 0;
    let mut d = 0;
    for (cmd, v) in default_input() {
        match cmd.as_str() {
            "up" => d -= v,
            "down" => d += v,
            "forward" => h += v,
            _ => panic!("unknown command '{}'", cmd),
        }
    }
    println!("{}", d * h);
}

fn part2() {
    let mut aim = 0;
    let mut h = 0;
    let mut d = 0;
    for (cmd, v) in default_input() {
        match cmd.as_str() {
            "up" => aim -= v,
            "down" => aim += v,
            "forward" => {
                h += v; 
                d += aim * v;
            }
            _ => panic!("unknown command '{}'", cmd),
        }
    }
    println!("{}", d * h);
}

fn main() {
    part1();
    part2();
}
