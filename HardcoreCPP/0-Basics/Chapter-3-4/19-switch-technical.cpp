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
	
	int n;
	
	cout<<"Enter a integer from 0- 9: ";
	cin>>n;
	// lets convert integer to string
    
	switch(n)
	{
		case 0:
		{
			cout<<"zero\n";
		}
		break;
		case 1:
		{
			cout<<"one\n";
		}
		break;
		case 2:
		{
			cout<<"three\n";
		}
		break;
		case 4:
		{
			cout<<"four\n";
		}
		break;
		case 5:
		{
			cout<<"five\n";
		}
		break;
		case 6:
		{
			cout<<"six\n";
		}
		break;
		case 7:
		{
			cout<<"seven\n";
		}
		break;
		case 8:
		{
			cout<<"eight\n";
		}
		break;
		case 9:
		{
			cout<<"nine\n";
		}
		break;
		default:
		{
			cout<<"Wrong number. PLease select integer between (& incl.) 0 and 9\n";	}
		
	} 

    
	return 0;
}