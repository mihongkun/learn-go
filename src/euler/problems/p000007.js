var i = 1;
var p = 2;
for (;i <= 10001;) {
	p++;	
	if(isPrime(p)) {
		i++;
	}
}

console.log(p)


function isPrime(n){
	for(let i = 2;i<Math.sqrt(n);i++) {
		if (n % i == 0) return false;
	}
	return true;
}
