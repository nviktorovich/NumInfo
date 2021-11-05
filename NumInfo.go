package NumInfo

// IsPrime func get int number (n) and returned bool value:
// true - if n is prime number, and false - if n is not prime number
func IsPrime(n int) bool {
	var flag bool
	if n < 2 {
		flag = false
	} else if n == 2 {
		flag = true
	} else {
		for i := 2; i < n/2; i++ {
			if n%2 == 0 {
				flag = false
				break
			}
		}
	}
	return flag
}

// IsOdd func get int number (n) and returned bool value: true - if n is odd, and false - if n is not odd
func IsOdd(n int) bool {
	return !(n&2 == 0)
}