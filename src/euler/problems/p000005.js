var t1 = new Date().getTime();
var t2 = t1;
t1 = process.uptime()*1000;

for(var i=1;;i++) {
	for(var j=1;j<=20;j++){
		if (i % j != 0) break;
	}
	if(j > 20) {
		console.log(i);
		break;
	}
}

t2 = process.uptime()*1000;
var interval = (t2 - t1);
console.log(interval / 1000 + "s")
