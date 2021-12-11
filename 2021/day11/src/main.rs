use array2d::Array2D;
use itertools::Itertools;
use std::collections::HashSet;

type Cave = Array2D<u32>;
type Coord = (i32, i32);

fn parse_input(input: &str) -> Cave {
    let r: Vec<_> = input.lines()
        .map(|l| {
            l.chars()
                .map(|c| c.to_digit(10))
                .map(Option::unwrap)
                .collect()
        })
        .collect();
    Cave::from_rows(&r)
}

fn default_input() -> Cave {
    parse_input(include_str!("input.txt"))
}

fn neighbours(c: Coord, max: Coord) -> Vec<Coord> {
    (c.0-1..=c.0+1).cartesian_product(c.1-1..=c.1+1)
        .filter(|n| {
            *n != c && 
            n.0 >= 0 && 
            n.1 >= 0 && 
            n.0 < max.0 && 
            n.1 < max.1
        })
        .collect()
}

fn next_state(c: &Cave) -> (u32, Cave) {
    let max = (c.num_columns() as i32, c.num_rows() as i32);

    let mut next = Cave::from_iter_row_major(
        c.elements_row_major_iter()
            .map(|v| v+1),
            c.num_rows(), c.num_columns(),
    );

    let mut to_fire: HashSet<Coord> = (0..next.num_columns())
        .cartesian_product(0..next.num_rows())
        .filter(|n| *next.get(n.1 as usize, n.0 as usize).unwrap() > 9)
        .map(|(x, y)| (x as i32, y as i32))
        .collect();
    let mut fired = HashSet::new();

    while to_fire.len() > 0 {
        let n = *to_fire.iter().next().unwrap();
        to_fire.remove(&n);

        fired.insert(n);
        for m in neighbours(n, max) {
            if fired.contains(&m) {
                continue
            }
            let v = next.get_mut(m.1 as usize, m.0 as usize).unwrap();
            *v += 1;
            if *v > 9 {
                to_fire.insert(m);
            }
        }
    }

    for (x, y) in &fired {
        next.set(*y as usize, *x as usize, 0).unwrap();
    }
    (fired.len() as u32, next)
}

fn part1(input: &Cave) -> u32 {
    let mut c = input.clone();
    let mut sum = 0;
    for _ in 0..100 {
        let (fired, next) = next_state(&c);
        c = next;
        sum += fired;
    }
    sum
}

fn part2(input: &Cave) -> u32 {
    let mut c = input.clone();
    let mut i = 1;
    let exp = (input.num_rows() * input.num_columns()) as u32;
    loop {
        let (fired, next) = next_state(&c);
        if fired == exp {
            return i
        }
        c = next;
        i += 1;
    }
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526");
    assert_eq!(part1(&i), 1656);
    assert_eq!(part2(&i), 195);
}


