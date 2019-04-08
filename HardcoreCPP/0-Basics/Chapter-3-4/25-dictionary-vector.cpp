
#include "std_lib_facilities.h"

int main()
{
	vector <string> words;

	for(string temp; cin>>temp;) // we use cin>>temp in condition of for loop
								 // so that if the entered data is not integer
								 // it will stop taking data and loop will exit
	{
		words.push_back(temp);
	}

	// We could have used while loop
	// But we have to limit the scope of temp variable
	// So in for loop the temp variable wont be available
	// after for loop


	cout<<"Number of words: "<<words.size()<<"\n";

	sort(words);
	for(int i = 0; i< words.size(); i++)
	{
		if(i == 0 || words[i-1] != words[i])
			cout<<words[i]<<'\n';
	}

	return 0;
}