package logic_test

import (
	"fmt"
	"github.com/og11423074s/go-test-basic/internal/logic"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Running tests")
	exitCode := m.Run()
	fmt.Println("Exiting with code", exitCode)
	os.Exit(exitCode)
}

func TestSum(t *testing.T) {
	got := logic.Sum(1, 2, 3, 4, 5)
	expected := 15

	assert.Equal(t, expected, got)

	got = logic.Sum(4, 5)
	expected = 9

	assert.Equal(t, expected, got)

}

func TestCheckEmail(t *testing.T) {
	got := logic.CheckEmail("test@test.com")
	expected := true
	assert.True(t, got)

	emailInvalid := "testtest.com.ar"
	got = logic.CheckEmail(emailInvalid)
	expected = false
	assert.False(t, expected, got)
}

func TestSumMatrix(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Sum of 1, 2 and 3",
			input:    []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Sum of 1, 2, 3 and 4",
			input:    []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "Sum of 1, 2, 3, 4 and 5",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Sum of empty",
			input:    []int{},
			expected: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := logic.Sum(c.input...)
			assert.Equal(t, c.expected, got)
		})

	}
}

func TestCheckEmailMatrix(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid email",
			input:    "test@gmail.com",
			expected: true,
		},
		{
			name:     "Invalid email",
			input:    "test@gmail",
			expected: false,
		},
		{
			name:     "Empty email",
			input:    "",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := logic.CheckEmail(c.input)
			assert.Equal(t, c.expected, got)
		})
	}

}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		_ = logic.Sum(1, 2, 3, 4, 5)
	}
}

func BenchmarkCheckEmail(b *testing.B) {
	for b.Loop() {
		_ = logic.CheckEmail("test@gmail.com")
	}
}
