// input name and age from user and shwo it back

#include "std_lib_facilities.h"

int main()
{
	cout<<" Please enter your first namew and age\n";
	string first_name = "???";
	int age = 0;

	cin>>first_name>>age;
	cout<<"Hello,"<<first_name<<"(age"<<age<<")\n"; // this keas string and devindes into two parts with whitespace and puts first into first variable adn second into secondd variable. since << operator is senstive to type if we try to put in into variable of different type it will be zero
	return 0;
}