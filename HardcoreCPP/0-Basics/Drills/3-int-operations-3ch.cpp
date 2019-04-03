// Prompt a user to input 2 int values and perform 
// operations on that

#include "std_lib_facilities.h"

int main()
{
	cout<<"Enter first integer: ";
	int val1, val2;
	cin>>val1;
	cout<<"Enter first integer: ";
	cin>>val2;

	if(val1 > val2)
	{
		cout<<"val1 is greater.\n";
	}
	else if(val2> val1)
	{
		cout<<"val2 is greater.\n";
		
	/*	// swap the variable to always have positivve difference
		val2 = val2+val1; 
		val1 = val2-val1;
		val2 = val2-val1;
   */
	}
	else
	{
		cout<<"Both are equal bro.\n";
	}

	cout<<"Sum : "<<val1+val2<<"\n";
	cout<<"Difference : "<<val1-val2<<"\n";
	cout<<"Product : "<<val1*val2<<"\n";
	float valF {val1};
	cout<<"Ratio : "<<valF/val2<<": 1\n";

	return 0;
}