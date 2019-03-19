package rando

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"regexp"
	"testing"
	"time"
)

var seed int64 = 42

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
	r1 := NewSeededRandom(seed)

	firstRun := make([]string, 0)
	for i := 0; i < 1000; i++ {
		firstRun = append(firstRun, r1.RandomString(5))
	}

	r2 := NewSeededRandom(seed)
	secondRun := make([]string, 0)
	for i := 0; i < 1000; i++ {
		secondRun = append(secondRun, r2.RandomString(5))
	}

	assert.EqualValues(t, firstRun, secondRun)
}

func TestRandom_RandomString(t *testing.T) {
	var table = []struct {
		length  int
		matches *regexp.Regexp
	}{
		{0, regexp.MustCompile("")},
		{100, regexp.MustCompile("[a-z0-9]{100}")},
	}

	r := NewSeededRandom(seed)
	for _, tt := range table {
		v := r.RandomString(tt.length)
		assert.Regexp(t, tt.matches, v)
	}
}

func TestRandom_SampleStringSlice(t *testing.T) {
	var table = []struct {
		values     []string
		sampleSize int
		shouldErr  bool
	}{

		{[]string{"a", "b", "c", "d", "e", "e", "f", "e"}, 3, false},
		{[]string{"a", "b", "c"}, 0, false},
		{[]string{"a", "b", "c"}, -1, true},
		{[]string{"a", "b", "c"}, 10, true},
	}

	r := NewSeededRandom(seed)
	for _, tt := range table {
		sampled, err := r.SampleStringSlice(tt.values, tt.sampleSize)

		if tt.shouldErr {
			assert.NotNil(t, err)
			return
		}

		assert.Len(t, sampled, tt.sampleSize)
	}
}
