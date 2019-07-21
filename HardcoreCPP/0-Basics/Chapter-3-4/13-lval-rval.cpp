#include "std_lib_facilities.h"

int main()
{
	int length = 20;
	int width =40;

	length = 99; // length is lvalue as it refers to 
				//  the object which is named length

				 // 99 is the rvalue as i t refers to
				//  the value

	width = length;
				// here length is rvalue as it refers to 
			   // the value in the object named length

				// here width is lvalue as it refers to 
			   // the object named width

	width = 100;
	/*
		We can think of object as a box (as it has memory) which has a type and name
			int
		  -----------------------
   length |						|
		  |			99			|
		  -----------------------
	*/
  	(width > length? width : length) =789 ;
  				// Will above work?  
  					// YES!!! Since ternary operator 
  					// preserves lvalue-ness     
    cout<<width<<" "<<length; 
	return 0;
}