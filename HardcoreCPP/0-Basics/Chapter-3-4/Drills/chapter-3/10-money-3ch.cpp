#include "std_lib_facilities.h"

int main()
{
	int pennies, nickels, dimes, quarters, half_dollar, full_dollars;
	
	cout<<"How many pennies you have? ";	
	cin>>pennies;
	cout<<"How many nickels you have? ";	
	cin>>nickels;
	cout<<"How many dimes you have? ";	
	cin>>dimes;
	cout<<"How many quarters you have? ";	
	cin>>quarters;
	cout<<"How many half dollar you have? ";	
	cin>>half_dollar;
	cout<<"How many one dollar you have? ";	
	cin>>full_dollars;

	int total_cent = 0;
	total_cent = pennies*1 + nickels*5 + dimes*10 + quarters*25 + half_dollar*50 + full_dollars*100;

	float dollar{total_cent};
	dollar = dollar/100;
	
	if(pennies>0)
	{
	cout<<"You have "<<pennies;
	if(pennies>1)
		cout<<" pennies\n";
	else
		cout<<" penny.\n";
	}

	if(nickels > 0)
	{
	cout<<"You have "<<nickels;
	if(nickels>1)
		cout<<" nickels\n";
	else
		cout<<" nickels.\n";
	}

	if(dimes>0)
	{
	cout<<"You have "<<dimes;
	if(dimes>1)
		cout<<" dimes\n";
	else
		cout<<" dime.\n";
	}

	if(quarters > 0)
	{
	cout<<"You have "<<quarters;
	if(quarters>1)
		cout<<" quarters\n";
	else
		cout<<" quarter.\n";
 	}

 	if(half_dollar > 0)
 	{   
    cout<<"You have "<<half_dollar;
    if(half_dollar>1)
		cout<<" half dollars\n";
	else
		cout<<" half dollar.\n";
	}

	if(full_dollars >0)
	{
		cout<<"You have "<<full_dollars;
    if(full_dollars>1)
		cout<<" full dollars\n";
	else
		cout<<" full dollar\n";
	
	}
    
    cout<<"You have total "<<total_cent;
    if(total_cent>1)
		cout<<" cents\n";
	else
		cout<<" cent.\n";

    cout<<"You have total $"<<dollar<<"\n";

	return 0;
}	