package test

import (
	"piscine"
	"testing"
)

func expect(actual, expect int, t *testing.T) {
	if actual != expect {
		t.Errorf("expected %d, but got %d", expect, actual)
	}
}

func TestCompare(t *testing.T) {
	expect(piscine.Compare("a", "a"), 0, t)
	expect(piscine.Compare("b", "a"), 1, t)
	expect(piscine.Compare("a", "b"), -1, t)
	expect(piscine.Compare("test", "t"), 1, t)
	expect(piscine.Compare("test", "testa"), -1, t)
	expect(piscine.Compare("test", "test"), 0, t)
}
