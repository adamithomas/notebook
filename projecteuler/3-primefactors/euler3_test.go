package primefactorization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	primes := [...]int{2, 3, 5, 7, 11, 13, 149}
	for _, n := range primes {
		assert.True(t, IsPrime(n), "%d should be prime", n)
	}

	composites := [...]int{4, 6, 8, 9, 10, 200, 1024}
	for _, n := range composites {
		assert.False(t, IsPrime(n), "%d should not be prime", n)
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	// Check that we get the right answer for Project Euler's example
	assert.Equal(t, LargestPrimeFactor(13195), 29,
		"the largest prime factor of 13195 is 29")

	// Check our calculated answer to the problem
	assert.Equal(t, LargestPrimeFactor(600851475143), 6857,
		"the largest prime factor of 600851475143 is 6857")
}
