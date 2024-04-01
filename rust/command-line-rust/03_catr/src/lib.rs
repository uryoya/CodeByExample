use clap::{App, Arg};
use std::error::Error;
use std::fs::File;
use std::io::{self, BufRead, BufReader};

type MyResult<T> = Result<T, Box<dyn Error>>;

pub fn run(config: Config) -> MyResult<()> {
    for filename in config.files {
        match open(&filename) {
            Err(err) => eprintln!("Failed to open {}: {}", filename, err),
            Ok(file) => {
                if config.number_nonblank_lines {
                    readlines_with_number_nonblank(file)?
                } else if config.number_lines {
                    readlines_with_number(file)?
                } else {
                    readlines(file)?
                }
            }
        }
    }
    Ok(())
}

#[derive(Debug)]
pub struct Config {
    files: Vec<String>,
    number_lines: bool,
    number_nonblank_lines: bool,
}

pub fn get_args() -> MyResult<Config> {
    let matches = App::new("catr")
        .version("0.1.0")
        .author("Ryoya Urano")
        .about("Rust cat")
        .arg(
            Arg::with_name("file")
                .value_name("FILE")
                .help("Input file(s)")
                .multiple(true)
                .default_value("-"),
        )
        .arg(
            Arg::with_name("number")
                .short("n")
                .long("number")
                .help("Number lines")
                .takes_value(false)
                .conflicts_with("number_nonblank"),
        )
        .arg(
            Arg::with_name("number_nonblank")
                .short("b")
                .long("number-nonblank")
                .help("Number nonblank lines")
                .takes_value(false),
        )
        .get_matches();

    Ok(Config {
        files: matches.values_of_lossy("file").unwrap(),
        number_lines: matches.is_present("number"),
        number_nonblank_lines: matches.is_present("number_nonblank"),
    })
}

pub fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {
    match filename {
        "-" => Ok(Box::new(BufReader::new(io::stdin()))),
        _ => Ok(Box::new(BufReader::new(File::open(filename)?))),
    }
}

pub fn readlines(body: Box<dyn BufRead>) -> MyResult<()> {
    for line in body.lines() {
        println!("{}", line?);
    }
    Ok(())
}

pub fn readlines_with_number(body: Box<dyn BufRead>) -> MyResult<()> {
    for (i, line) in body.lines().enumerate() {
        println!("{:>6}\t{}", i + 1, line?);
    }
    Ok(())
}

pub fn readlines_with_number_nonblank(body: Box<dyn BufRead>) -> MyResult<()> {
    let mut i = 1;
    for line in body.lines() {
        match &*line? {
            "" => println!(""),
            line => {
                println!("{:>6}\t{}", i, line);
                i += 1;
            }
        }
    }
    Ok(())
}
