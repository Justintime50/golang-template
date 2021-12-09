package project

import (
	"testing"
)

func TestMyProject(t *testing.T) {
	outString := myProject("Justintime50")
	expectedValue := "hello Justintime50"

	if outString != expectedValue {
		t.Errorf("myProject does not match, got: %s want: %s", outString, expectedValue)
	}
}
