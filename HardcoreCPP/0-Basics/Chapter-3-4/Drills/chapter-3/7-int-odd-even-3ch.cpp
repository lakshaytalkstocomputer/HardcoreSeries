#include "std_lib_facilities.h"

int main()
{
	int num;
	cout<<"Enter a number: ";
	cin>>num;
	if(num%2 == 0)
	{
		cout<<num<<" is an even number.\n";
	}
	else
	{
		cout<<num<<" is a odd number.\n";
	}
	return 0;
}