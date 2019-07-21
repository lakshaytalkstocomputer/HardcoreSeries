// program to check conditional operator in strings

#include "std_lib_facilities.h"

int main()
{
	cout<<"Please enter two name: ";
	string first;
	string second;

	// read two strings
	cin>>first>>second;

	if(first == second)
		cout<<"That's the same name twice\n";
	if(first < second) // comapres every word 
		cout<<first<<" is alphabetically before "<<second<<"\n";
	if(first > second) // comapres every word 
		cout<<first<<" is alphabetically after "<<second<<"\n";

	return 0;
}