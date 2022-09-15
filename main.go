package main

import "fmt"

func main() {
	 Discount(120000)
}

func Discount(x int) {
	b:=0
	c:=0

	
	if x>=50000&& x<=70000 {
		b=(x*25)/100
		c=25
		if b>=20000 {
			b=20000
		}
	}else if x>=70000{
		b=(x*50)/100
		c=50
		if b>=50000 {
			b=45000
		}
	}else{
		b+=0
		c=0
	}
	d :=x-b

	fmt.Println("Total belanja : Rp.",x ,"-Diskon :",c,"%, Potongan: Rp.",b,"-Total pembayaran: Rp.",d)
}

