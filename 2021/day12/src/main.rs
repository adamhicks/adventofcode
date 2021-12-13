use std::collections::HashMap;

type CaveMap = HashMap<String, Vec<String>>;

fn parse_input(input: &str) -> CaveMap {
    input.trim().lines()
        .map(|l| l.split_once("-").unwrap())
        .fold(CaveMap::new(), |mut cm, (from, to)| {
            cm.entry(from.to_string())
                .or_default()
                .push(to.to_string());
            cm.entry(to.to_string())
                .or_default()
                .push(from.to_string());
            cm
        })
}

fn default_input() -> CaveMap {
    parse_input(include_str!("input.txt"))
}

type CanVisitFn = fn(&String, &HashMap<String, i32>) -> bool;

fn is_upper(s: &String) -> bool {
    return s.chars().all(|c| c.is_ascii_uppercase())
}

fn visit_part1(n: &String, visited: &HashMap<String, i32>) -> bool {
    if is_upper(n) {
        return true;
    }
    if n == "start" {
        return false;
    }
    *visited.get(n).unwrap() == 0
}

fn visit_part2(n: &String, visited: &HashMap<String, i32>) -> bool {
    if is_upper(n) {
        return true;
    }
    if n == "start" {
        return false;
    }

    if *visited.get(n).unwrap() == 0 {
        return true;
    }
    if visited.iter().any(|(v, c)| !is_upper(v) && *c >= 2) {
        return false;
    }
    true
}

struct RouteMap {
    path: Vec<String>,
    visited: HashMap<String, i32>,
}

fn extend_route(r: &RouteMap, n: &String) -> RouteMap {
    let mut p = r.path.clone();
    let mut v = r.visited.clone();
    p.push(n.clone());

    let c = v.get_mut(n).unwrap();
    *c += 1;

    RouteMap{
        path: p,
        visited: v,
    }
}

fn find_routes(m: &CaveMap, can_visit: CanVisitFn) -> Vec<Vec<String>> {
    let mut stk = vec![
        RouteMap{
            path: vec!["start".to_string()],
            visited: m.keys().map(|k| (k.to_string(), 0)).collect(),
        },
    ];

    let mut ret = Vec::new();

    while stk.len() > 0 {
        let route = stk.pop().unwrap();
        let last = route.path.last().unwrap();

        if last == "end" {
            ret.push(route.path);
            continue;
        }

        for to in &m[last] {
            if can_visit(&to, &route.visited) {
                stk.push(extend_route(&route, &to));
            }
        }
    }
    ret
}

fn part1(input: &CaveMap) -> usize {
    find_routes(input, visit_part1).len()
}

fn part2(input: &CaveMap) -> usize {
    find_routes(input, visit_part2).len()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test() {
        let i1 = parse_input("start-A
start-b
A-c
A-b
b-d
A-end
b-end");
    assert_eq!(part1(&i1), 10);
    assert_eq!(part2(&i1), 36);

    let i2 = parse_input("fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW");
    assert_eq!(part1(&i2), 226);
    assert_eq!(part2(&i2), 3509);
}
