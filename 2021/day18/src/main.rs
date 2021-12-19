use std::str::Chars;
use itertools::Itertools;

#[derive(Clone, Debug)]
enum Node {
    Lit(u32),
    Br(Box<Node>, Box<Node>),
}

impl Node {
    fn parse(input: &str) -> Self {
        Self::parse_node(&mut input.chars())
    }

    fn parse_node(it: &mut Chars<'_>) -> Self {
        match it.next().unwrap() {
            '[' => {},
            c => {
                return Node::Lit(c.to_digit(10).unwrap());
            }
        }
        let left = Self::parse_node(it);
        match it.next().unwrap() {
            ',' => {},
            _ => unreachable!(),
        }
        let right = Self::parse_node(it);
        match it.next().unwrap() {
            ']' => {},
            _ => unreachable!(),
        }
        return Node::Br(Box::new(left), Box::new(right))
    }

    fn pair_mut(&mut self) -> Option<(&mut Node, &mut Node)> {
        match self {
            Node::Br(left, right) => Some((left, right)),
            Node::Lit(_) => None,
        }
    }

    fn to_string(&self) -> String {
        match self {
            Node::Lit(v) => v.to_string(),
            Node::Br(left, right) => format!(
                "[{},{}]", 
                left.to_string(), 
                right.to_string(),
            ),
        }
    }

    fn mag(&self) -> u32 {
        match self {
            Node::Lit(v) => *v,
            Node::Br(left, right) => {
                left.mag() * 3 + right.mag() * 2
            }
        }
    }

}

fn explode(n: &mut Node) -> bool {
    explode_node(n, 0) != None
}

fn explode_node(n: &mut Node, depth: usize) -> Option<(u32, u32)> {
    if let Some((left, right)) = n.pair_mut() {
        match (left, right) {
            (Node::Lit(l), Node::Lit(r)) => {
                if depth < 4 {
                    return None
                }
                let ret = Some((*l, *r));
                *n = Node::Lit(0);
                return ret;
            }
            (left, right) => {
                if let Some((l, r)) = explode_node(left, depth+1) {
                    fill_left(right, r);
                    return Some((l, 0));
                }
                if let Some((l, r)) = explode_node(right, depth+1) {
                    fill_right(left, l);
                    return Some((0, r));
                }
            }
        }
    }
    None
}

fn split(n: &mut Node) -> bool {
    match n {
        Node::Lit(v) => {
            if *v >= 10 {
                let l = Node::Lit(*v/2);
                let r = Node::Lit(*v/2 + *v%2);
                *n = Node::Br(Box::new(l), Box::new(r));
                return true;
            }
            false
        },
        Node::Br(left, right) => {
            if split(left) {
                return true;
            }
            if split(right) {
                return true;
            }
            false
        },
    }
}

fn fill_left(n: &mut Node, val: u32) {
    match n {
        Node::Lit(v) => *v += val,
        Node::Br(left, _) => fill_left(left, val),
    }
}

fn fill_right(n: &mut Node, val: u32) {
    match n {
        Node::Lit(v) => *v += val,
        Node::Br(_, right) => fill_right(right, val),
    }
}

fn parse_input(input: &str) -> Vec<Node> {
    input.lines()
        .map(Node::parse)
        .collect()
}

fn default_input() -> Vec<Node> {
    parse_input(include_str!("input.txt"))
}

fn add(a: Node, b: Node) -> Node {
    let mut t = Node::Br(Box::new(a), Box::new(b));
    loop {
        if explode(&mut t) {
            continue
        }
        if split(&mut t) {
            continue
        }
        break
    }
    t
}

fn part1(input: Vec<Node>) -> u32 {
    input.into_iter().reduce(add).unwrap().mag()
}

fn part2(input: Vec<Node>) -> u32 {
    input.into_iter()
        .permutations(2)
        .map(|v| {
            v.into_iter().reduce(add).unwrap().mag()
        })
        .max()
        .unwrap()
}

fn main() {
    let i = default_input();
    println!("{}", part1(i.clone()));
    println!("{}", part2(i));
}

#[test]
fn test_add() {
    let i = parse_input("[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]");
    let n = i.into_iter().reduce(add).unwrap();
    assert_eq!(n.to_string(), "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]");
}

#[test]
fn test_example() {
    let i = parse_input("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]");
    assert_eq!(part1(i), 3488);
}

#[test]
fn test_example2() {
    let i = parse_input("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]");
    assert_eq!(part1(i.clone()), 4140);
    assert_eq!(part2(i), 3993);
}