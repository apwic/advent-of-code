use std::fs::File;
use std::io::{self, BufRead, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);
    let mut sum: i32 = 0;

    for line in reader.lines() {
        let line = line?;

        let strings: Vec<&str> = line.split("").collect();
        let n = strings.len();

        let mut a: i32 = 0;
        let mut b: i32 = 0;

        let mut i = 0;
        let mut j = n - 1;

        while i <= j {
            if a != 0 && b != 0 {
                sum += a * 10 + b;
                break;
            }

            if a == 0 {
                match strings[i].parse::<i32>() {
                    Ok(x) => a = x,
                    Err(_) => i += 1,
                }
            }

            if b == 0 {
                match strings[j].parse::<i32>() {
                    Ok(x) => b = x,
                    Err(_) => j -= 1,
                }
            }
        }
    }

    println!("{}", sum);

    Ok(())
}
