// Input three in number and print them in sorted 

#include "std_lib_facilities.h"

int main()
{
	string a, b, c;
	cout<<"Enter three strings: ";
	cin>>a>>b>>c;

	if(a<=b && a<=c)
	{
		cout<<a<<",";
		if(b<=c)
		{
			cout<<b<<","<<c<<"\n";
		}
		else
		{
			cout<<c<<","<<b<<"\n";
		}		
	}
	else if(b<=c)
	{
		cout<<b<<",";
		if(a<=c)
		{
			cout<<a<<","<<c<<"\n";
		}
		else
		{
			cout<<c<<","<<a<<"\n";	
		}		
	}
	else
	{
		cout<<c<<",";
		if(b<=a)
		{
			cout<<b<<","<<a<<"\n";
		}
		else
		{
			cout<<a<<","<<b<<"\n";	
		}		
	}
	return 0;
}