use std::collections::HashMap;

fn parse_input(input: &str) -> Vec<String> {
    input.lines().map(str::to_string).collect()
}

fn default_input() -> Vec<String> {
    parse_input(include_str!("input.txt"))
}

fn open_lookup() -> HashMap<char, char> {
    let mut r = HashMap::new();
    r.insert('(', ')');
    r.insert('[', ']');
    r.insert('{', '}');
    r.insert('<', '>');
    r
}

fn scores() -> HashMap<char, u64> {
    let mut r = HashMap::new();
    r.insert(')', 3);
    r.insert(']', 57);
    r.insert('}', 1197);
    r.insert('>', 25137);
    r
}

fn scores2() -> HashMap<char, u64> {
    let mut r = HashMap::new();
    r.insert(')', 1);
    r.insert(']', 2);
    r.insert('}', 3);
    r.insert('>', 4);
    r
}

fn find_corrupt(l: &str) -> Option<char> {
    let open = open_lookup();

    let mut stk = Vec::new();

    for c in l.chars() {
        if open.contains_key(&c) {
            stk.push(c);
        } else {
            match stk.pop() {
                Some(op) => {
                    let cl = open.get(&op).unwrap();
                    if *cl != c {
                        return Some(c)
                    }
                }
                None => return Some(c),
            }
        }
    }
    return None
}

fn complete(l : &str) -> Option<String> {
    let open = open_lookup();

    let mut stk = Vec::new();

    for c in l.chars() {
        if open.contains_key(&c) {
            stk.push(c);
        } else {
            match stk.pop() {
                Some(op) => {
                    let cl = open.get(&op).unwrap();
                    if *cl != c {
                        return None
                    }
                }
                None => return None,
            }
        }
    }

    Some(stk.into_iter()
        .rev()
        .filter_map(|c| open.get(&c))
        .collect()
    )
}

fn part1(input: &Vec<String>) -> u64 {
    let sc = scores();

    input.iter()
        .filter_map(|l| find_corrupt(l))
        .filter_map(|c| sc.get(&c))
        .sum()
}

fn part2(input: &Vec<String>) -> u64 {
    let sc2 = scores2(); 
    let mut scores : Vec<_> = input.iter()
        .filter_map(|l| complete(l))
        .map(|s| {
            s.chars()
                .filter_map(|c| sc2.get(&c))
                .fold(0, |acc, s| acc*5 + s )
        })
        .collect();
    scores.sort();
    scores[scores.len()/2]
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
    let i = parse_input("[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]");
    assert_eq!(part1(&i), 26397);
    assert_eq!(part2(&i), 288957);
}


