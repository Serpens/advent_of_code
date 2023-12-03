use std::io;
use std::io::prelude::*;

fn main() {
    let digits = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
    let mut result = 0;

    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let mut first = 0;
        let mut last = 0;
        let str_line = line.expect("");
        for c in str_line.chars() {
            if digits.contains(&c) {
                first = c as i32 - 48;
                break;
            }
        }
        for c in str_line.chars().rev() {
            if digits.contains(&c) {
                last = c as i32 - 48;
                break;
            }
        }
        result += first * 10 + last;
    }
    println!("{}", result);
}