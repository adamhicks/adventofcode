use array2d::Array2D;
use itertools::Itertools;
use std::collections::HashSet;

type Map = Array2D<u32>;
type Coord = (usize, usize);

fn parse_input(input: &str) -> Map {
    let m : Vec<Vec<u32>> = input.trim()
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| c.to_digit(10).unwrap())
                .collect()
        })
        .collect();
    Map::from_rows(&m)
}

fn default_input() -> Map {
    parse_input(include_str!("input.txt"))
}

fn neighbours(c: Coord) -> Vec<Coord> {
    let mut n = vec![
        (c.0+1, c.1),
        (c.0, c.1+1),
    ];
    if c.0 > 0 {
        n.push((c.0-1, c.1))
    }
    if c.1 > 0 {
        n.push((c.0, c.1-1))
    }
    n
}

fn part1(m: &Map) -> u32 {
    let mut tot = 0;
    for y in 0..m.num_rows() {
        for x in 0..m.num_columns() {
            let c = (x, y);
            let h = m.get(c.1, c.0).unwrap();
            let lower = neighbours(c).iter()
                .filter_map(|n| m.get(n.1, n.0))
                .filter(|v| *v<=h)
                .count();
            if lower == 0 {
                tot += h+1
            }
        }
    }
    tot
}

fn get_basin(c: Coord, m: &Map) -> HashSet<Coord> {
    let mut stk = vec![c];
    let mut basin = HashSet::new();

    while stk.len() > 0 {
        let c = stk.pop().unwrap();
        match m.get(c.1, c.0) {
            Some(h) => if *h == 9 { continue },
            None => continue,
        }
        basin.insert(c);
        for n in neighbours(c) {
            if basin.contains(&n) {
                continue
            }
            stk.push(n)
        }
    }
    basin
}

fn part2(m: &Map) -> usize {
    let mut unvisit = (0..m.num_columns())
        .cartesian_product(0..m.num_rows())
        .collect::<HashSet<Coord>>();

    let mut basins = Vec::new();

    while unvisit.len() > 0 {
        let c = unvisit.iter().next().unwrap();
        let mut rem = HashSet::new();
        rem.insert(*c);

        let h = m.get(c.1, c.0).unwrap();
        if *h != 9 {
            let b = get_basin(*c, m);
            basins.push(b.len());
            rem = rem.union(&b).map(|c|*c).collect();
        }
        unvisit = &unvisit - &rem;
    }
    basins.sort();
    basins.iter().rev().take(3).product()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("2199943210
3987894921
9856789892
8767896789
9899965678"
    );
    assert_eq!(part1(&i), 15);
    assert_eq!(part2(&i), 1134);
}


