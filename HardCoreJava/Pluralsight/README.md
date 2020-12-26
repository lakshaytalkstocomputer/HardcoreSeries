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
    