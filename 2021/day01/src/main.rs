use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;


fn part1() {
    let mut inc = 0;
    let mut last = 0;

    if let Ok(lines) = read_lines("input.txt") {
        for line in lines {
            if let Ok(l) = line {
                let cur: i32 = l.parse().unwrap();

                if last != 0 {
                    if cur > last {
                        inc += 1;
                    }
                }
                last = cur
            }
        }
    }
    println!("{}", inc);
}

fn part2() {
    let mut v = Vec::new();

    if let Ok(lines) = read_lines("input.txt") {
        for line in lines {
            if let Ok(l) = line {
                let cur: i32 = l.parse().unwrap();
                v.push(cur)
            }
        }
    }

    let mut inc = 0;

    for n in 3..=v.len()-1 {
        let prev = v[n-3];
        if v[n] > prev {
            inc += 1;
        }
    }
    println!("{}", inc);    
}

fn main() {
    part1();
    part2();
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
