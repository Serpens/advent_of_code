use std::io;
use std::io::prelude::*;
use std::cmp;

fn main() {
    let stdin = io::stdin();
    let mut result = 0;
    for line in stdin.lock().lines() {
        let str_line = line.expect("");
        let colon_idx = str_line.find(":").expect("");
        let balls_str = &str_line[(colon_idx+1)..];
        let mut max_red = 0;
        let mut max_green = 0;
        let mut max_blue = 0;
        for split in balls_str.split(";") {
            for ball_split in split.split(",") {
                let ball_num: i32 = ball_split[
                        1..(ball_split[1..].find(" ").expect("") + 1)
                    ].parse().unwrap();
                if ball_split.contains("red") {
                    max_red = cmp::max(max_red, ball_num);
                } else if ball_split.contains("green") {
                    max_green = cmp::max(max_green, ball_num);
                } else if ball_split.contains("blue") {
                    max_blue = cmp::max(max_blue, ball_num);
                }         
            }
        }
        result += max_red * max_green * max_blue;
    }
    println!("{}", result);
}
