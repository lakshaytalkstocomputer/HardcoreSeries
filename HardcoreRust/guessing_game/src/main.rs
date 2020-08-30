// Import as we say in python or include in C, C++
//  input output from std library.
use std::io;
use rand::Rng;

fn main(){
    // // print here is macro not a function, '!' at end
    // print!("Please enter a number: ");
    // println!();
    // // Create a variable to store the value entered by user.
    // //   Look String is not imported, this comes in something
    // //   that is known as "prelude" in rust.
    // let mut guess = String::new();
    //
    //
    // let output = io::stdin()
    //                         .read_line(&mut guess)
    //                         .expect("Error in reading!");
    //
    //
    // println!("you guessed {}", guess);
    // println!("OUt put was {}", output);
    //
    // // Experimenting
    // // println!("Please enter a number: {} ",io::stdin()
    // //                         .read_line(&mut guess)
    // //                         .expect("Error in reading!"));
    // // println!("You entered {}", guess);

    println!("Guess the number: ");

    let secret_number = rand::thread_rng()
                                    .gen_range(1, 101);

    println!("The secret number is : {}", secret_number);
    println!("Please input your guess! ");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed : {}", guess)
}

