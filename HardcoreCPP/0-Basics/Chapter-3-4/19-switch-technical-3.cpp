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
	4. You can use several case labels for a single case.
	*/
    cout<<"PLease enter a digit\n";
    char a;
    cin>>a;
    switch(a)
    {
    	case '0':case'2':case'4':case'6':case'8':
    		cout<<"Number is even\n";					
    	break;
    	
    	case '1':case'3':case'5':case'7':case'9':
    			cout<<"Number is odd\n";					
       	break;
    	default:
    	cout<<"Wrong number!!";
    }

	// Lets check a number is even or odd
	return 0;
}