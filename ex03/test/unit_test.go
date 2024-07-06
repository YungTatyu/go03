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

func TestIndex(t *testing.T) {
	expect(piscine.Index("this is test", "is"), 2, t)
	expect(piscine.Index("this is test", "test"), 8, t)
	expect(piscine.Index("this is test", "testt"), -1, t)
	expect(piscine.Index("", ""), -1, t)
	expect(piscine.Index("", "test"), -1, t)
	expect(piscine.Index("test", ""), -1, t)
}
