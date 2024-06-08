package otp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOTPLength(t *testing.T) {
	expected := 64
	result := len(Generate())
	assert.Equal(t, expected, result, "Generate() should return an Generate with a length of 64 characters")
}

func TestGenerateOTPEmptyString(t *testing.T) {
	otp := Generate()
	assert.NotEmpty(t, otp, "Generate() should not return an empty string")
}

func TestGenerateOTPUniqueness(t *testing.T) {
	// Generate multiple OTPs and ensure they are unique
	otpSet := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		otp := Generate()
		assert.NotContains(t, otpSet, otp, "Generate() should return unique OTPs")
		otpSet[otp] = struct{}{}
	}
}

func TestGenerateRandomNumberInRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		num := generateRandomNumber()
		assert.True(t, num >= 100000 && num < 1000000, "generateRandomNumber() returned an out-of-range number: %d", num)
	}
}

func TestFormatAsSixDigitValidity(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "000000"},
		{12345, "012345"},
		{999999, "999999"},
		{1000000, "1000000"},
	}

	for _, test := range tests {
		result := formatAsSixDigit(test.input)
		assert.Equal(t, test.expected, result, "formatAsSixDigit(%d) returned %s, expected %s", test.input, result, test.expected)
	}
}
