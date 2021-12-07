fn parse_input(input: &str) -> Vec<i64> {
    input.trim().split(",")
        .map(str::parse)
        .map(Result::unwrap)
        .collect()
}

fn default_input() -> Vec<i64> {
    parse_input(include_str!("input.txt"))
}

fn linear_distance(input: &Vec<i64>, val: i64) -> i64 {
    input.iter().fold(0, |sum, v| sum + (v - val).abs())
}

fn sum_distance(input: &Vec<i64>, val: i64) -> i64 {
    input.iter().fold(0, |sum, v| {
        let dis = (v - val).abs();
        let d = (dis * (dis + 1)) / 2;
        sum + d
    })
}

fn min_diff(input: &Vec<i64>, diff: fn(&Vec<i64>, i64) -> i64) -> (i64, i64) {
    let mut i = input.iter().sum::<i64>() as i64 / input.len() as i64;
    let mut d = diff(input, i);

    loop {
        let above = diff(input, i+1);
        if above < d {
            i += 1;
            d = above;
            continue
        }
        let below = diff(input, i-1);
        if below < d {
            i -= 1;
            d = below;
            continue
        }
        return (i, d)
    }
}

fn part1(input: &Vec<i64>) -> i64 {
    let (_, d) = min_diff(input, linear_distance);
    d
}

fn part2(input: &Vec<i64>) -> i64 {
    let (_, d) = min_diff(input, sum_distance);
    d
}


fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("16,1,2,0,4,2,7,1,2,14");
    assert_eq!(part1(&i), 37);
    assert_eq!(part2(&i), 168);
}
