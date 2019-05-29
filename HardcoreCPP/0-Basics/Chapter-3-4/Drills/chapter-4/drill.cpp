#include "std_lib_facilities.h"


int main()
{
int a{0} ,b{0};
int max{0};
	while(cin>>a && cin>>b)
	{
		// print the number as it is
		cout<<a<<" "<<b<<"\n";

		// print " smaller number is : " followed by smaller number 
		// print " larger number is : " followed by larger number
		if( a!= b)
		{
			cout<<"smaller number is: "<<((a<b)?a:b);
			cout<<"\nlarger number is :"<<((a>b)?a:b)<<"\n";
		}
		else
		{
			cout<<"the numbers are equal.\n";
		}
		//
	
	}	
}