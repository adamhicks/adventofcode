use std::collections::HashMap;

type Pair = (char, char);
type Grammar = HashMap<Pair, Vec<Pair>>;

fn parse_input(input: &str) -> (String, Grammar) {
    let (init, rules) = input.split_once("\n\n").unwrap();
    (
        init.to_string(),
        rules.lines()
            .map(|l| {
                let (from, ins) = l.split_once(" -> ").unwrap();
                let mut c = from.chars();
                let a = c.next().unwrap();
                let b = c.next().unwrap();
                let i = ins.chars().next().unwrap();
                ((a, b), vec![(a, i), (i, b)])
            })
            .collect()
    )
}

fn default_input() -> (String, Grammar) {
    parse_input(include_str!("input.txt"))
}

fn next(st: &HashMap<Pair, i64>, g: &Grammar) -> HashMap<Pair, i64> {
    st.iter()
        .fold(HashMap::new(), |mut m, (from, c)| {
            for to in g[from].iter() {
                let e = m.entry(*to).or_default();
                *e += c;
            }
            m
        })
}

fn count_chars(st: &HashMap<Pair, i64>, init: &Vec<char>) -> HashMap<char, i64> {
    st.iter()
        .fold(HashMap::new(), |mut m: HashMap<char, i64>, (p, c) | {
            *m.entry(p.0).or_default() += c;
            *m.entry(p.1).or_default() += c;
            m
        })
        .iter()
        .map(|(c, count)| {
            let mut v = *count;
            if c == init.first().unwrap() {
                v += 1;
            }
            if c == init.last().unwrap() {
                v += 1;
            }
            (*c, v/2)
        })
        .collect()
}

fn run_expansion(init: &str, g: &Grammar, count: usize) -> i64 {
    let chs: Vec<_> = init.chars().collect();

    let mut state: HashMap<Pair, i64> = HashMap::new();

    for c in chs.windows(2) {
        let p = (c[0], c[1]);
        let e = state.entry(p).or_default();
        *e += 1;
    }

    for _ in 0..count {
        state = next(&state, g);
    }
    let counts = count_chars(&state, &chs);
    counts.values().max().unwrap() - counts.values().min().unwrap()
}

fn part1(input: &(String, Grammar)) -> i64 {
    let (init, g) = input;
    run_expansion(init, &g, 10)
}

fn part2(input: &(String, Grammar)) -> i64 {
    let (init, g) = input;
    run_expansion(init, &g, 40)
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C");
    assert_eq!(part1(&i), 1588);
    assert_eq!(part2(&i), 2188189693529);
}
