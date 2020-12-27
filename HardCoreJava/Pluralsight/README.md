# 1. Statement Structure
*  Statement ends with semicolon.

# 2. Comments
*  Supports three type of comments
*  a) Line Comment     // 
*  b) Block Comment    /* */
*  c) JavaDoc Comment  /** */

Note : We can put line comment in block comment but not block comment in block comment.

# 3. Packages
*   a) All lowercase
*   b) Use reverse domain name notation to assure global uniqueness. 
*   c) Add firther qualifier to assuire uniqueness within a company or group.
*   d) Type Names are qualified by their Package 
       Main becomes com.pluralsight.example.Main
*   e) Package Name affect source code organization 
```
    src
     |___com
           |___example
                     |___Main        
```

# 4. Variables and Data Types
*    a) variable and valid variable names and access modifiers
*    b) Prmitive Data Type
*    c) how prmitive Data types are stored?
*    d) Arithmetic operators
*    e) Data Type Conversions

## Variables
* Named Storage.
* Strongly Typed.
* *final* modifier to disable change of value.
* What are different ways to intilize a variable?
```
   a) int somevariable;
      someVariable = 100;
   b) int someVariable = 100;
```
* Variable Names cannot start with numbers.
* When using `final` keyword, we cannot change the value of variable once set.
```
        final int somevariable;

        int someOtherVariable = 100;
        somevariable = someOtherVariable;
```
* Above code works because value of `someVariable` was not set. But we cannot again try to set the value of the `someVariable`.

## Primitive DataType
* Prmitive Data Types are the foundation of all other types. 
* Various primitive data types are :-
    * Integer
    * Floating Point
    * Character
    * Boolean

### Integer 
* These are also of four types :-
    * byte   `byte numOfEnglishLetter = 26;`
    * short  `short feetInAMile = 5280;`
    * int    `int milesToSun = 92960000;`
    * long   `long milesInAlightYear = 5879000000000L`
* We need to mention L in long otherwise it will be taken as int.
* [Reference for : Why do we need to mention L even after ddeclaring the variable as long?](https://stackoverflow.com/questions/13134956/what-is-the-reason-for-explicitly-declaring-l-or-ul-for-long-values)

### Float
* There are two types :-
    * float  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `float somefloat  = 42.15f;`
    * double &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `double somDOuble = 0.0000;`
* We can also add suffix `d` for double variables.

### Character
* Single unicode.
* Uses single quotes.
* For characters not present on keyboard we can use `\u` followed by 4-digit hex value.

### Boolean
* Stores true/false.

## How prmitive Data types are stored?
* Primitive Stored By Value not By reference.

## Arithmetic Operators
* There are three type of Arithmetic Operators that act on Primitive Data Types
    * Basic
        * Produces a result.
        * No impact on value used in the operation.
    * Prefix/PostFix
        * Increase or decrease a value.
        * Replace orginal value.
    * Compound Argument
        * Operate on a value.
* Operator Precedence
    * Follows this order :
        * PostFix
        * Prefix
        * Multiplicative
        * Additive
    * Override Precedence with parenthesis.
    * Equal precendence evaluated left to right.
    * Nested Parenthesis evaluated from inside out.

## Type Conversion
* There are two ways with which we can convert the data type.
    * Implicit Type Conversion
        * Conversion done automatically by compiler.
    * Exmplicit Type COnversion 
        * Conversion performed explicitly in code using cast operator.
        ```
        long some = 1212L;
        int someIntValue = (int) some;
        ```
# 5. Conditional Logic
* Relational Operators
* Conditonal Assignment ``` condition ? true-value ? false-value```
* if-else
* if-else_if-else
* Switch with break

## Block Statement
* Group together multiple statements.
* Variable scope limited to that block.

# 6. Looping and Arrays
## Loops
* for loop
    * normal for loop
        * `for( <intiliazation> ; <condition> ; <updator>) {}`
    * for each
        * `for(dataType loop-variable: array){}`
* while loop
* do-while loop

## Arrays
* Collection of elements of same type
* zero based indexing
* `.length` attribute for knowing the length.
* Initialize array :
    * Ways to intialize an array :
        * Directly:
            * `float[] someArray = { 0.01f, 0.02f}`
            * Java will automatically set the length.
        * Normal
            ```
            int[] someArr = new int[3];
            someArr[0] =1;
            someArr[1] =1;
            someArr[2] =1;
            ```           

# 7. Methods
* Named Parameter List
* Allows Empty Parameter List
* Can return something or return void.
* Variables are scoped to the method.
* Parameters are passed as value.
* Allows usage of empty return in which case the return type shoudl be void.


# 8. Strings
* String Class
    * Stores a sequesnce of Unicode characters.
    * Literals are stored in double quotes.
    * Allows += operations.
    * String objects are immutable.
        * Variables hold the reference to the value.
    * String Equality
        * Use Equals method to check whether every character in two strings are same.
        * Use == to check if two variables hodl same reference.
    * Interning a string.
        ```
        String s1 ="I LOVE"
        String s2 = "I"
        s2 += " LOVE"
        String s3 = s1.intern();
        String s4 = s2.intern();

        s1 == s2 (false)
        s3 == s4 (true)
        ```
        * This provides a canonicalized value.
        * Enables reliable == comparison.
        * when an intern method is apllied to a string, java looks if a intern object of that string is created, if not it creates and gives the reference.
* String Builder
    * Need to extract strign from Stringbuilder Object
    * Mutable Strings
    * Automatic length adaptation.
    ```
    String location ="Florida";
    int flightNumber = 175;
    StringBuilder sb = new StringBuilder(40);

    sb.append("I flew to");
    sb.append(location);
    sb.append(" on flight #);
    sb.append(flightnumber);

    String message = sb.toString();
    ```
    * Allows `indexOf`,`insertAt` methods. 

# 9. Classes and Inheritence
* Declaring Classes
* Class Mmebers
* Working with objects
* Encapsulation and access modifiers
* Field Accessors and mutators

## a) Declaring Classes
* Declared with Class Key Word followed by the class name.
* Class Name should be Capital.
* Body of class is contained within brackets.
* Java source file name normallu has the name as class.

### i) Class
* A class is mad eup of state and executable code.
* We use memeber to represent these. Following are the members:-
    * Fields (Store object state)
    * Methods (Executable code that manipulates state and perform operations)
    * Constructors (Executable code used during object creation to set initial state)
    
* Creating variable of type _MyClass_  will not create object but a place to hold reference.
    * Actual object is created using _**new**_ keyword.
* Classes are known as reference types.
* WHen you create an array of __MyClass__ , you are not creating arrya of objects,
you are creating an array which can hold refernce to four objects. You need to create those four objects specifically.
```
MyClass[] someVar = new MyClass[4];

This creates an array which can hold refernce to 4 objects of type MyClass. Objects are not created yet.
```

## ii) Encapsulation and Access Modifiers
* Different Access Modifiers are :-
    * Private  ( only inside the class)
    * Public   ( from anywhere)
    * default  (only within same package)

## iii) Special Reference
* this ( contenxt fo current object)
* null ( no reference to anything)

## iv) Field Value
* Java sets the default value of the fields.
* Various default values of types are :-
    * byte, short, int, long -> 0
    * float, double -> 0.0
    * char -> '/u0000'
    * boolean -> false
    * Reference Types -> null
* We can change the default value of the fields usign following ways:-
    * Field Initializers
    * Constrcutors
    * Intialization blocks

### a) Field Initializers
* Specifiy field's initial value as part od the field's declaration
* Can be an equation.
* can include method fields.
* can include method calls.
```
public class Earth{
    long circumferenceInMiles = 294901;
    long circumferenceInKms = Math.round(circumferenceInMiles * 1.6d);    
}
```

### b) Constrcutors
* Name of constructor is the Name of the class
* Every class needs to have atleast one constructor.
* If we dont provide one, java automatically creates one for the class which is public.
* We can have any number of constructors we need but we need to take care of the following:-
    * Each must have unique aprameter list
    * Different number of parameters
    * Different parameter types
* Chaining constuctor
    * Must be first line of constuctor
    * Use the this keyword followed by parameter list
    ```
    public class Passenger{
        public Passenger(int freeBags){
            this(freeBags > 1? 25.0d : 50.0d);
            this.freeBags = freeBags;
        }

        public Passenger(int freeBags, int checkBags){
            this(freeBags);
            this.checkedBags = checkedBags;
        }

        public Passenger(double perBagFee){
            this.perBagFee = perBagFee;
        }
    }
    ```
* Constructor Visibility
    * Cosntructors can be non public

### c) Initialization Block
* Share code across all constrcutors
* Place code within brackets outside of any method or cosntrcutor
* Does not include any parameter list.


## v) Static Members
* Static memebers are shared class - wide
    * Not assoscaited with individual instacne
* Decalred using the static keyword
    * Accessible using the class name
* Diferent types of Static members:
    * Field
        * A value not assosciated with a apscific isntance.
        * All instances access the same value.
    * Method
        * Performs an action not ties to a specific instance.
        * Has access to only static members.
### a) Static Import statement
* Used with static methods
* Allows method anme to be used without being class - qualified
```
import static com.package.name.ClassName.staticMethodName;

staticMethodName();
```  
* _import static com.package.name.ClassName.*_ allows all the static members  methods to be imported.   