package main

import (
        "fmt"
)


func main(){
        var temp,result uint64
        var f,s uint64 = 1,2
        result = uint64(2)
        fmt.Println(1,"\t\t\t奇\t\t\t",1) 
        fmt.Println(2,"\t\t\t偶\t\t\t",2)
        for i:=1;;i++ {
                temp = f + s            
                f = s
                s = temp
                if temp % 2 ==0 {
                        result += temp
			if result > 4000000 {
                        break
             		}
                        fmt.Println(i + 2,"\t\t\t偶\t\t\t",temp)
                } else {
                        fmt.Println(i + 2,"\t\t\t奇\t\t\t",temp)
                }
                        
        }
        fmt.Println("Result:",result)
}
