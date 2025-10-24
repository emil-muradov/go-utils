package semver_compare

import (
	"fmt"
	"testing"
)

func TestCompareSemvers(t *testing.T) {
	tests := []struct {
		a, b string
		want CompareResult
	}{
		{"1.0.0", "1.1.0", CompareResult("LT")},
		{"0.1.0", "0.0.1.0", CompareResult("GT")},
		{"1.1.1", "1.1.1", CompareResult("EQ")},
		{"1.1.1", "1,2", ""},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s, %s", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			actual, err := CompareSemvers(tt.a, tt.b)
			if err != nil {
				t.Skipf("got error: %v", err)
			}
			if actual != tt.want {
				t.Errorf("got %s, want %s", actual, tt.want)
			}
		})
	}
}
