#include "std_lib_facilities.h"


int main()
{

/*double a{0} ,b{0};
double max{0};

double larger{0}, smaller{0};
double diff{0};
	while(cin>>a && cin>>b)
	{
		// print the number as it is
		cout<<a<<" "<<b<<"\n";
		larger = ((a>b)?a:b);
		smaller = ((a<b)?a:b);
        
        diff = larger - smaller;
        //double delta = 1.0/100;
		
		// print " smaller number is : " followed by smaller number 
		// print " larger number is : " followed by larger number
		if( a!= b)
		{
			cout<<"smaller number is: "<<smaller<<"\n";
			cout<<"larger number is :"<<larger<<"\n";
			
			if(diff <= (1.0/100))
			{
				cout<<"the numbers are almost equal.\n";
			} 

		}
		
		else
		{
			cout<<"the numbers are equal.\n";
		}
		//
	
	}	
	*/

	double input{0};
	double smaller{99999};
	double larger{-99999};
	
	double sum{0};
	int count{0};
	vector<double> data;
    string unit;
	cout<<"Enter value along with unit (cm, m, in, ft)\n";
	while(cin>>input && cin>>unit)
	{
		/*if(input<smaller)
		{
			smaller = input;
		}
		if(input>larger)
		{
			larger = input;
		}

		cout<<smaller<<" the smallest so far\n";
		cout<<larger<<" the largest so far\n";
		*/
		if(unit== "cm" || unit== "m" || unit== "in" || unit== "ft" )
		{
			//cout<<input<<" "<<unit<<"\n";
			
			count += 1;
			double value{0};
			if(unit == "cm")
			{
				value = input/100;
				 
			}
			else if(unit == "m")
			{
				value += input;
				
			}
			else if(unit == "in")
			{
				value = (input*2.54)/100; 
				
			}
			else if(unit == "ft")
			{
				value = ((input*12)*2.54)/100; 
			}

			if(value<smaller)
			{
				smaller = value;
			}
			if(value>larger)
			{
				larger = value;
			}

			data.push_back(value);
			sum += value;
		}
		else
		{
			cout<<"Rejected.\n";

		}

	}

	cout<<"Largest : "<<larger<<" m"<<"\n";
	cout<<"Smallest : "<<smaller<<" m"<<"\n";
	cout<<"Count : "<<count<<"\n";

	sort(data);

	for(double d : data)
	{
		cout<<d<<" ";
	}
}