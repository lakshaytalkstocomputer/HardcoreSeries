#include "std_lib_facilities.h"

int main()
{
	int i = 0;
	int ch = 0;
	while(i<26)
	{
		ch = int('a' + i);

		cout<<char(ch)<<" "<<ch<<"\n";
		++i;
	}
	return 0;
}