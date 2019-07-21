//#include <iostream>
#include <bits/stdc++.h>
using namespace std;

ifstream f("data.in");
ofstream g("data.out");
void countSort(int a[], int n)
{
	// finding the max
	int max = -1;
	for(int j =0; j<n;j++)
	{
		if (max<a[j])
			max = a[j];
	}
    

  // testing the input array
  //	cout<<"max : "<<max<<"\n";
//	for(int i = 0;i<n;i++)
	// {
	// 	cout<<a[i]<<" ";
	// }
	// cout<<"\n";


	// finding the frequency
	int count[max+1] = {0};   
    

	 for(int i= 0 ;i<n;i++)
	 {
	 	++count[a[i]];
	 }
      
     

	// count index by bubble adding the count array

	for(int i =1;i< max+1; i++)
	{
		count[i] += count[i-1];
	} 

	
    // 1. start picking elements from original array
    // 2. use it to find index from count array
    // 3. decrease by the value in count array
    // 4. put the element in new array at location from step 2
    
    int finalArray[n+1]={-99};
    for(int i= 0; i < n;i++ )
    {
    	int index = count[a[i]]--;
    	finalArray[index] = a[i];
    }
    

    for(int i = 1;i<n+1;i++)
    {
    	g<<finalArray[i]<<" ";
    }
    g<<"\n";
	
}

int main()
{
	int t;
	int n;
	int d;
	int data[100];
	f>>t;
	for (int i = 0; i<t; i++)
	{
		if(f>>n)
		{
			for (int j = 0; j<n;j++)
			{
				f>>d;
				data[j]=d;
			}
			countSort(data,n);
		}
	}
    
	
	return 0;
}

