use std::cmp;
use std::collections::HashSet;
use itertools::Itertools;

type Coord = (i32, i32);
type Enhance = HashSet<usize>;

#[derive(Clone)]
struct LightSet {
    lights: HashSet<Coord>,
    min: Coord,
    max: Coord,
    border: bool,
}

impl LightSet {
    fn from_lights(lights: HashSet<Coord>, border: bool) -> Self {
        let min = min_coord(&lights);
        let max = max_coord(&lights);
        LightSet{lights, min, max, border}
    }

    fn get_at(&self, c: Coord) -> bool {
        if c.0 < self.min.0 || 
            c.0 > self.max.0 || 
            c.1 < self.min.1 || 
            c.1 > self.max.1 {
            return self.border
        }
        self.lights.contains(&c)
    }
}

fn index_lights(input: &str) -> Vec<usize> {
    input.chars()
        .enumerate()
        .filter(|(_, c)| *c == '#')
        .map(|(idx, _)| idx)
        .collect()
}

fn parse_input(input: &str) -> (Enhance, LightSet) {
    let (lights, init) = input.trim().split_once("\n\n").unwrap();
    (
        index_lights(lights).into_iter().collect(), 
        LightSet::from_lights(
            init.lines().enumerate()
                .flat_map(|(y, l)| index_lights(l)
                    .into_iter()
                    .map(move |x| (x as i32, y as i32))
                )
                .collect(),
            false,
        ),
    )
}

fn default_input() -> (Enhance, LightSet) {
    parse_input(include_str!("input.txt"))
}

#[allow(dead_code)]
fn print_lights(l :&LightSet) {
    println!("{:?} -> {:?}", l.min, l.max);

    for s in (l.min.1-2..=l.max.1+2)
        .map(|y| {
            (l.min.0-2..=l.max.0+2)
                .map(|x| {
                    if l.get_at((x, y)) { '█' } else { '░' }
                })
               .collect::<String>()
        })
    {
        println!("{}", s);
    }
}

fn min_coord(l: &HashSet<Coord>) -> Coord {
    l.iter().fold((0,0), |acc, c| (cmp::min(acc.0, c.0), cmp::min(acc.1, c.1)))
}

fn max_coord(l: &HashSet<Coord>) -> Coord {
    l.iter().fold((0,0), |acc, c| (cmp::max(acc.0, c.0), cmp::max(acc.1, c.1)))
}

fn enhance_result(c: Coord, state: &LightSet, enhance: &Enhance) -> bool {
    let idx = (-1..=1).cartesian_product(-1..=1)
        .map(|(y, x)| (c.0 + x, c.1 + y))
        .fold(0, |mut acc, n| {
            acc <<= 1;
            if state.get_at(n) { acc += 1; }
            acc
        });
    enhance.contains(&idx)
}

fn expand(state: &LightSet, enhance: &Enhance) -> LightSet {
    let lights = (state.min.0-1..=state.max.0+1)
        .cartesian_product(state.min.1-1..=state.max.1+1)
        .filter(|c| enhance_result(*c, state, enhance))
        .collect();
    let border = (state.border && enhance.contains(&511)) || 
        (!state.border && enhance.contains(&0));

    LightSet::from_lights(lights, border)
}

fn expand_n(state: &LightSet, enhance: &Enhance, n: usize) -> LightSet {
    if n == 0 {
        return state.clone()
    }
    let mut st = expand(&state, &enhance);
    for _ in 1..n {
        st = expand(&st, &enhance);
    }
    st
}

fn part1(input: &(Enhance, LightSet)) -> usize {
    expand_n(&input.1, &input.0, 2).lights.len()
}

fn part2(input: &(Enhance, LightSet)) -> usize {
    expand_n(&input.1, &input.0, 50).lights.len()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test_example() {
    let i = parse_input("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###");
    assert_eq!(part1(&i), 35);
    // assert_eq!(part2(&i), 3621);
}

#[test]
fn test_example2() {
    let en = HashSet::from([0, 495]);
    let mut st = LightSet::from_lights(HashSet::from([(0,0)]), false);

    print_lights(&st);
    st = expand(&st, &en);
    print_lights(&st);
    st = expand(&st, &en);
    print_lights(&st);

    assert_eq!(1, 0);
}
