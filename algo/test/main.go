package test

func recursion(n int) int {
	if n == 1 {
		return n
	}

	res := recursion(n - 1)
	return n + res
}

func tailRecursion(n, res int) int {
	if n == 0 {
		return res
	}
	res += n
	n -= 1
	return tailRecursion(n, res)
}
