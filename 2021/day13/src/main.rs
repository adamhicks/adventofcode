use std::collections::HashSet;

type Coord = (usize, usize);

fn parse_input(input: &str) -> (HashSet<Coord>, Vec<(String, usize)>) {
    let (coords, folds) = input.split_once("\n\n").unwrap();
    (
        coords.lines()
            .map(|l| {
                let a: Vec<usize> = l.split(",")
                    .map(str::parse)
                    .map(Result::unwrap)
                    .collect();
                (a[0], a[1])
            }
            )
            .collect(),
        folds.lines()
            .map(|l| {
                let (along, v) = l.split_once("=").unwrap();
                (along.to_string(), v.parse().unwrap())
            })
            .collect(),
    )
}

fn default_input() -> (HashSet<Coord>, Vec<(String, usize)>) {
    parse_input(include_str!("input.txt"))
}

fn fold(v: usize, lim: usize) -> usize {
    if v > lim {
        lim - (v - lim)
    } else {
        v
    }
}

fn fold_coords(coords: &HashSet<Coord>, f: &(String, usize)) -> HashSet<Coord> {
    coords.iter()
            .map(|c| {
                match f.0.as_str() {
                    "fold along x" => (fold(c.0, f.1), c.1),
                    "fold along y" => (c.0, fold(c.1, f.1)),
                    _ => unreachable!(),
                }
            })
            .collect()
}

fn part1(input: &(HashSet<Coord>, Vec<(String, usize)>)) -> usize {
    fold_coords(&input.0, &input.1[0]).len()
}

fn part2(input: &(HashSet<Coord>, Vec<(String, usize)>)) -> String {
    let mut c = input.0.clone();
    for f in &input.1 {
        c = fold_coords(&c, f);
    }
    let max_x = c.iter().map(|c| c.0).max().unwrap();
    let max_y = c.iter().map(|c| c.1).max().unwrap();

    let mut s = String::new();

    for y in 0..=max_y {
        for x in 0..=max_x {
            if c.contains(&(x, y)) {
                s += "#";
            } else {
                s += ".";
            }
        }
        s += "\n";
    }
    s
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5");
    assert_eq!(part1(&i), 17);
    assert_eq!(part2(&i), "#####
#...#
#...#
#...#
#####
");
}
