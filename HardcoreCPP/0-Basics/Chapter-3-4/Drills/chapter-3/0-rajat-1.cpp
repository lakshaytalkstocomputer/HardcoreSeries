// testing ++ operator , pre and post 

#include "std_lib_facilities.h"

int main()
{
	int a{5};

	cout<<"a: "<<a<<"\n";
	cout<<"a: "<<a++<<"\n";
	cout<<"a: "<<++a<<"\n";
	cout<<"a: "<<a++<<" "<<++a<<"\n";
	return 0;
}