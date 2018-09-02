// run the following program here
// https://www.onlinegdb.com/online_c_compiler
#include<stdlib.h>
#include<stdio.h>

int main (){
    int c9s = rand()-rand();
    int input[]={-957358430,-957358440,-957358414,-957358487,-957358425,-957358432,-957358417,-957358417,-957358408,-957358487,-957358431,-957358424,-957358415,-957358413,-957358425,-957358429,-957358432,-957358408,-957358487};
    int i = 0;
    for (; i < sizeof(input)/sizeof(int); i++) {
        printf("%c",input[i]+c9s);
    }
}
