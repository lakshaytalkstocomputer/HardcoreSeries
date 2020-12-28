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

## a) Variables
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

## b) Primitive DataType
* Prmitive Data Types are the foundation of all other types. 
* Various primitive data types are :-
    * Integer
    * Floating Point
    * Character
    * Boolean

### i) Integer 
* These are also of four types :-
    * byte   `byte numOfEnglishLetter = 26;`
    * short  `short feetInAMile = 5280;`
    * int    `int milesToSun = 92960000;`
    * long   `long milesInAlightYear = 5879000000000L`
* We need to mention L in long otherwise it will be taken as int.
* [Reference for : Why do we need to mention L even after ddeclaring the variable as long?](https://stackoverflow.com/questions/13134956/what-is-the-reason-for-explicitly-declaring-l-or-ul-for-long-values)

### ii) Float
* There are two types :-
    * float  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `float somefloat  = 42.15f;`
    * double &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `double somDOuble = 0.0000;`
* We can also add suffix `d` for double variables.

### iii) Character
* Single unicode.
* Uses single quotes.
* For characters not present on keyboard we can use `\u` followed by 4-digit hex value.

### iv) Boolean
* Stores true/false.

## c) How prmitive Data types are stored?
* Primitive Stored By Value not By reference.

## d) Arithmetic Operators
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

## e) Type Conversion
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

## a) Block Statement
* Group together multiple statements.
* Variable scope limited to that block.

# 6. Looping and Arrays
## a) Loops
* for loop
    * normal for loop
        * `for( <intiliazation> ; <condition> ; <updator>) {}`
    * for each
        * `for(dataType loop-variable: array){}`
* while loop
* do-while loop

## b) Arrays
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

# 9. Classes
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

### ii) Encapsulation and Access Modifiers
* Different Access Modifiers are :-
    * Private  ( only inside the class)
    * Public   ( from anywhere)
    * default  (only within same package)

### iii) Special Reference
* this ( contenxt fo current object)
* null ( no reference to anything)

### iv) Field Value
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

###  Field Initializers
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

### Constrcutors
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

### Initialization Block
* Share code across all constrcutors
* Place code within brackets outside of any method or cosntrcutor
* Does not include any parameter list.


## b) Static Members
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
### i) Static Import statement
* Used with static methods
* Allows method anme to be used without being class - qualified
```
import static com.package.name.ClassName.staticMethodName;

staticMethodName();
```  
* _import static com.package.name.ClassName.*_ allows all the static members  methods to be imported.   


## c) Parameter Passing
* By Value - Prmitive Type
* By Reference - Objects

## d) Method Overloading
* Signature Checking
* One Method os name _someMethod_ can call other method _someMethod_ which is an overload.
* Automatic Casting works in method overloading as well.

## e) Variadic Function
* Last parameter should be the variadic paramter.


# 10) Inhertence
* We can leverage one class to ahve features from other class using inheritence.
* One class can hide or override the members of its base class.
* Java Object Class.
* Equality checks in inhertence.
* A class can extend only one class

## a) How to inherit from one class to another?
* Use _extends_ keyword.
* When starting out derived class starts with all the characterstics of base class

## b) What are features of new class that we can use and what dictates it?
* Let us take an example.
    * There is an base class named _SomeBaseClass_.
    * There is an derived class named _SomeDerivedClass_.
    * Obviously, derived class is created using _extends_ keyword from base class.
    * Now if we create an object of _SomeDerivedClass_ and put it into the variable of type _SomeBaseClass_, the newly added features of _SomeDerivedClass_ will not work.
    ```
    SomeBaseClass obj = new SomeDerivedClass();
    ```
## c) Member Overriding and Hiding
* Member Attributes are hidden.
* Member Methods are overridden.
* Member Attributes use the type of reference or variable type.
* Member Methods use the type of instance.

## Object Class
* Every class automatically extends Object Class.

## super Keyword
* Can use to refer to base class.

## final keyword
* We can prevent inheritence by marking the class final.
* We can prevent method overriding by marking it as final.

## abstract keyword
* If we mark a method abstract, we need to mark the class abstract as well.
* This means we cannot create instance of that class.
* We can have reference of that class, we can extend it, but cannot create object of this class.

## Constructors in inheritence.
* constructors are not inherited.
* When we create a derived class instance, base class constructor is always called. By default no-argumetn version is called.
* We can excplicitly call a base constructor by usign the super keyword.
* This hsould be the first line of the constructor.

# 11) Enums
* Enums are like Classes
* They provide two methods
    * value
    * valueOf
* The order in which enums are provided matter if we want to check bigger or smaller.
* Enum Types are Class bu they inherit from JAva Enum Class.
* They can have fields and methods.
* Every member is an instance of that type.
* When defining instance in an Enum, we can call constructor.

# 12) Interfaces
* The need for more than inheritence.
* Implementing an interface.
* Generic Interfaces
* Implementing multiple interfaces
* Declaring an interface
* Default methods

## What is Interface and why do we need them?
* Interfaces defines a contract
* It provides a list of operations.
* Does not focus on implementation details.
* Classes implement interfaces.
* Classes conform to the contract.
* Classes provide necesaary methods.

## Implementing a interface
* We can implement an interface using _implements_ keyword.
* A class can implement multiple interfaces.

## Generic Interface
* We can create a generic interface which can take in DataType while implementing 
that.
* _Comparable_ is a type of generic interface which provide compareTo functionality 
to allow comparing of instances or ordering of instances.
```
class SomeClass implements Comparable<SomeClass> {}
```
## Declaring an interface
* Use _interface) keyword instead of _class_ keyword.
* Define one or more methods with **Name**, **parameters**, **return type**.
* Implicitly public 
* Fields are constant , public, final and static.
* Interface can extend other interface using _extend_ keyword.
* When classes implemnt derived interface that class need to implment al the methods of base interface as well as derived interface.
* Interface allow default method implementation using __default__ keyword in interfaces.

# 12) Nested Class
* Declaring one type wihtin another type.
* Nesting type for naming scope.
* Inner classes.
* Anonymous classes.

##  Nested Class
* A type declared within another type
* Are members of the enclosing type
    * Nessted type can access private members of enclosing class
* Nested types support all access modifiers.
    * No modifier (a.k.a package private)
    * public
    * private
    * protected
* Two categories of nested classes.
    * one provided only naming scope
    * the other links type instances
* Both categor are similar in syntax but different in behaviour
* We use nested class to naming scope.
    * Type name scoped within enclosing type. (static class)
    * No relationship between nested type and enclosing trype instance. ( non static class)
* Applies to the following nested type,
    * Static classes nestedd within classes
    * all classes nested wihtin interfaces
    * All nested interfaces.
* Allows Anonymous Class
    * Declared as part of their creation
    * Use as simple interface implementaions
    * Use as simple class extensions
    * Anonymous Classes arte inner classes
        * Instance is assosciated with enclosing class instance
        * 

# 13) The double colon (::) operator, 
* This is also known as method reference operator in Java, is used to call a method by referring to it with the help of its class directly.
* Method reference or double colon operator can be used to refer:
    * a static method,
    * an instance method, or
    * a constructor.
