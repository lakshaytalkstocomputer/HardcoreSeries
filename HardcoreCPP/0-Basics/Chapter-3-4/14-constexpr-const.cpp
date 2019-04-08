#include "std_lib_facilities.h"

constexpr int maxi = 100;
using namespace std;
void use (int n)
{
	constexpr int c1 = maxi +7;
	//constexpr int c2 = n +7; // error: we don't the value of c2

	const int c2 = n +7;
	cout<<"c2: "<<c2<<"\n";

}

constexpr int sumi(int n)
{
	return n*n;
}

int main()
{
	use(20);
	cout<<"sumi : "<<sumi(1500)<<"\n";
	return 0;

}
