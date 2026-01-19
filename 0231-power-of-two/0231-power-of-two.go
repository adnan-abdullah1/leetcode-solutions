func IterisPowerOfTwo(n int) bool {

	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}

	return n == 1
}

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
    
	if n <= 0 || n%2 != 0 {
		return false
	}

	return isPowerOfTwo(n / 2)

}