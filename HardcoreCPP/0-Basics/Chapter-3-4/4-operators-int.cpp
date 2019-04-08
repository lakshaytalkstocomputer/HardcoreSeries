// 3- operators program to wrok on int

#include "std_lib_facilities.h"

int main()
{
	cout<<"Please enter integer: ";
	int n;
	cin>>n;

	cout<<"n == "<<n
		<<"\nn+1 == "<<n+1
		<<"\nthree time n  == "<<3*n
		<<"\nn squared == "<<n*n // return in since both int
	    <<"\nhalf of n == "<<n/2 // return int since both int
	    <<"\nsquare root of n == "<<sqrt(n) // return float
	    <<"\nn % 4 == "<<n%4 // return int
	    <<"\n";

	return 0;
}