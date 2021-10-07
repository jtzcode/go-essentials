package client

import (
	"packages/libs"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(libs.GetFibonacci(5))
}
