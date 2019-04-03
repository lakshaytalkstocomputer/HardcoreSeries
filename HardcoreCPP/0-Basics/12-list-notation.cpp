#include "std_lib_facilities.h"
int main()
{
	/* 
		list notation | universal and uniform intialization
	*/
	double x {2.7};
	cout<<"(double) x: "<<x<<"\n";

	int y {x}; // this gives warning
	cout<<"(int) y: "<<y<<"\n";

	int z {55};
	cout<<"(int) z: "<<z<<"\n";
	
	char c {z}; // this gives warning
	cout<<"(char) c: "<<c<<"\n";

    char ch {97}; 
	cout<<"(char) ch: "<<ch<<"\n";
    
    /* 
    	normal intilization 
    */

    double a = 2.7;
    cout<<"(double) a: "<<a<<"\n";

    int b = a; // does not gives warning
	cout<<"(int) :b "<<b<<"\n";

	int d = 55;
	cout<<"(int) d: "<<d<<"\n";
	
	char e = d;  // does not gives warning
	cout<<"(char) e: "<<e<<"\n";

    char g = 97;
	cout<<"(char) g: "<<g<<"\n";
    
}