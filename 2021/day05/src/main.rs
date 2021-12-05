use array2d::Array2D;

type Coord = (usize, usize);
type Line = (Coord, Coord);

fn parse_coord(input: &str) -> Coord {
    let v : Vec<usize> = input.split(",")
        .map(str::parse)
        .map(Result::unwrap)
        .collect();
    (v[0], v[1])
}

fn default_input() -> Vec<Line> {
    parse_input(include_str!("input.txt"))
}

fn parse_input(input: &str) -> Vec<Line> {    
    input.lines()
        .map(|line| {
            let (a, b) = line.split_once(" -> ").unwrap();
            (parse_coord(a), parse_coord(b))
        })
        .collect()
}


fn between(a: usize, b: usize) -> Vec<usize> {
    if a == b {
        return vec![a]
    } else if a <= b {
        (a..=b).collect()
    } else {
        (b..=a).rev().collect()
    }
}

fn line_coords(l: Line, diag: bool) -> Vec<Coord> {
    let (x1, y1) = l.0;
    let (x2, y2) = l.1;

    let x_v = between(x1, x2);
    let y_v = between(y1, y2);

    if x_v.len() == 1 {
        return y_v.iter().map(|&y| (x_v[0], y)).collect()
    } else if y_v.len() == 1 {
        return x_v.iter().map(|&x| (x, y_v[0])).collect()
    }
    if !diag {
        return Vec::new()
    }
    return x_v.iter()
        .zip(y_v.iter())
        .map(|v| (*v.0, *v.1))
        .collect()
}

fn map_lines(size: usize, lines: Vec<Line>, diag: bool) -> Array2D<i64> {
    let mut grid = Array2D::filled_with(0, size, size);
    for l in lines {
        for c in line_coords(l, diag) {
            grid.get_mut(c.1, c.0).map(|v| *v += 1 );
        }
    }
    grid
}

fn count_two_plus(grid: Array2D<i64>) -> usize {
    grid.elements_row_major_iter()
        .filter(|&v| *v >= 2)
        .count()
}

fn part1() {
    let i = default_input();
    let sum = count_two_plus(map_lines(1000, i, false));
    println!("{}", sum);
}

fn part2() {
    let i = default_input();
    let sum = count_two_plus(map_lines(1000, i, true));
    println!("{}", sum);
}

fn main() {
    part1();
    part2();
}


#[test]
fn test_input() -> Vec<Line> {
    let s = "0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2";
    return parse_input(s)
}

#[test]
fn test1() {
    let i = test_input();
    let sum = count_two_plus(map_lines(10, i, false));
    assert_eq!(sum, 5);
}

#[test]
fn test2() {
    let i = test_input();
    let sum = count_two_plus(map_lines(10, i, true));
    assert_eq!(sum, 12);
}