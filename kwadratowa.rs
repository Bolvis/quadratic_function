use std::{io, fmt, error};
use std::num::ParseFloatError;
use std::fmt::{Display, Formatter};

enum QuadraticResult {
    Double(f32, f32),
    Single(f32),
    None
}

impl Display for QuadraticResult {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), fmt::Error> {
        match self {
            QuadraticResult::Double(x1, x2) => write!(f, "{}, {}", x1, x2),
            QuadraticResult::Single(x) => write!(f, "{}", x),
            QuadraticResult::None => write!(f, "no result"),
        }
    }
}

#[derive(Debug)]
enum QuadraticError {
    IoError(io::Error),
    NumberParse(ParseFloatError),
    AssumptionViolation,
}

impl error::Error for QuadraticError { }

impl Display for QuadraticError {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), fmt::Error> {
        match self {
            QuadraticError::IoError(io) => write!(f, "IO error: {}", io),
            QuadraticError::NumberParse(parse) => write!(f, "Input error: {}", parse),
            QuadraticError::AssumptionViolation => write!(f, "a cannot be 0"),
        }
    }
}

impl From<io::Error> for QuadraticError {
    fn from(io_error: io::Error) -> Self {
        QuadraticError::IoError(io_error)
    }
}

impl From<ParseFloatError> for QuadraticError {
    fn from(parse_error: ParseFloatError) -> Self {
        QuadraticError::NumberParse(parse_error)
    }
}

fn main() {
    match execute() {
        Ok(result) => println!("Result: {}", result),
        Err(error) => eprintln!("Error: {}", error),
    }
}

fn execute() -> Result<QuadraticResult, QuadraticError> {
    println!("A:");
    let a = read_number()?;
    println!("B:");
    let b = read_number()?;
    println!("C:");
    let c = read_number()?;

    if a == 0.0 { return Err(QuadraticError::AssumptionViolation) }

    Ok(match (b * b) - (4.0 * a * c) {
        delta if delta < 0.0 => QuadraticResult::None,
        delta if delta == 0.0 => QuadraticResult::Single(-b / (2.0 * a)),
        delta => {
            let delta_sqrt = delta.sqrt();
            QuadraticResult::Double((-b - delta_sqrt) / (2.0 * a), (-b + delta_sqrt) / (2.0 * a))
        },
    })
}

fn read_number() -> Result<f32, QuadraticError> {
    let mut string = String::new();
    io::stdin().read_line(&mut string)?;
    Ok(string.trim().parse()?)
}
