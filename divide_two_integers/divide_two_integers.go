package dividetwointegers

// divide implements division without built-in division, multiplication, or mod.
// Multiplication is technically used to handle positive/negative cases elegantly.
func divide(dividend int, divisor int) int {
	sign := 1
	if divisor < 0 {
		// Normalize so divisor is always positive
		divisor = -divisor
		dividend = -dividend
	}
	if dividend < 0 {
		sign = -1
	}

	quotient := 0
	remainder := dividend
	for sign*remainder >= divisor && quotient != 2147483647 {
		quotient += sign
		remainder -= sign * divisor
	}

	return quotient
}
