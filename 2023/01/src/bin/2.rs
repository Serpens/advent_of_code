use std::io;
use std::io::prelude::*;
use regex::Regex;


fn string2digit(s: &str) -> i32 {
    let result = match s {
        "zero" => 0,
        "one" => 1,
        "two" => 2,
        "three" => 3,
        "four" => 4,
        "five" => 5,
        "six" => 6,
        "seven" => 7,
        "eight" => 8,
        "nine" => 9,
        &_ => 100
    };
    if result == 100 {
        panic!("{}", s)
    };
    result
}


fn main() {
    let pattern = r"(0|1|2|3|4|5|6|7|8|9|zero|one|two|three|four|five|six|seven|eight|nine)";
    let digits_regex = Regex::new(pattern).unwrap();
    let mut result = 0;

    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let mut first_str = "";
        let mut last_str = "";
        let str_line = line.expect("");
        
        let nchar = str_line.chars().count();
        for i in 0..nchar {
            let found = digits_regex.find(&str_line[i..nchar]);
            if found.is_none() {
                break;
            }
            if first_str == "" {
                first_str = found.expect("").as_str()
            }
            last_str = found.expect("").as_str();
        }

        let mut first = 0;
        if first_str.chars().count() == 1 {
            first = first_str.chars().next().unwrap() as i32 - 48;
        } else {
            first = string2digit(first_str);
        }

        let mut last = 0;
        if last_str.chars().count() == 1 {
            last = last_str.chars().next().unwrap() as i32 - 48;
        } else {
            last = string2digit(last_str);
        }
        result += first * 10 + last;
    }
    println!("{}", result);
}