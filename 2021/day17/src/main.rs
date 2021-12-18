use itertools::Itertools;

type Range = (i64, i64);
type TargetArea = (Range, Range);

fn split_range(input: &str) -> (i64, i64) {
    let (_, range) = input.split_once("=").unwrap();
    let v: Vec<_> = range.split("..")
        .map(str::parse)
        .map(Result::unwrap)
        .collect();
    (v[0], v[1])
}

fn parse_input(input: &str) -> TargetArea {
    let (_, tline) = input.split_once(":").unwrap();
    let (xc, yc) = tline.trim().split_once(", ").unwrap();
    (split_range(xc), split_range(yc))
}

fn default_input() -> TargetArea {
    parse_input(include_str!("input.txt"))
}

fn min_x(ta: TargetArea) -> i64 {
    let mut x = 0;
    while (x * (x+1))/2 < ta.0.0 {
        x += 1;
    }
    x
}

fn max_y(ta: TargetArea) -> i64 {
    -ta.1.0 - 1
}

fn hits(init: (i64, i64), target: TargetArea) -> bool {
    let mut c = (0, 0);
    let mut vel = init;
    loop {
        c = (c.0+vel.0, c.1+vel.1);
        if c.0 >= target.0.0 && 
            c.0 <= target.0.1 && 
            c.1 >= target.1.0 && 
            c.1 <= target.1.1 {
            return true;
        }
        if c.1 < target.1.0 {
            return false;
        }

        let mut vel_x = vel.0;
        if vel_x > 0 {
            vel_x -= 1;
        }
        vel = (vel_x, vel.1-1);
    }
}

fn part1(input: TargetArea) -> i64 {
    let y = max_y(input);
    (y * (y + 1)) / 2
}

fn part2(input: TargetArea) -> usize {
    (min_x(input)..=input.0.1)
        .cartesian_product(input.1.0..=max_y(input))
        .filter(|init| hits(*init, input))
        .count()
}

fn main() {
    let i = default_input();
    println!("{}", part1(i));
    println!("{}", part2(i));
}

#[test]
fn test_example() {
    let i = parse_input("target area: x=20..30, y=-10..-5");
    assert_eq!(part1(i), 45);
    assert_eq!(part2(i), 112);
}
