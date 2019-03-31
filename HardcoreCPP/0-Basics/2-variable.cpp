// checking some simple variable initialization programs

#include "std_lib_facilities.h"

int main()
{
	string name = "Lakshay";
    int age = 17;
    float height = 144.4;
    double some_number = height * 2;
    
    cout<<name<<" "<<age<<" "<<height<<""<<some_number<<"\n";
    
    string name1;
    cout<<name1<<" \n"; // this oputputs empty 

	/*
		below code compiles and runs
	    but the output is not proper.
	    It shows funny things on console.
	*/
	// from here
	    name1 = 22;
	    //cout<<name<"  \n"; //error
	    cout<<name1<<"  \n";
		
		name1 = height;
		cout<<name1<<"  \n";
		
	    name1 = some_number;
		cout<<name1<<"  \n";
	// upto here

	/* This code has error: 
			cannot convert string to int 
			in assignment.

	age = name;
    cout<<"age: "<<age<<" \n";
	*/

 	/* code below this works fine 
 	   but loses some information
 	*/
	
	// from here	
		age = height;
		cout<<"age: "<<age<<" \n";

		age = some_number;
		cout<<"age: "<<age<<" \n";
	// upto here

	
	return 0;
}