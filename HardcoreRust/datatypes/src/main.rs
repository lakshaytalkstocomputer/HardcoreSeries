fn main() {
    // Integer Type
    let first_integer : u32  = 1;
    println!("32 bit unsigned integer : {}", first_integer);

    let second_integer : u8 = 255;
    println!("8 bit unsigned integer: {}", second_integer);

    // // Let's do integer overflow!
    // // // Let's set the value of second_integer to 256

    // let second_integer : u8 = 256;
    // println!("8 bit unsigned integer with value set to 256: {}", second_integer);

    // Above code gives us error in debug mode but would be "wrapped" in
    //   prod build.

    // Floating Type

    // // Be default rust gives us Floating point 64
    // let first_float = 2.2;

    // // Lets check in operations

    // let add_op: u64 = 2.2 + 6;
    // print!("Add op : {}", add_op);

    // Above code is an error

    let divide_op_float_float = 6.0 / 3.0;
    println!("Divide op float with float  : {}", divide_op_float_float);

    // let divide_op_flaot_int = 6.0 / 3;
    // println!("Divide op float with int : {} ",divide_op_flaot_int );
    // Above code is an error

    let mult_op_float_float = 6.0 * 3.0;
    println!("Multiply op float with float  : {}", mult_op_float_float);



}
