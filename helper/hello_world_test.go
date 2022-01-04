package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Faqih")
	assert.Equal(t, "Hello Faqih", result, "Result must be 'Hello Faqih'")
	fmt.Println("TestHelloWorld with Assert Done")
}


func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Faqih")
	if result != "Hello Faqih" {
		panic("Result is not Hello Faqih")
	}
}

