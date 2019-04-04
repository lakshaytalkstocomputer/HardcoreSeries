#include "std_lib_facilities.h"

int main()
{
	 string operation;
	 float num1;
	 float num2;

	 cout<<"Enter expression in postfix : ";
	 cin>>operation>>num1>>num2;

	 if(operation == "+")
	 {
	 	cout<<"Answer: "<<num1+num2<<"\n";
	 }
	 else if(operation == "-")
	 {
	 	cout<<"Answer: "<<num1-num2<<"\n";
	 }
	 else if(operation == "*")
	 {
	 	cout<<"Answer: "<<num1*num2<<"\n";
	 }
	 else if(operation == "/")
	 {
	 	cout<<"Answer: "<<num1/num2<<"\n";
	 }
	 else
	 {
	 	cout<<"Invalid operation!";
	 }

	return 0;
}