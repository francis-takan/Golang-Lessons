package main

import (
	"testing"
)

func mainTest(t *testing.T) {
	if add(5, 5) != 10 {
		t.Error("Test fail")
	}
}
