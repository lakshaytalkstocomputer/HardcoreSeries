// take two number and print ratio of the number in the form "a:b"


#include<iostream>
using namespace std;

int calcualte_gcd(int a, int b);

int main()
{
	int a;
	int b;
	int x, y;
	cout<<"Enter two numbers: ";
	cin>>a>>b;
	int gcd;

	gcd = calcualte_gcd(a,b);
    
    x = a/gcd;
    y = b/gcd;

	cout<<"Ratio = "<<x<<":"<<y<<"\n";

	return 0;
}

int calcualte_gcd(int a, int b)
{
	if (b == 0)
	{
		return a;
	}
	return calcualte_gcd(b, a%b);
}