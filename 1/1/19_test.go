package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDLinearTwoEuclid(t *testing.T) {
	var d, x, y int

	d, x, y = GCDLinearTwoEuclid(12, 36)
	assert.Equal(t, 12, d)
	assert.Equal(t, d, 12*x+36*y)

	d, x, y = GCDLinearTwoEuclid(132, 36)
	assert.Equal(t, 12, d)
	assert.Equal(t, d, 132*x+36*y)

	d, x, y = GCDLinearTwoEuclid(5, 3)
	assert.Equal(t, 1, d)
	assert.Equal(t, d, 5*x+3*y)
}

type BinRat struct {
	a, div int
}

func (br *BinRat) DivTwo() {
	br.div *= 2
}

func Add(right, left BinRat) BinRat {
	if right.div >= left.div {
		return BinRat{
			a:   right.a + (right.div/left.div)*left.a,
			div: right.div,
		}
	} else {
		return BinRat{
			a:   left.a + (left.div/right.div)*right.a,
			div: left.div,
		}
	}
}

func Sub(right, left BinRat) BinRat {
	left.a = -left.a

	return Add(right, left)
}

func GCDLinearTwoEuclid(a, b int) (gcd, x, y int) {
	m, n := a, b

	p := BinRat{a: 1, div: 1}
	q := BinRat{a: 0, div: 1}
	s := BinRat{a: 0, div: 1}
	t := BinRat{a: 1, div: 1}

	d := 1

	// Inv. GCD(a;b) = d*GCD(m;n)
	// Inv. m = ma*p+mb*q, n = na*s+nb*t
	// Inv. d*m = a*p+b*q, d*n = a*s+b*t, so a=d*ma, b=d*mb etc.
	for !(m == 0 || n == 0) {
		if m%2 == 0 && n%2 == 0 {
			d *= 2

			m /= 2
			p.DivTwo()
			q.DivTwo()

			n /= 2
			s.DivTwo()
			t.DivTwo()
		} else if m%2 == 0 {
			m /= 2
			p.DivTwo()
			q.DivTwo()
		} else if n%2 == 0 {
			n /= 2
			s.DivTwo()
			t.DivTwo()
		} else if m >= n {
			m -= n

			p = Sub(p, s)
			q = Sub(q, t)
		} else {
			n -= m

			s = Sub(s, p)
			t = Sub(t, q)
		}
	}

	if m != 0 {
		a /= d
		b /= d

		rem := p.div
		if q.div > rem {
			rem = q.div
		}

		x = (rem / p.div) * p.a
		y = (rem / q.div) * q.a

		for rem > d {
			if x%2 == 0 && y%2 == 0 {
				x /= 2
				y /= 2
				rem /= 2
			} else {
				x += b
				y -= a
			}
		}

		return d * m, x, y
	} else {
		a /= d
		b /= d

		rem := s.div
		if t.div > rem {
			rem = t.div
		}

		x = (rem / s.div) * s.a
		y = (rem / t.div) * t.a

		for rem > d {
			if x%2 == 0 && y%2 == 0 {
				x /= 2
				y /= 2
				rem /= 2
			} else {
				x += b
				y -= a
			}
		}

		return d * n, x, y
	}

}
