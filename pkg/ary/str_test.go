package ary_test

import (
	"peanut/pkg/ary"
	"testing"
)

func TestInArray(t *testing.T) {
	isInTests := []struct {
		description string
		element     string
		array       []string
		expect      bool
	}{
		{
			description: "String value with true",
			element:     "dog",
			array:       []string{"dog", "cat", "monkey", "pig", "chicken"},
			expect:      true,
		},
		{
			description: "String value with false",
			element:     "apple",
			array:       []string{"dog", "cat", "monkey", "pig", "chicken"},
			expect:      false,
		},
		{
			description: "String value with empty array",
			element:     "apple",
			array:       []string{},
			expect:      false,
		},
		{
			description: "String value with empty element",
			element:     "",
			array:       []string{"dog", "cat", "monkey", "pig", "chicken"},
			expect:      false,
		},
	}

	for _, tc := range isInTests {
		if result := ary.InArray(tc.element, tc.array); result != tc.expect {
			t.Errorf("Failed test case: %s", tc.description)
		}
	}
}
