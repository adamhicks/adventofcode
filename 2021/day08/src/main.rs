use std::collections::HashMap;
use std::collections::HashSet;

type CharSet = HashSet<char>;
type SegMap = HashMap<usize, CharSet>;

fn parse_input(input: &str) -> Vec<(Vec<CharSet>, Vec<CharSet>)> {
    input.trim().lines().map(| line | {
        let (signal, output) = line.split_once(" | ").unwrap();
        (
            signal.split(" ").map(|w| w.chars().collect::<CharSet>()).collect(), 
            output.split(" ").map(|w| w.chars().collect::<CharSet>()).collect(),
        )
    }).collect()
}

fn default_input() -> Vec<(Vec<CharSet>, Vec<CharSet>)> {
    parse_input(include_str!("input.txt"))
}

fn got_one(v: Vec<CharSet>) -> CharSet {
    if v.len() != 1 {
        panic!("found non-single signal {:?}", v);
    }
    return v[0].clone()
} 

fn find_len(s: &Vec<CharSet>, l: usize) -> CharSet {
    got_one(s.iter()
        .filter(|s| s.len() == l)
        .map(CharSet::clone)
        .collect()
    )
}

fn find_nine(sigs: &Vec<CharSet>, m: &SegMap) -> CharSet {
    let four = m.get(&4).unwrap();
    got_one(sigs.iter()
        .filter(|s| s.len() == 6 && s.is_superset(four))
        .map(CharSet::clone)
        .collect())
}

fn find_zero(sigs: &Vec<CharSet>, m: &SegMap) -> CharSet {
    let nine = m.get(&9).unwrap();
    let one = m.get(&1).unwrap();
    got_one(sigs.iter()
        .filter(|s| s.len() == 6 && *s != nine && one.is_subset(s))
        .map(CharSet::clone)
        .collect())
}

fn find_three(sigs: &Vec<CharSet>, m: &SegMap) -> CharSet {
    let one = m.get(&1).unwrap();
    got_one(sigs.iter()
        .filter(|s| s.len() == 5 && one.is_subset(s))
        .map(CharSet::clone)
        .collect())
}

fn find_five(sigs: &Vec<CharSet>, m: &SegMap) -> CharSet {
    let nine = m.get(&9).unwrap();
    let three = m.get(&3).unwrap();
    got_one(sigs.iter()
        .filter(|s| s.len() == 5 && *s != three && s.is_subset(nine))
        .map(CharSet::clone)
        .collect())
}

fn find_six(sigs: &Vec<CharSet>, m: &SegMap) -> CharSet {
    let one = m.get(&1).unwrap();
    got_one(sigs.iter()
        .filter(|s| s.len() == 6 && (*s - one).len() == 5)
        .map(CharSet::clone)
        .collect())
}

fn find_two(sigs: &Vec<CharSet>, m: &SegMap) -> CharSet {
    let four = m.get(&4).unwrap();
    got_one(sigs.iter()
        .filter(|s| s.len() == 5 && s.union(four).count() == 7)
        .map(CharSet::clone)
        .collect())
}

fn decode_map(signal: &Vec<CharSet>) -> SegMap {
    let mut m = SegMap::new();
    m.insert(1, find_len(signal, 2));
    m.insert(7, find_len(signal, 3));
    m.insert(4, find_len(signal, 4));
    m.insert(8, find_len(signal, 7));
    m.insert(9, find_nine(signal, &m));
    m.insert(0, find_zero(signal, &m));
    m.insert(3, find_three(signal, &m));
    m.insert(5, find_five(signal, &m));
    m.insert(6, find_six(signal, &m));
    m.insert(2, find_two(signal, &m));
    m
}

fn decode_output(signal: &Vec<CharSet>, output: &Vec<CharSet>) -> usize {
    let m = decode_map(signal);
    let mut out = 0;
    for o in output {
        out *= 10;
        for (num, val) in &m {
            if val == o {
                out += num;
                break
            }
        }
    }
    out
}

fn part1(signals: &Vec<(Vec<CharSet>, Vec<CharSet>)>) -> usize {
    signals.iter().map(| (_, o) | {
        o.iter()
            .filter(|w| {
                w.len() == 2 || w.len() == 3 || w.len() == 4 || w.len() == 7
            })
            .count()
    }).sum()
}

fn part2(signals: &Vec<(Vec<CharSet>, Vec<CharSet>)>) -> usize {
    signals.iter().map(|(s,o)| decode_output(s, o)).sum()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input(
        "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
");
    assert_eq!(part1(&i), 26);
    assert_eq!(part2(&i), 61229);
}
