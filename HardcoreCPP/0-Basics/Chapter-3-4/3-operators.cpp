// program to exercise some operators in float

#include "std_lib_facilities.h"

int main()
{
	cout<<"Please enter a floating point number: ";
	float n;
	cin>>n;
	cout<<"n == "<<n
		<<"\nn+1 == "<<n+1
		<<"\nthree times n == "<<3*n
		<<"\ntwice n == "<<2*n
		<<"\nn squared == "<<n*n
		<<"\nhalf of n == "<<n/2
		<<"\nsquare root of n == "<<sqrt(n)
	//	<<"\nn % 4 == "<<n%4 // error : invalid operand of type float and int
		<<"\n";
	return 0;
}