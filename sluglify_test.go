package sluglify

import (
	"testing"
)

func TestGetSlugString(t *testing.T) {
	input := "-123abcdeЦУЦ&#&&$^$^7-azcx23"
	want := "abcdetsutsand-andzzz"

	got := GetSlugString(input, 20, "zzz")

	if got != want {
		t.Errorf("GetSlugString() = %q, want %q", got, want)
	}

	if len(got) != 20 {
		t.Errorf("GetSlugString() = %q, length must be 20", got)
	}
}
