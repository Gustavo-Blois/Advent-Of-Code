use std::{fs::File, io::Read};
macro_rules! number_map {
    () => {
        vec![
            ("1", '1'),
            ("2", '2'),
            ("3", '3'),
            ("4", '4'),
            ("5", '5'),
            ("6", '6'),
            ("7", '7'),
            ("8", '8'),
            ("9", '9'),
            ("one", '1'),
            ("two", '2'),
            ("three", '3'),
            ("four", '4'),
            ("five", '5'),
            ("six", '6'),
            ("seven", '7'),
            ("eight", '8'),
            ("nine", '9'),
        ]
    };
}
fn main() -> std::io::Result<()>{

    let mut file = File::open("input")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    let mut value = 0;
    let lines = contents.lines();
    let mut iterator = 0;
    for line in lines {
        let first:Option<u32> = get_first(&line);
        let last:Option<u32> = get_last(&line);

        if let (Some(first), Some(last)) = (first,last){
            value += first *10 + last;
            println!("Linha {} -> Primeiro {} -> Ultimo {} -> Resultado {}\n",iterator,first,last,(first*10+last));
        }
        iterator +=1;
    }
    println!("Total: {}",value);
    Ok(())
}

fn get_first(line:&str) -> Option<u32>{

    let digit_map = number_map!();
    for i in 0..line.len() {
        for &(word, digit) in &digit_map {
            if line[i..].starts_with(word) {
               return digit.to_digit(10);
            }
        }
    }
    return None;
}


fn get_last(line:&str) -> Option<u32> {
    
    let digit_map = number_map!();
    let mut last:Option<u32> = None;

    for i in 0..line.len(){
        for &(word,digit) in &digit_map {
            if line[i..].starts_with(word) {
                last = digit.to_digit(10);
            }
        }
    }
    return last;
}