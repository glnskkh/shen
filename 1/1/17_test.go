package shen

import "fmt"

func ExampleGCDDeikstra() {
	// 2*LCD(12;32) = 72 |=>
	fmt.Println(GCDDeikstra(12, 36))
	// Output: 72
}

func GCDDeikstra(a, b int) int {
	m := a
	n := b

	u := b
	v := a

	// Inv. GCD(a;b) = GCD(m;n)
	// Inv. m*u+n*v = 2*a*b is constant
	for !(m == 0 || n == 0) {
		if m >= n { // (m-n)*u+n*(u+v)=m*u-n*u+n*u+n*v=m*u+n*v
			m -= n
			v += u
		} else { // m*(u+v)+(n-m)*u=m*u+m*v+n*u-m*u=m*v+n*u
			n -= m
			u += v
		}
	}

	// 2nd inv. => either m*u or n*v = 2*a*b
	// But a*b = GCD(a;b)*LCD(a;b) |=>
	// either u or v = 2*LCD(a;b)

	if m != 0 {
		return u
	} else {
		return v
	}
}
