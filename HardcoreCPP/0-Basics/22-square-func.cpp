#include "std_lib_facilities.h"

int squared(int x)
{
	int ans = 0;
	for(int i = 0; i< x; ++i)
	{
		ans += x;
	}

	return ans;
}

int main()
{	
	cout<<"Enter a number you want to square: ";
	int num;
	cin>>num;

	cout<<"Square of "<<num<<" is "<<squared(num)<<'\n';
	return 0;
}