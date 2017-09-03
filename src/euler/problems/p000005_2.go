package main

import (
	"fmt"
)

func main() {
	const N int = 20
	var eles [N / 2]int
	var result int = 1
	var k int = 0
	fmt.Println("len(eles)",len(eles))	
	for i:=N;i>N/2;i--{
		result *= i
		eles[k] =  i
		fmt.Println("k:",k)
		k++
	}
	temp := eles
	for i:=0;i<len(eles) - 1;i++ {
		for j:=i+1;j<len(eles);j++{
			for k:=N/2;k >= 2; k--{
			if eles[i] % k ==0 && eles[j] % k == 0{
					eles[i] /= k
					eles[j] /= k
					result /= k
			
				} 
			}
		}
	}
	
	for i:=0;i<len(eles) - 1;i++ {
                for j:=i+1;j<len(eles);j++{
                        for k:=N/2;k >= 2; k--{
                                if eles[i] % k ==0 && eles[j] % k == 0{
                                        eles[i] /= k
                                        eles[j] /= k
                                        result /= k    
                                }
                        }
                }
        }
	for i:=0;i<len(eles) - 1;i++ {
                for j:=i+1;j<len(eles);j++{
                        for k:=N/2;k >= 2; k--{
                                if eles[i] % k ==0 && eles[j] % k == 0{
                                        eles[i] /= k
                                        eles[j] /= k
                                        result /= k    
                                }
                        }
                }
        }
	fmt.Println(eles)	
	
	fmt.Println(result)	
	for i:=0;i<len(temp);i++{
		fmt.Println(temp[i],result/temp[i])
	}		


}

