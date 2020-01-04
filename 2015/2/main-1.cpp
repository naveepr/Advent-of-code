#include <iostream>
#include <fstream>
#include <string>
#include <cstdlib>

using namespace std;

int main(int argc, const char *argv[])
{
    ifstream ob("input.txt");
    string s;
    int l, b, h, left_index, right_index, area;
    unsigned long res =0;

    while(getline(ob, s)) {
       left_index =-1;
       right_index = s.find_first_of('x');
       l = stoi(s.substr(0, right_index));
    
       //cout <<l<<endl;
       left_index = right_index+1;
       right_index = s.find_last_of('x');
       b = stoi(s.substr(left_index, right_index - left_index));
      
       //cout << b<<endl;
       left_index =right_index+1;
       h = stoi(s.substr(left_index, s.length()-left_index));

       //cout<< h <<endl;
       area = min(min(l*b, b*h), l*h);
       res += (2*(l*b + b*h + h*l) + area);
    }
    
    cout<<res<<endl;

    return res;
}
