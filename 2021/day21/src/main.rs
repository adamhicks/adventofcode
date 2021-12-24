use std::cmp;
use std::collections::HashMap;
use itertools::Itertools;

fn parse_input(input: &str) -> Vec<usize> {
    input.lines()
        .map(|l| {l.split_once(": ").unwrap().1})
        .map(str::parse)
        .map(Result::unwrap)
        .collect()
}

fn default_input() -> Vec<usize> {
    parse_input(include_str!("input.txt"))
}


#[derive(Copy, Clone, Debug, Hash, PartialEq, Eq)]
struct Game {
    p: usize,
    p1: (usize, usize),
    p2: (usize, usize),
}

impl Game {
    fn init(pos: &Vec<usize>) -> Self {
        Game{p: 0, p1:(pos[0], 0), p2: (pos[1], 0)}
    }

    fn is_over(&self, max: usize) -> bool {
        self.p1.1 >= max || self.p2.1 >= max
    }

    fn get_mut(&mut self) -> &mut(usize, usize) {
        match self.p {
            0 => &mut self.p1,
            1 => &mut self.p2,
            _ => unreachable!(),
        }
    }

    fn move_player(&mut self, mv: usize) {
        let p = self.get_mut();
        p.0 += mv;
        if p.0 > 10 {
            p.0 %= 10;
        }
        p.1 += p.0;
        self.p = (self.p + 1) % 2;
    }
}

fn play_game(pos: &Vec<usize>) -> usize {
    let mut g = Game::init(pos);
    let mut die = (1..=100).cycle();
    let mut rolls = 0;

    while !g.is_over(1000) {
        let d = &mut die;
        let m = d.take(3).sum::<usize>() % 10;
        rolls += 3;

        g.move_player(m);
    }

    cmp::min(g.p1.1, g.p2.1) * rolls
}

fn probabilities(die: usize, sides: usize) -> Vec<(usize, usize)> {
    (1..=die).map(|_| 1..=sides)
        .multi_cartesian_product()
        .map(|r| r.iter().sum())
        .fold(HashMap::new(), |mut m, s| {
            *m.entry(s).or_default() += 1;
            m
        })
        .into_iter()
        .sorted()
        .collect()
}

fn play_many_games(pos: &Vec<usize>) -> usize {
    let p = probabilities(3, 3);
    let mut games = HashMap::new();
    games.insert(Game::init(pos), 1);

    let mut wins = (0, 0);

    while games.len() > 0 {
        games = games.iter()
            .filter_map(|(g, c)| {
                if g.is_over(21) {
                    if g.p1.1 > g.p2.1 {
                        wins.0 += c;
                    } else {
                        wins.1 += c;
                    }
                    return None;
                }

                Some(p.iter().map(|(mv, count)| {
                    let mut g2 = g.clone();
                    g2.move_player(*mv);
                    (g2, count * *c)
                }))
            })
            .flatten()
            .fold(HashMap::new(), |mut m, (g, c)| {
                *m.entry(g).or_default() += c;
                m
            })
    }
    cmp::max(wins.0, wins.1)
}

fn part1(input: &Vec<usize>) -> usize {
    play_game(input)
}

fn part2(input: &Vec<usize>) -> usize {
    play_many_games(input)
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test_example() {
    let i = parse_input("Player 1 starting position: 4
Player 2 starting position: 8");
    assert_eq!(part1(&i), 739785);
    assert_eq!(part2(&i), 444356092776315);
}
