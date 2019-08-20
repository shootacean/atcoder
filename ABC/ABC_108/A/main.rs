use std::io::{self};

fn main() {
    let k: i32 = read_line().parse().unwrap();
    println!("{}", format!("{}", solve(k)));
}

fn solve(n: i32) -> i32 {
    let e = get_even_count(n);
    return e * (n - e);
}

fn get_even_count(n: i32) -> i32 {
    n / 2
}

fn read_line() -> String {
    let mut s = String::new();
    io::stdin().read_line(&mut s).unwrap();
    return s.trim_right().to_owned();
}