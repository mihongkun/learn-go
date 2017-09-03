#include <stdio.h>
#include <time.h>


int main(){
	clock_t startTime, endTime;startTime = clock();  //获取开始时间
	for(int i = 1;;i++) {
		int j = 1;
		for(;j <= 20;j++) {
			if(i % j != 0) break;
		}
		if(j > 20) {
			printf("%d\n",i);
			break;
		}
	}
	endTime = clock();//获取结束时间
	printf( "%fs \n", (float)(endTime - startTime)/ CLOCKS_PER_SEC );
}
