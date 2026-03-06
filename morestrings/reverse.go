// Package with additional functions to manipulate strings
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
// from the How to Write Go Code post
// the capital letter means this function is exported
// and can be used by other packages that import this
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
