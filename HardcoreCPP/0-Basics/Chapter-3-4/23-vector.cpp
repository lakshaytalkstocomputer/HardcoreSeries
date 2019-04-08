#include "std_lib_facilities.h"

int main()
{
	vector <int> data;

	for(int value; cin>>value;)  // we use cin>>value in condition of for loop
								 // so that if the entered data is not integer
								 // it will stop taking data and loop will exit
	{
		data.push_back(value);
	}

	// We could have used while loop
	// But we have to limit the scope of temp variable
	// So in for loop the temp variable wont be available
	// after for loop

	int sum = 0;

	for(int x: data)
		sum += x;

	cout<<"Sum of entered integers is: "<<sum<<"\n";
}