package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 1, Factorial(0))
	assert.Equal(t, 1, Factorial(1))
	assert.Equal(t, 2, Factorial(2))
	assert.Equal(t, 6, Factorial(3))
	assert.Equal(t, 24, Factorial(4))
	assert.Equal(t, 120, Factorial(5))
}
func TestRecursiveFactorial(t *testing.T) {
	assert.Equal(t, 1, RecursiveFactorial(0))
	assert.Equal(t, 1, RecursiveFactorial(1))
	assert.Equal(t, 2, RecursiveFactorial(2))
	assert.Equal(t, 6, RecursiveFactorial(3))
	assert.Equal(t, 24, RecursiveFactorial(4))
	assert.Equal(t, 120, RecursiveFactorial(5))
}
func TestTailRecursiveFactorial(t *testing.T) {
	assert.Equal(t, 1, TailRecursiveFactorial(1,0))
	assert.Equal(t, 1, TailRecursiveFactorial(1,1))
	assert.Equal(t, 2, TailRecursiveFactorial(1,2))
	assert.Equal(t, 6, TailRecursiveFactorial(1,3))
	assert.Equal(t, 24, TailRecursiveFactorial(1,4))
	assert.Equal(t, 120, TailRecursiveFactorial(1,5))
}