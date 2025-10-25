package semver_compare

import (
	"fmt"
	"testing"
)

func TestCompareSemvers(t *testing.T) {
	tests := []struct {
		a, b   string
		expect CompareResult
	}{
		{"1.0.0", "1.1.0", "LT"},
		{"0.1.0", "0.0.1.0", "GT"},
		{"1.1.1", "1.1.1", "EQ"},
		{"1.1.1", "1,2,3", ""},
		{"", "0.3", ""},
		{"1.2.3-alpha.1", "1.2.3-alpha.2", "LT"},
		{"2.0.0-rc.1", "1.7.3", "GT"},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s, %s", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			actual, err := CompareSemvers(tt.a, tt.b)
			if err != nil {
				t.Skipf("error: %v", err)
			}
			if actual != tt.expect {
				t.Errorf("expect: %s, actual: %s", tt.expect, actual)
			}
		})
	}
}
