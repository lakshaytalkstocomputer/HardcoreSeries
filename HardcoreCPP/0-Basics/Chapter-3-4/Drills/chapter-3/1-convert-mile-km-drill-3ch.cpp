// crate a promt to entere number of miles and convert that and show in kilmeters

#include "std_lib_facilities.h"

int main()
{
	cout<<"Please enter number of miles travelled: ";
	float miles;
	cin>>miles;
	cout<<"You travelled "<<miles*1.609<<" km(s)\n";
	return 0;
}