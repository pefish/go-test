package test

import "testing"

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 1, 2)
}

func TestIn(t *testing.T) {
	In(t, []interface{}{"aa", "bb"}, "aa")
}

func TestNotIn(t *testing.T) {
	NotIn(t, []interface{}{"aa", "bb"}, "aaa")
}
