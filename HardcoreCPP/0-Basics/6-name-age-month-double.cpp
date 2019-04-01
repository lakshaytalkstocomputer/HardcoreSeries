// read the age and first anme adn print age in months adn take age in double
#include "std_lib_facilities.h"

int main()
{
	cout<<"Please enter name and age: ";
	float age=0;
	string first_name="???";

	cin>>first_name>>age;

	int age1 = age;
	age1 = age1*12 + (age - age1)*10;
	cout<<"Hello "<<first_name<<",( age : "<<age1<<" months )\n";
	return 0;
}