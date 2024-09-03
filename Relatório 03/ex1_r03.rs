use std::io;
fn main() {
    let mut n1_input = String::new();
    let mut n2_input = String::new();
    let mut op = String::new();
    let result: i32;
    let num1: i32; // cria number1 como inteiro 32 bits
    let num2: i32; // cria number2 como inteiro 32 bits

    io::stdin()
        .read_line(&mut n1_input)
        .expect("failed to read line");
    io::stdin()
        .read_line(&mut n2_input)
        .expect("failed to read line");
    io::stdin().read_line(&mut op).expect("failed to read line");

    num2 = n1_input.trim().parse().unwrap();
    num1 = n2_input.trim().parse().unwrap();

    if op.trim() == "+" {
        result = num1 + num2;
        println!("Soma {} + {} = {}", num1, num2, result);
    } else {
        result = num1 * num2;
        println!("Multiplicacao {} * {} = {}", num1, num2, result);
    }
}
