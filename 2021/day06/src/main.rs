fn parse_input(input: &str) -> Vec<usize> {
    input.trim().split(",")
        .map(str::parse)
        .map(Result::unwrap)
        .collect()
}

fn default_input() -> Vec<usize> {
    parse_input(include_str!("input.txt"))
}

fn sim_fish(start: &Vec<usize>, days: usize) -> i64 {
    let mut week = vec![0i64;7];
    for v in start {
        week[*v] += 1
    }

    let mut incub = vec![0i64;2];

    for i in 0..days {
        let idx = i % 7;
        let ready = incub.pop().unwrap();
        incub.insert(0, week[idx]);
        week[idx] += ready;
    }
    week.iter().chain(incub.iter()).sum()
}

fn part1() {
    let i = default_input();
    println!("{:?}", sim_fish(&i, 80));
}

fn part2() {
    let i = default_input();
    println!("{:?}", sim_fish(&i, 256));
}


fn main() {
    part1();
    part2();
}

#[test]
fn test() {
    let i = parse_input("3,4,3,1,2");
    assert_eq!(sim_fish(&i, 18), 26);
    assert_eq!(sim_fish(&i, 256), 26984457539);
}
