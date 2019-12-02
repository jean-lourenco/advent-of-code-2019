use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let fuel = calculate(&calculate_fuel);
    assert!(3297896 == fuel);
    println!("Part One: {}", fuel);

    let fuel = calculate(&calculate_fuel_rec);
    assert!(4943969 == fuel);
    println!("Part Two: {}", fuel);
}

fn calculate(calc: &dyn Fn(i32) -> i32) -> i32 {
    let file = File::open("input.txt").unwrap();
    BufReader::new(file)
        .lines()
        .map(|x| x.unwrap().parse::<i32>().unwrap())
        .map(|x| calc(x))
        .sum()
}

fn calculate_fuel(mass: i32) -> i32 {
    mass / 3 - 2
}

// TODO: implement Tail Recursion
fn calculate_fuel_rec(mass: i32) -> i32 {
    let fuel = calculate_fuel(mass);

    if fuel > 0 {
        fuel + calculate_fuel_rec(fuel)
    } else {
        0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn cases_part_one() {
        assert_eq!(calculate_fuel(12), 2);
        assert_eq!(calculate_fuel(14), 2);
        assert_eq!(calculate_fuel(1969), 654);
        assert_eq!(calculate_fuel(100756), 33583);
    }

    #[test]
    fn cases_part_two() {
        assert_eq!(calculate_fuel_rec(14), 2);
        assert_eq!(calculate_fuel_rec(1969), 966);
        assert_eq!(calculate_fuel_rec(100756), 50346);
    }
}