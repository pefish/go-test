package test

import "testing"

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 1, 2)
}