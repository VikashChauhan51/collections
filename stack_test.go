package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test-specific setup function
func testSetup() func() {
	// Perform test-specific setup tasks
	println("Test-specific setup")
	// Return a teardown function to be called after the test
	return func() {
		println("Test-specific teardown")
	}
}

func Test_first(t *testing.T) {
	teardown := testSetup()
	defer teardown()
	fmt.Print("first")
	assert.Nil(t, nil)
}

func Test_second(t *testing.T) {
	teardown := testSetup()
	defer teardown()
	fmt.Print("Test_second")
	assert.Nil(t, nil)
}
