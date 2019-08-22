fn main() {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    let s = s.trim_right().to_owned();
    let vec: Vec<&str> = s.split_whitespace().collect();
    let n: i32 = vec[0].parse().unwrap_or(0);
    let i: i32 = vec[1].parse().unwrap_or(0);
    println!("{}", format!("{}", n - i + 1));
}
