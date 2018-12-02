use std::path::Path;
use std::error::Error;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;

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
    // Loop over changes vector and increment frequency
    for i in &changes {
        frequency += i;
    }

    println!("{}", frequency);
    Ok(())
}
