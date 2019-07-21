/*
	In Switch statement: 
	 1. The value on which we switch must be of an 
	 	integer, char, or enumeration type.
	 	In particular, It MUST NOT be of STRING type.
	 2. The values in the case lebels must be constant expressions.
	 	In particular, you  CANNOT use VARIABLE in case label.
	 3. You CANNOT use the same value for two case labels.
	 4. You can use several case labels for a single case.
	 5. Don't forget to end each case with a break.
	 	Unfortunately, the compiler probalby won't wwarn you if you forget.


*/

#include "std_lib_facilities.h"

int main()
{
	/*
	1. The value on which we switch must be of an 
	 	integer, char, or enumeration type.
	 	In particular, It MUST NOT be of STRING type.
    */

    string n;
    cout<<"Do you like fish? \n";
    cin>>n;
    
    /*
    switch(n) // error : n is string we cannot use it!!!!!!
    {
    	case "Yes":
    		cout<<"Good!!\n";
    	break;
    	case "No":
    	    cout<<"No!\n";
    	break;
    	default:
    	cout<<"Wrong choice!";
    }
    */

    // To use string we have to use if else or map 
    char n1 = 0;
    if(n == "yes" || n == "Yes")
    	n1 = 'y';
    else if(n == "no" || n == "No")
    	n1 = 'n';

    switch(n1)
    {
    	case 'y':
    		cout<<"Good!!\n";
    	break;
    	case 'n':
    	    cout<<"No!!!!!!!! You should eat!\n";
    	break;
    	default:
    	cout<<"Wrong choice!";
   
    }
	return 0;
}