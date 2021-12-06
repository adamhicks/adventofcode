use array2d::Array2D;

type Board = Array2D<i64>;

fn parse_input(input: &str) -> (Vec<i64>, Vec<Board>) {
    let mut i = input.split("\n\n");

    let draws : Vec<i64> = i.next().unwrap()
        .split(",")
        .map(str::parse)
        .map(Result::unwrap)
        .collect();

    let boards : Vec<Board> = i
        .map(|g| {
            let a : Vec<Vec<i64>> = g.lines()
                .map(|line| {
                    line.split_whitespace()
                        .map(str::parse)
                        .map(Result::unwrap)
                        .collect()
                })
                .collect();
            Array2D::from_rows(&a)
        })
        .collect();

    (draws, boards)
}


fn default_input() -> (Vec<i64>, Vec<Board>) {
    parse_input(include_str!("input.txt"))
}

fn find_num(num: i64, board: &Board) -> Option<(usize, usize)> {
    for r in 0..board.num_rows() {
        for c in 0..board.num_columns() {
            if *board.get(r,c).unwrap() == num {
                return Some((r, c));
            }
        }
    }
    return None
}

fn sum_unpicked(picked: Array2D<bool>, board: &Board) -> i64 {
    picked.elements_row_major_iter()
        .zip(board.elements_row_major_iter())
        .fold(0, |sum, i| {
            let (p, v) = i;
            if !*p { sum + v } else { sum }
        })
}

fn when_bingo(draws: &Vec<i64>, board: &Board) -> (usize, i64) {
    let mut picked = Array2D::filled_with(false, board.num_rows(), board.num_columns());

    for (round, num) in draws.iter().enumerate() {
        if let Some((r, c)) = find_num(*num, board) {
            picked.set(r, c, true).unwrap();

            if picked.row_iter(r).all(|b| *b) || 
                picked.column_iter(c).all(|b| *b) {
                return (round, num * sum_unpicked(picked, board))
            }
        }
    }
    (draws.len(), 0)
}

fn part1() {
    let (draws, boards) = default_input();

    let first_bingo = boards.iter()
        .map(|b| when_bingo(&draws, &b))
        .min_by_key(|b| b.0)
        .unwrap();
    
    println!("{}", first_bingo.1);
}

fn part2() {
    let (draws, boards) = default_input();

    let last_bingo = boards.iter()
        .map(|b| when_bingo(&draws, &b))
        .max_by_key(|b| b.0)
        .unwrap();
    
    println!("{}", last_bingo.1);
}

fn main() {
    part1();
    part2();
}

#[test]
fn test_input() -> (Vec<i64>, Vec<Board>) {
    let s = "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7";
    parse_input(s)
}

#[test]
fn test1() {
    let (draws, boards) = test_input();

    let first_bingo = boards.iter()
        .map(|b| when_bingo(&draws, &b))
        .min_by_key(|b| b.0)
        .unwrap();
    
    println!("{}", first_bingo.1);
}
