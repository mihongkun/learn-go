package main


import (
	"fmt"
	"math"
)


func main(){
//	fmt.Println("99999是回文吗?",isPalindromic(99999))
	fmt.Println("最终结果: ",findByByte(3))	

}

func findByByte(n int) int {
	var result int 
	n2 := n * 2
	x := int(math.Pow(10,float64(n2))-1)
	basex :=int(math.Pow(10,float64(n)))
	basei :=int(math.Pow(10,float64(n-1)))
	for i:=x;;i-- {
		if isPalindromic(i) {
			fmt.Println("回文:",i)
			for j:=basex;j > basei;j-- {
				if i % j == 0  && i / j < basex && i / j > basei {
					result = i
					fmt.Println("三位数i:\t\t\t",i / j,"\t\t\t三位数j\t\t\t",j)
					return result
				}  
			}	
		}	
	}
	return -1
}

func isPalindromic(x int) bool{
	var xstr string = fmt.Sprintf("%d",x)
	var xstrLen int = len(xstr)
	for i:=0;i <= xstrLen / 2;i++ {
		fmt.Println(xstr[xstrLen - i - 1],xstr[i])
		if xstr[xstrLen - i - 1] != xstr[i] {
			return false
		}
	}
	return true



}





