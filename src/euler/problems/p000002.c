#include <stdio.h>

int fuc4()  
{//Fn=4*F(n-3)+F(n-6)  
    int sum=0;  
    int fn;  
    int fn_3=8;  
    int fn_6=2;  
    while(fn_6<4000000)  
    {  
        sum+=fn_6;  
        fn=4*fn_3+fn_6;       
        fn_6=fn_3;  
        fn_3=fn;  
    }  
    return  sum;  
} 

int main(){
	printf("result=%d",fuc4());
	return 0;
}
