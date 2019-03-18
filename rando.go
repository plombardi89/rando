package rando

import (
	"math/rand"
	"time"
)

// Random provides convenient utility funcs for generating random data
type Random struct {
	rand *rand.Rand
}

// NewRandom returns a new Random seeded with the current UTC time converted to nanoseconds.
func NewRandom() *Random {
	return NewSeededRandom(time.Now().UTC().UnixNano())
}

// NewSeededRandom returns a new Random seeded with the provided value.
func NewSeededRandom(seed int64) *Random {
	return &Random{rand: rand.New(rand.NewSource(seed))}
}

// RandomString returns a random string of the given length from a standard lowercase-only alphanumeric alphabet.
func (r *Random) RandomString(length int) string {
	return r.RandomStringUsingCustomAlphabet(length, []rune("0123456789abcdefghijklmnopqrstuvwxyz"))
}

// RandomStringUsingCustomAlphabet return a random string of the given length from a custom alphabet.
func (r *Random) RandomStringUsingCustomAlphabet(length int, alphabet []rune) string {
	res := make([]rune, length)
	for i := range res {
		res[i] = alphabet[r.rand.Intn(len(alphabet))]
	}

	return string(res)
}

// RandomSelectionFromStringSlice returns a random string from a given slice of strings.
func (r *Random) RandomSelectionFromStringSlice(values []string) string {
	if len(values) == 0 {
		return ""
	}

	return values[rand.Intn(len(values))]
}

// RandomBool returns a random boolean value.
func (r *Random) RandomBool() bool {
	return rand.Intn(2) == 0
}
