// count the number of words and print the repeated words

#include "std_lib_facilities.h"

int main()
{
	 int word_count = 0;

	 string previous = " ";
	 string current;

	 while(cin>>current)
	 {
	 	++word_count;
	 	if(previous == current)
	 	{
	 		cout<<"Repeated words: "<<current<<'\n';
	 	}
	 	previous = current;
	 }
	 cout<<"Total word count: "<<word_count<<'\n';
}