package rando

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

// TestNewRandom ensures creating a Random without a specific seed yields two distinct generators at different times.
func TestNewRandom(t *testing.T) {
	r1 := NewRandom()

	firstRun := make([]string, 0)
	for i := 0; i < 1000; i++ {
		firstRun = append(firstRun, r1.RandomString(5))
	}

	time.Sleep(1 * time.Second)
	r2 := NewRandom()
	secondRun := make([]string, 0)
	for i := 0; i < 1000; i++ {
		secondRun = append(secondRun, r2.RandomString(5))
	}

	assert.False(t, reflect.DeepEqual(firstRun, secondRun))
}

// TestNewSeededRandom ensures creating a Random with a specific seed yields the same random number generator.
func TestNewSeededRandom(t *testing.T) {
	r1 := NewSeededRandom(42)

	firstRun := make([]string, 0)
	for i := 0; i < 1000; i++ {
		firstRun = append(firstRun, r1.RandomString(5))
	}

	r2 := NewSeededRandom(42)
	secondRun := make([]string, 0)
	for i := 0; i < 1000; i++ {
		secondRun = append(secondRun, r2.RandomString(5))
	}

	assert.EqualValues(t, firstRun, secondRun)
}

