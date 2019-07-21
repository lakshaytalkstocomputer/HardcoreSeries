#include "std_lib_facilities.h"

using namespace std;

void compare_vector(vector<string> v1,vector<string> v2)
{
	int i = 0;
	int j = 0;
    
    bool found = false;
	int v1_len = v1.size();
	int v2_len = v2.size();
    
    int count_loop =0;
	while(i != v1_len && j != v2_len)
	{
		++count_loop;
		if(v1[i] > v2[j])
		{
			++j;
		}
		else if(v1[i] < v2[j])
		{
			++i;
		}
		else if(v1[i] == v2[j])
		{
			cout<<"Found. Your culprit: "<<v1[i]<<": "<<v2[j]<<"\n";
			found = true;
			break;
		}
	}
	cout<<"loop ran for: "<<count_loop<<"\n";
	if(!found)
	{
		cout<<"Sorry try again!\n";	
	}
	
}

int main()
{
	vector<string> v1= {"Hello","Hi","Bye"};
	vector<string> v2= {"ok","no","tata","Hello","Zye"};
	
	/*
		1. Sort the Array 
		2. Implement merge like routine to find same
	*/
	sort(v1);
	sort(v2);
	compare_vector(v1,v2);


	
	return 0;
}