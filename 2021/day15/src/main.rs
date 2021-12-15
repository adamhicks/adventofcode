use array2d::Array2D;
use itertools::Itertools;
// use std::collections::HashMap;
use std::collections::BinaryHeap;
use std::cmp::Ordering;

type Cave = Array2D<usize>;
type Coord = (usize, usize);

fn parse_input(input: &str) -> Cave {
    let r: Vec<_> = input.lines()
        .map(|l| {
            l.chars()
                .map(|c| c.to_digit(10))
                .map(Option::unwrap)
                .map(|c| c as usize)
                .collect()
        })
        .collect();
    Cave::from_rows(&r)
}

fn default_input() -> Cave {
    parse_input(include_str!("input.txt"))
}

fn neighbours(c: Coord, max: Coord) -> Vec<Coord> {
    let mut r = Vec::new();
    if c.0 > 0 {
        r.push((c.0-1, c.1));
    }
    if c.1 > 0 {
        r.push((c.0, c.1-1));
    }
    if c.0 < max.0 {
        r.push((c.0+1, c.1));
    }
    if c.1 < max.1 {
        r.push((c.0, c.1+1));
    }
    r
}

#[derive(Clone, Eq, PartialEq)]
struct State {
    cost: usize,
    position: Coord,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other.cost.cmp(&self.cost)
            .then_with(|| self.position.cmp(&other.position))
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn min_cost(input: &Cave) -> Option<usize> {
    let end = (input.num_columns()-1, input.num_rows()-1);

    let mut dist = Array2D::filled_with(
        usize::MAX, 
        input.num_rows(), 
        input.num_columns(),
    );
    dist.set(0, 0, 0).unwrap();

    let mut heap = BinaryHeap::new();
    heap.push(State{cost: 0, position: (0,0)});

    while let Some(s) = heap.pop() {
        if s.position == end {
            return Some(s.cost);
        }

        if s.cost > *dist.get(s.position.1, s.position.0).unwrap() {
            continue;
        }

        for n in neighbours(s.position, end){
            let c = input.get(n.1, n.0).unwrap();
            let next = State{cost: s.cost+c, position: n};

            if next.cost < *dist.get(n.1, n.0).unwrap() {
                dist.set(n.1, n.0, next.cost).unwrap();
                heap.push(next);
            }
        }
    }
    None
}

fn expand_map(input: &Cave, size: usize) -> Cave {
    let rows = (0..input.num_rows()*size)
            .map(|y| {
                (0..input.num_columns()*size)
                    .map(|x| {
                        let in_x = x % input.num_columns();
                        let in_y = y % input.num_rows();
                        let scale = x / input.num_columns() + y / input.num_rows();
                        let v = *input.get(in_y, in_x).unwrap() + scale;
                        if v > 9 {
                            v % 9
                        } else {
                            v
                        }
                    })
                    .collect::<Vec<usize>>()
            })
            .collect::<Vec<_>>();
    Cave::from_rows(&rows)
}

fn part1(input: &Cave) -> usize {
    min_cost(&input).unwrap()
}

fn part2(input: &Cave) -> usize {
    let m = expand_map(&input, 5);
    min_cost(&m).unwrap()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581");
    assert_eq!(part1(&i), 40);
    assert_eq!(part2(&i), 315);
}

#[test]
fn test_expansion() {
    let mut m = Cave::filled_with(8, 1, 1);
    m = expand_map(&m, 5);
    assert_eq!(m.as_rows(), vec![
        vec![8,9,1,2,3],
        vec![9,1,2,3,4],
        vec![1,2,3,4,5],
        vec![2,3,4,5,6],
        vec![3,4,5,6,7],
    ]);
}

