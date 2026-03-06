package morestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestReverseRunes_empty_asserts(t *testing.T) {
	in := ""
	result := ReverseRunes(in)
	assert.Equal(t, "", result, "empty in, empty out")
}

func TestReverseRunes_asserts(t *testing.T) {
	in := "The answer to the Ultimate Question of Life, The Universe, and Everything"
	want := "gnihtyrevE dna ,esrevinU ehT ,efiL fo noitseuQ etamitlU eht ot rewsna ehT"
	result := ReverseRunes(in)
	assert.Equal(t, want, result, "non-empty in, reversed out")
}

func TestReverseRunesWithErr_empty(t *testing.T) {
	in := ""
	want := ""
	if got, err := ReverseRunesWithErr(in); err == nil || got != want {
		t.Errorf("ReverseRunesWithErr(%q) = %q, want %q, err = %v", in, got, want, err)
	}
}

func TestReverseRunesWithErr_empty_asserts(t *testing.T) {
	in := ""
	result, err := ReverseRunesWithErr(in)
	assert.NotNil(t, err, "empty in, error expected")
	assert.Equal(t, "", result, "empty in, empty out")
}

func TestReverseRunesWithErr(t *testing.T) {
	in := "The answer to the Ultimate Question of Life, The Universe, and Everything"
	want := "gnihtyrevE dna ,esrevinU ehT ,efiL fo noitseuQ etamitlU eht ot rewsna ehT"
	if got, err := ReverseRunesWithErr(in); err != nil || got != want {
		t.Errorf("ReverseRunesWithErr(%q) = %q, want %q, err = %v", in, got, want, err)
	}
}

func TestReverseRunesWithErr_asserts(t *testing.T) {
	in := "The answer to the Ultimate Question of Life, The Universe, and Everything"
	want := "gnihtyrevE dna ,esrevinU ehT ,efiL fo noitseuQ etamitlU eht ot rewsna ehT"
	result, err := ReverseRunesWithErr(in)
	assert.NoError(t, err, "non-empty in, no error expected")
	assert.Equal(t, want, result, "non-empty in, reversed out")
}
