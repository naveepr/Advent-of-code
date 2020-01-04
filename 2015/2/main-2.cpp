#include <iostream>
#include <fstream>
#include <string>
#include <cstdlib>
#include <set>

using namespace std;

int main(int argc, const char *argv[])
{
    ifstream ob("input.txt");
    string s;
    int l, b, h, left_index, right_index;
    unsigned long res =0;

    while(getline(ob, s)) {
       left_index =-1;
       right_index = s.find_first_of('x');
       l = stoi(s.substr(0, right_index));
        
       left_index = right_index+1;
       right_index = s.find_last_of('x');
       b = stoi(s.substr(left_index, right_index - left_index));

       left_index =right_index+1;
       h = stoi(s.substr(left_index, s.length()-left_index));
       
       res += (l<b) ? 2*l: 2*b;
       res += (l<b) ? (b<h ? 2*b:2*h):(l<h ? 2*l:2*h);

       res += l*b*h;
    }
    
    cout<<res<<endl;

    return res;
}
