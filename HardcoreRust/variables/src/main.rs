fn main() {
    /* Mutability adn Immutability */
    let mut mutable_number = 9;
    println!("The value of mutable_number is: {} ", mutable_number);

    mutable_number = 27;
    println!("The value of mutable_number has been changed to {} ", mutable_number);

    mutable_number = 30;
    println!("The  value of mutable_number has again been changed to  {}", mutable_number);

    let immutable_number = 10;
    println!("The Value of immutable_number: {}", immutable_number);

    // this is an error
    // // immutable_number = 11;

    /* Constants */

    const constant_number : u32 = 111;

    println!("COnstant Number: {} ", constant_number);

    // this is is not allowed as variable of this already exists as constant Number
    // // let constant_number = 222;
    // // println!("Constant Number After shadowing it with let : {} ", constant_number);

    // below is not allowed. Why? because const are immutable and cannot be changed to mutable!
    // // const mut another_constant_number :u32 = 122;

    // below is also not allowed. Why? Because not datatype!
    // // const another_constant_numebr = 123;

    /* Shadowing */

    let some_number = 1;
    println!("Value of some_number is : {}", some_number);

    let some_number =2;
    println!("Value of some_number has been changed to : {}", some_number);

    // Why Shadowing?  Why not just use x = 5?
    // // Since we cannot change the value of variable if mut is not mentioned
    // //  Therefor e we need to mention mut.

    let mut shadow_number = 1010;
    println!("Shoadow Number : {}", shadow_number)
    // // this is not allowed! Why? Because it changes the type of the variable
    // // shadow_number = "hello";

}
