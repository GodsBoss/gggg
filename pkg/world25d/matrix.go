package world25d

// matrix represents the following matrix:
//    a b c
// (  d e f )
//    0 0 1
type matrix struct {
	a float64
	b float64
	c float64
	d float64
	e float64
	f float64
}

func (m matrix) apply(x, y float64) (float64, float64) {
	return x*m.a + y*m.b + m.c, x*m.d + y*m.e + m.f
}

// identityMatrix is the matrix which leaves vectors unchanged if applied.
var identityMatrix = matrix{
	a: 1,
	e: 1,
}

func translationMatrix(tx, ty float64) matrix {
	return matrix{
		a: 1,
		c: tx,
		e: 1,
		f: ty,
	}
}

func multiplyMatrices(m1, m2 matrix) matrix {
	return matrix{
		a: m1.a*m2.a + m1.b*m2.d,
		b: m1.a*m2.b + m1.b*m2.e,
		c: m1.a*m2.c + m1.b*m2.f + m1.c,
		d: m1.d*m2.a + m1.e*m2.d,
		e: m1.d*m2.b + m1.e*m2.e,
		f: m1.d*m2.c + m1.e*m2.f + m1.f,
	}
}
