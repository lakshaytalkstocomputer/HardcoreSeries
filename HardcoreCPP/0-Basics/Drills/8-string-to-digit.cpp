#include "std_lib_facilities.h"

int main()
{
	cout<<"Enter a spelled out number: ";
	string number;
	cin>>number;

	if(number == "zero" || number == "Zero")
	{
		cout<<"Digit: 0\n";
	}
	else if(number == "one" || number == "One")
	{
		cout<<"Digit: 1\n";
	}
	else if(number == "two" || number == "Two")
	{
		cout<<"Digit: 2\n";
	}
	else if(number == "three" || number == "Three")
	{
		cout<<"Digit: 3\n";
	}
	else if(number == "four" || number == "Four")
	{
		cout<<"Digit: 4\n";
	}
	else
	{
		cout<<"Not a number I Know, Sorry!\n";
	}
}