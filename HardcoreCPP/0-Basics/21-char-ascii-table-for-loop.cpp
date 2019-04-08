#include "std_lib_facilities.h"

int main()
{
	int ch = 0;

	for (int i = 0; i<26; i++)
	{
		ch = int('a' + i);
		cout<<char(ch)<<" "<<ch<<"\n";
	}
	return 0;
}