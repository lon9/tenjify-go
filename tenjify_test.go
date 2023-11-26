package tenjify

import (
	"testing"
)

func TestTenjify(t *testing.T) {
	res := Tenjify("sample.png", 60, 100, false, false)
	t.Log("\n" + res)
}

func TestTenjifyReverse(t *testing.T) {
	res := Tenjify("sample.png", 60, 100, true, false)
	t.Log("\n" + res)
}

func TestTenjifyFillBlank(t *testing.T) {
	res := Tenjify("sample.png", 60, 100, true, true)
	t.Log("\n" + res)
}
