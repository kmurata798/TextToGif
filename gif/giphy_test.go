package gif

import (
	"testing"
)

func TestModifyInput(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
