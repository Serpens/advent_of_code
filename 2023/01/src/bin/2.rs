use std::io;
use std::io::prelude::*;
use regex::Regex;

fn main() {
    let pattern = r"(0|1|2|3|4|5|6|7|8|9|zero|one|two|three|four|five|six|seven|eight|nine)";
    let digits_regex = Regex::new(pattern).unwrap();
    let mut result = 0;

    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let mut first = 0;
        let mut last = 0;
        let mut str_line = line.expect("");
        
        for i in digits_regex.find_iter(&str_line) {
            println!("{:?}", i);
        }

        println!("{}", first * 10 + last);
        result += first * 10 + last;
    }
    println!("{}", result);
}