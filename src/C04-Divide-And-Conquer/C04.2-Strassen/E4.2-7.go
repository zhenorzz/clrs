package C04_2_Strassen

func ComplexMultiply(x complex64, y complex64) complex64 {
	a := real(x)
	b := imag(x)
	c := real(y)
	d := imag(y)
	P1 := (a + b) * c
	P2 := (c + d) * b
	P3 := (a - b) * d
	return complex(P1-P2, P2+P3)
}
