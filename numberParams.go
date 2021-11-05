package NumberParams

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

func IsOdd(n int) bool {
	return n&2 != 0
}