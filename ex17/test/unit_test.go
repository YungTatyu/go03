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

func TestAtoiBase(t *testing.T) {
	expect(piscine.AtoiBase("9223372036854775807", "0123456789"), 9223372036854775807, t)
	expect(piscine.AtoiBase("9223372036854775808", "0123456789"), -9223372036854775808, t)
	expect(piscine.AtoiBase("7FFFFFFFFFFFFFFF", "0123456789ABCDEF"), 9223372036854775807, t)
	expect(piscine.AtoiBase("7FFFFFFFFFFFFFFF", ""), 0, t)
	expect(piscine.AtoiBase("7FFFFFFFFFFFFFFF", "a"), 0, t)
	expect(piscine.AtoiBase("7FFFFFFFFFFFFFFF", "aa"), 0, t)
	expect(piscine.AtoiBase("7FFFFFFFFFFFFFFF", "12-"), 0, t)
	expect(piscine.AtoiBase("7FFFFFFFFFFFFFFF", "12+"), 0, t)
}
