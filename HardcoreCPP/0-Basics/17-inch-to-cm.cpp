// convert from inches to centimeters or
// centimeters to inches

#include "std_lib_facilities.h"

int main()
{
	constexpr double cm_per_inch = 2.54; 
	//no. of centimeter in an inch
	
	double length = 1; // length in inches or centimeters

	char unit = 0;
	cout<<"Please enter a length followed by a unit (c or i): ";
	cin>>length>>unit;

	/*
	if(unit == 'i')
	{
		cout<<length<<" in =="<<cm_per_inch*length<<"cm\n";
	}
	else
	{
		cout<<length<<"cm == "<<length/cm_per_inch<<"in\n";
	}
	*/

	if(unit == 'i')
	{
		cout<<length<<" in =="<<cm_per_inch*length<<" cm\n";
	}
	else if(unit == 'c')
	{
		cout<<length<<"cm == "<<length/cm_per_inch<<" in\n";
	}
	else
	{
		cout<<"Unknown unit, Please try again!\n";
	}
}
