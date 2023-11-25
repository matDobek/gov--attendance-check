package predicates

import "testing"

func TestPresent(t *testing.T) {
	tests := [][]any{
		{[]string{}, false},
		{[]string{"a"}, true},
	}

	for _, test := range tests {
		arg0 := test[0].([]string)
		expectation := test[1]
		result := Present(arg0)

		if result != expectation {
			t.Errorf("for %#v, got: %v, expected: %v", arg0, result, expectation)
		}
	}
}

func TestContains(t *testing.T) {
	tests := [][]any{
		{[]string{}, "", false},
		{[]string{}, "a", false},
		{[]string{"a"}, "", false},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{"a", "b", "c"}, "a", true},
	}

	for _, test := range tests {
		arg0 := test[0].([]string)
		arg1 := test[1].(string)
		expectation := test[2]
		result := Contains(arg0, arg1)

		if result != expectation {
			t.Errorf("for %#v, %#v got: %v, expected: %v", arg0, arg1, result, expectation)
		}
	}
}
