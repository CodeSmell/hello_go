package morestrings

import (
	"testing"
)

func TestReverseRunes_empty(t *testing.T) {
	in := ""
	want := ""
	if got := ReverseRunes(in); got != want {
		t.Errorf("ReverseRunes(%q) = %q, want %q", in, got, want)
	}
}

func TestReverseRunes(t *testing.T) {
	in := "The answer to the Ultimate Question of Life, The Universe, and Everything"
	want := "gnihtyrevE dna ,esrevinU ehT ,efiL fo noitseuQ etamitlU eht ot rewsna ehT"
	if got := ReverseRunes(in); got != want {
		t.Errorf("ReverseRunes(%q) = %q, want %q", in, got, want)
	}
}

func TestReverseRunesWithErr_empty(t *testing.T) {
	in := ""
	want := ""
	if got, err := ReverseRunesWithErr(in); err == nil || got != want {
		t.Errorf("ReverseRunesWithErr(%q) = %q, want %q, err = %v", in, got, want, err)
	}
}

func TestReverseRunesWithErr(t *testing.T) {
	in := "The answer to the Ultimate Question of Life, The Universe, and Everything"
	want := "gnihtyrevE dna ,esrevinU ehT ,efiL fo noitseuQ etamitlU eht ot rewsna ehT"
	if got, err := ReverseRunesWithErr(in); err != nil || got != want {
		t.Errorf("ReverseRunesWithErr(%q) = %q, want %q, err = %v", in, got, want, err)
	}
}
