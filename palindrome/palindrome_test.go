package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("a"))
	assert.True(t, IsPalindrome("aba"))
	assert.True(t, IsPalindrome("kodok"))
	assert.True(t, IsPalindrome(""))
	
	assert.False(t, IsPalindrome("ab"))
	assert.False(t, IsPalindrome("abab"))
	assert.False(t, IsPalindrome("kodcok"))
	assert.False(t, IsPalindrome("faqih"))
}