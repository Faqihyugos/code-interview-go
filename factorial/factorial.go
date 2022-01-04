package factorial

func Factorial(value int) int {
	if value <= 0 {
		return 1
	}

	result := 1
	for i := value; i >= 1; i-- {
		result *= i
	}
	return result
}

func RecursiveFactorial(value int) int {
	if value <= 0 {
		return 1
	} else {
		// 5 * factorial(4)
		// 5 * 4 * factorial(3)
		// 5 * 4 * 3 * factorial(2)
		// 5 * 4 * 3 * 2 * factorial(1)
		// 5 * 4 * 3 * 2 * 1 * factorial(0)
		// 5 * 4 * 3 * 2 * 1 * 1
		return value * RecursiveFactorial(value-1)
	}
}

// syarat tail recursive : jgn ada data yg dikalikan function recurvise
func TailRecursiveFactorial(total, value int) int {
	// factorial(1, 5) => factorial(5,4) => factorial(20,3)
	if value <= 0 {
		return total
	} else {
		return TailRecursiveFactorial(total*value, value-1)
	}
}