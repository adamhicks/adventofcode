use itertools::Itertools;
use itertools::iproduct;
use std::collections::HashSet;
use std::collections::HashMap;
use std::collections::VecDeque;

type Point = (i32, i32, i32);
type PointCloud = HashSet<Point>;
// Column major transform matrix
type Transform = (Point, Point, Point);

fn parse_input(input: &str) -> Vec<PointCloud> {
    input.split("\n\n")
        .map(|s| {
            let mut i = s.lines();
            i.next();
            i.map(|l| {
                let p: Vec<i32> = l.split(",")
                    .map(str::parse)
                    .map(Result::unwrap)
                    .collect();
                (p[0], p[1], p[2])
            }).collect()
        })
        .collect()
}

fn default_input() -> Vec<PointCloud> {
    parse_input(include_str!("input.txt"))
}

fn mul(p1: Point, p2: Point) -> i32 {
    p1.0 * p2.0 + p1.1 * p2.1 + p1.2 * p2.2
}

fn point_to_matrix(p: Point) -> Transform {
    ((p.0, 0, 0), (0, p.1, 0), (0, 0, p.2))
}

fn transform(p: Point, t: Transform) -> Point {
    let ti = invert(t);
    (mul(p, ti.0), mul(p, ti.1), mul(p, ti.2))
}

fn dot(t1: Transform, t2: Transform) -> Transform {
    let t2i = invert(t2);
    (
        (mul(t1.0, t2i.0), mul(t1.0, t2i.1), mul(t1.0, t2i.2)),
        (mul(t1.1, t2i.0), mul(t1.1, t2i.1), mul(t1.1, t2i.2)),
        (mul(t1.2, t2i.0), mul(t1.2, t2i.1), mul(t1.2, t2i.2)),
    )
}

// Swap rows for columns
fn invert(t: Transform) -> Transform {
    ((t.0.0, t.1.0, t.2.0),
     (t.0.1, t.1.1, t.2.1),
     (t.0.2, t.1.2, t.2.2))
}

fn all_transforms(t: Transform) -> Vec<Transform> {
    let mut id = point_to_matrix((1,1,1));
    let mut ret = vec![id];
    for _ in 0..3 {
        id = dot(id, t);
        ret.push(id);
    }
    ret
}

fn unique_transforms() -> Vec<Transform> {
    let yaw = ((0, 0, -1), (0, 1, 0), (1, 0, 0));
    let pitch = ((1, 0, 0), (0, 0, -1), (0, 1, 0));
    let roll = ((0, -1, 0), (1, 0, 0), (0, 0, 1));

    iproduct!(all_transforms(yaw), all_transforms(pitch), all_transforms(roll))
        .map(|(t1, t2, t3)| dot(dot(t1, t2), t3))
        .collect::<HashSet<_>>()
        .iter().copied().collect()
}

fn transform_cloud(pc: &PointCloud, t: Transform) -> PointCloud {
    pc.iter().map(|p| transform(*p, t)).collect()
}

fn move_cloud(pc: &PointCloud, mv: Point) -> PointCloud {
    pc.iter()
        .map(|p| (p.0 + mv.0, p.1 + mv.1, p.2 + mv.2))
        .collect()
}

fn abs_dist(p1: Point, p2: Point) -> i32 {
    (p2.0 - p1.0).abs() + (p2.1 - p1.1).abs() + (p2.2 - p1.2).abs()
}

fn dist(p1: Point, p2: Point) -> Point {
    (p2.0 - p1.0, p2.1 - p1.1, p2.2 - p1.2)
}

fn orient(src: &PointCloud, refc: &PointCloud) -> Result<(Point, PointCloud), ()> {
    for t in unique_transforms() {
        let ref_pc = transform_cloud(refc, t);
        let distances: HashMap<Point, usize> = src.iter()
            .cartesian_product(ref_pc.iter())
            .fold(HashMap::new(), |mut m, (p1, p2)| {
                let d = dist(*p2, *p1);
                *m.entry(d).or_default() += 1;
                m
            });

        let common_dist = distances.iter()
            .max_by_key(|(_, v)| *v).unwrap();

        if *common_dist.1 >= 12 {
            let scan = *common_dist.0;
            return Ok((scan, move_cloud(&ref_pc, scan)));
        }
    }
    Err(())
}

fn merge(clouds: &[PointCloud]) -> (Vec<Point>, PointCloud) {
    let mut pc = clouds[0].clone();
    let mut left: VecDeque<_> = clouds[1..].iter().collect();
    let mut scanners = vec![(0,0,0)];

    while left.len() > 0 {
        let nxt = left.pop_front().unwrap();
        match orient(&pc, nxt) {
            Ok((scan, cloud)) => {
                pc = pc.union(&cloud).copied().collect();
                scanners.push(scan);
            },
            Err(_) => left.push_back(nxt),
        }
    }
    (scanners, pc)
}

fn part1(input: &[PointCloud]) -> usize {
    merge(input).1.len()
}

fn part2(input: &[PointCloud]) -> i32 {
    merge(input).0.iter().permutations(2)
        .map(|v| abs_dist(*v[0], *v[1]))
        .max().unwrap()
}

fn main() {
    let i = default_input();
    println!("{}", part1(&i));
    println!("{}", part2(&i));
}

#[test]
fn test_example() {
    let i = parse_input(include_str!("test_input.txt"));
    assert_eq!(part1(&i), 79);
    assert_eq!(part2(&i), 3621);
}
