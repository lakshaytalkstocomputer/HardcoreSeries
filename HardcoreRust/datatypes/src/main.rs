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

    // Tuple
    let tup = (1,"2", '3');

    println!("First element in tuple: {} ", tup.0);
    println!("Second element in tuple: {} ", tup.1);
    println!("Third element in tuple: {} ", tup.2);

    // // Lets try to access invalid element in tuple
    // println!("Invalid element : {} ", tup.3);

    // Above code gives us an error

    // // Let us try to assign a value
    // tup[0] = 5;
    // println!("First element in tuple after assigning new value: {} ", tup.0);

    // Above code gives us error

    // tup = (7,8,9);
    // println!("First element in tuple after assigning new value: {} ", tup.0);
    // println!("Second element in tuple after assigning new value: {} ", tup.1);
    // println!("Third element in tuple after assigning new value: {} ", tup.2);

    // aAbove code gives us error

    // // Let stry to make tup mutable
    let mut tup = (1,2,3);
    println!("First element in mut tuple: {} ", tup.0);
    println!("Second element in mut tuple: {} ", tup.1);
    println!("Third element in mut tuple: {} ", tup.2);

    tup = (7,8,9);
    println!("First element in mut tuple after assigning new value: {} ", tup.0);
    println!("Second element in mut tuple after assigning new value: {} ", tup.1);
    println!("Third element in mut tuple after assigning new value: {} ", tup.2);

    // // let us try to acces the value of tup using square bracket
    //  println!("First element in mut tuple after assigning new value: {} ", tup[0]);
    // Above code gives us the error

    // Arrays
    let arr = [1,2,3,4,6];
    println!("First element in arr is : {}", arr[0]);

    // arr[0] = 10;
    // println!("First element in arr after assigning new value is : {}", arr[0]);
    // Above code gives us the error

    let mut new_arr = [11,12,13];
    println!("First element in arr is : {}", new_arr[0]);

    new_arr[0] = 1;
    println!("First element in mutable arr after assignning a new values is : {}", new_arr[0]);
    println!("Second element in mutable arr after assignning a new values is : {}", new_arr[1]);
    println!("Third element in mutable arr after assignning a new values is : {}", new_arr[2]);


}
