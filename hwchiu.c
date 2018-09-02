// run the following program here
// https://www.onlinegdb.com/online_c_compiler
#include<stdlib.h>
#include<stdio.h>

int w(int n) {
    return n? n%10+(w(n/10)<<1):0;
}

int main (){
    int c9s = rand()-rand();
    int input[]={958358508,957469498,958368508,957359507,958359497,958358498,958368497,958368497,958369498,957359507,958358507,958359498,958368507,958368597,958359497,958358597,958358498,958369498,957359507};
    int i = 0;
    for (; i < sizeof(input)/sizeof(int); i++) {
        printf("%c",w((input[i]-c9s)));
    }
}
