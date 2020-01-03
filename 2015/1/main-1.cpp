#include <iostream>
#include <fstream>

using namespace std;

int main(int argc, const char *argv[])
{
    string fileName;
    char c;
    long int count=0;

    cout<<"Enter the name of the file"<<endl;
    cin >> fileName;
    
    ifstream is(fileName);

    while(is.get(c)) {
        switch(c) {
            case '(':
            case '{':
            case '[':
                        count++;
                        break;
            case ')':
            case '}':
            case ']':
                        count--;
                        break;
            default:
                        break;
        }
    }

    is.close();
    cout<<count<<endl;
    return 0;
}
