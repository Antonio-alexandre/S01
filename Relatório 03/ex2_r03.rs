fn preenche_arr(arr: &mut [i32; 10], valor: i32) {
    for i in 0..arr.len() {
        arr[i] = i as i32 * valor;
    }
}

fn main() {
    let mut array: [i32; 10] = [0; 10];
    let valor = 3;

    preenche_arr(&mut array, valor);

    println!("Array preenchido: {:?}", array);
}
