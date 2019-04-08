// convert yen ('y'), kroner ('k'), pounds ('p') into dollar ('$')

#include "std_lib_facilities.h"

int main()
{
	constexpr double yen_to_dollar = 0.009 ;
	constexpr double kroner_to_dollar = 0.12;
	constexpr double pound_to_dollar = 1.3;

	double amount = 0; // stores aamount entered by user
	char unit = 0;     // stores the unit entered by user

	cout<<"Enter the amount follow by the currency( y, k, p for yen , kroner , pound resp. ): ";
	cin>>amount>>unit;

	if(unit == 'y')
	{
		cout<<amount<<" y ="<<yen_to_dollar*amount<<" $\n";
	}
	else if(unit == 'k')
	{
		cout<<amount<<" k ="<<kroner_to_dollar*amount<<" $\n";
	}
	else if(unit == 'p')
	{
		cout<<amount<<" p ="<<pound_to_dollar*amount<<" $\n";
	}
	else
	{
		cout<<"Sorry I don't know any currency with "<<unit<<" unit.\n";
	}
	return 0;
}