use std::path::Path;
use std::error::Error;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::collections::HashMap;

// Main returns an io::Result<T>
fn main() -> io::Result<()> {

    // Specify file to open
    let path = Path::new("input.txt");
    let display = path.display();

    // Open file in read-only mode
    let file = match File::open(&path) {
        Err(why) => panic!("couldn't open {}: {}", display, why.description()),
        Ok(file) => file,
    };

    // Vector to hold frequency changes
    let mut changes = Vec::new();

    // Buffered reader to read each line in file
    let reader = BufReader::new(file);

    // Loop over each line
    for line in reader.lines() {
        let mut s = line.unwrap();
        // Use parse to convert string to i32 and then push
        // onto the changes vector.
        match s.parse::<i32>() {
            Err(why) => panic!("couldn't convert string to i32: {}", why.description()),
            Ok(change) => changes.push(change),
        }
    }

    // Mutable i32 to hold total frequency
    let mut frequency: i32 = 0;
    // HashMap to hold each time a fequency is seen
    let mut seen = HashMap::new();
    // Index for changes vector
    let mut i = 0;
    // Loop until duplicate frequency is seen
    loop {
        frequency += changes[i];

        // Get or insert frequency into seen HashMap
        let count = seen.entry(frequency).or_insert(0);
        // Increment the frequency seen
        *count += 1;

        // Found the duplicate frequency
        if *count > 1 {
            println!("{}", frequency);
            break;
        }

        // Increment index or reset to zero
        i += 1;
        if i == changes.len() {
            i = 0;
        }
    }

    Ok(())
}
