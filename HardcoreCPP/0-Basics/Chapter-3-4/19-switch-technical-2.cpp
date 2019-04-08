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
	2. The values in the case labels must be constant expressions.
	 	In particular, you  CANNOT use VARIABLE in case label.
	3.  You CANNOT use the same value for two case labels.
	*/

	char y = 'y'; // for yes
	constexpr char n = 'n'; // for No
	constexpr char m = 'm'; // for maybe
	cout<<"Do you like fish? (yes (y), no (n), or maybe(m)? : \n";
	char a;
	cin>>a;

	switch(a)
	{
		//case y:  error: Cannot use variable as case label!!! 
		case 'y':
		cout<<"Good!!!\n";
		break;
		
		case n: // okay, as the variable is constexpr now!!
		cout<<"WHy??? You should eat!!!!\n";
		break;
		
		case m:
		cout<<"Well You should try to like it!!\n";
		break;
		/*
		case 'm': error: Cannot use same calse labels as above case has same value
		cout<<"Try Try Try Hard!!! \n"; 
		*/		
		default:
		cout<<"Wrong choice!!!";
	}
	return 0;
}