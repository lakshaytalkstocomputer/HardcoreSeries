#include "std_lib_facilities.h"

int main()
{
	vector<double> temps;

	for(double temp; cin>>temp;)
		temps.push_back(temp);

	// compute mean temperature
	double sum = 0;
	for(double data: temps)
		sum+=data;

	cout<<"Average temp: "<<sum/temps.size()<<"\n";

	// compute median of temperature
	  // sort first

	sort(temps);
	cout<<"Median temperature: "<<temps[temps.size()/2]<<"\n";
	
}