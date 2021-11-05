package NumInfo

import "errors"

// IsPrime func get int number (n) and returned bool value: true - if n is prime number, and false - if n is not prime number
func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else {
		for i := 2; i < n/2; i++ {
			if n%2 == 0 {
				return false
			}
		}
	}
	return true
}

// IsOdd func get int number (n) and returned bool value: true - if n is odd, and false - if n is not odd
func IsOdd(n int) bool {
	return !(n%2 == 0)
}

// IsBigger func get two int numbers, and returned true if n1 > n2, else - false
func IsBigger(n1, n2 int) bool {
	return n1 > n2
}

// GetMaxDivider func get int type number, and returned bigger divider
func GetMaxDivider(n int) (int, error) {
	if n == 0 {
		err := errors.New("zero number has not divider")
		return 0, err
	} else if n > 1 {
		for i := n - 1; i > 0; i-- {
			if n%i == 0 {
				return i, nil
			}
		}
	} else {
		for i := n + 1; i < 0; i++ {
			if n%i == 0 {
				return i, nil
			}
		}
	}
	return n, nil
}
