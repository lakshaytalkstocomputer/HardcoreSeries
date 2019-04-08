#include "std_lib_facilities.h"

int main()
{
	//int a = 197;
	//int b = 10;

	//double d = double(a) / b;
	//char ch = char(a);
	//cout<<"d : "<<d<<"\n";
	//cout<<"ch: "<<ch<<"\n"; 

	double d3 = 2.5;
	cout<<"d3 : "<<d3<<"\n";
	int i = 2;
	cout<<"i : "<<i<<"\n";

	double d2 = d3/i;
	cout<<"d2 : "<<d2<<"\n";

	int i2 = d3/i;
	cout<<"i2 : "<<i2<<"\n";
	
	int i3{d3/i};
	cout<<"i3 : "<<i3<<"\n";

	int i4{ double(d3)/i};
	cout<<"i4 : "<<i4<<"\n";

	d2 = d3/i;
	cout<<"d2 : "<<d2<<"\n";
	
	i2 = d3/i;
	cout<<"i2 : "<<i2<<"\n";
	return 0;
}