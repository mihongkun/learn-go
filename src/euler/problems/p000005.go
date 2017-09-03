package main

import(
	"fmt"
	"container/list"
)

func main(){
	var list list.List = filte(20)
	var primeList list.List = getPrimeList(20)
	for e:=list.Front();e != nil;e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()
}





func filte(n int) list.List{
	var result list.List
	for i:=2;i <= n;i ++ {
		j :=n
		for ;j>= n / 2;j-- {
				   if j % i == 0 {break}
		}
		if i > n / 2  || (j == n / 2 && j % i != 0) {
			result.PushBack(i)
		}
	}
	return result
}	

func isPrime(n int) bool{
	for i:=2;i <int( math.Sqrt(float64(n)));i++ {
		if n % i == 0 {
			return false
		}
	}
	
	return true
}

func getPrimeList(n int) list.List {
	m = n / 2 
	var result list.List
	for i:=2;i<n;i++{
		if isPrime(i) {
			result.PushBack(i)
		}
	}
	return result
}

func getBothPrime(list list.List, pl list.List) (primes list.List) {
	for e:=list.Front();e != nil;e := list.Next() {
		if list.Next() != nil  {
			for pe:=pl.Front();pe != nil;pe:= pl.Next() {
				if e.Value % pe.Value && list.Next().Value % pe.Value {
					primes.PushBack(pe.Value)
					list.PushBack(e.Value)
					list.Remove(e)
					list.Remeve(list.Next())
						
				}
			}
		}
	}		

}

