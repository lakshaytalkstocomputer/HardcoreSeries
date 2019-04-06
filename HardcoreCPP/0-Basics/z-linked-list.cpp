// creating a linked list
#include "std_lib_facilities.h"

// creating a skelton of node
class Node
{
private:
	int info;
	Node* next;
public:
	Node()
	{
		info = 0;
		this->next = NULL;
	}
	Node(int data)
	{
		info = data;
		this->next = NULL;
	}
	Node(int data,Node n)
	{
		info = data;
		this->next = &n;
	}

	void set_data(int x)
	{
		info = x; 
	}

	void set_next_info(Node n)
	{
		this->next = &n;
	}

	int get_data()
	{
		return info;
	}

	Node* get_next_node()
	{
		return next;
	}

};


int main()
{
	Node a = new Node();
	cout<<"a: "<<a.get_data()<<"\n";
	cout<<"a's Next Node: "<<a.get_next_node()<<"\n";
	return 0;
}