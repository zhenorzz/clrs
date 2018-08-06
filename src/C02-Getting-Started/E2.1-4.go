package C02_Getting_Started

func Add(a []int, b []int) []int {
	n := len(a)
	c := make([]int, n+1)
	carry := 0
	for i := n - 1; i >= 0; i-- {
		c[i+1] = (a[i] + b[i] + carry) % 2
		if a[i]+b[i]+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}
	c[0] = carry
	return c
}
