package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	testCases := []struct {
		input  []byte
		length int
		err    error
	}{
		{[]byte("Hello, World!"), 13, nil},
		{[]byte("Привет, Мир!"), 12, nil},
		{[]byte(""), 0, nil},
		{[]byte{227, 129, 255}, 0, ErrInvalidUTF8},
	}

	for _, tc := range testCases {
		length, err := GetUTFLength(tc.input)
		if length != tc.length || err != tc.err {
			t.Errorf("GetUTFLength(%v) = %d, %v, want %d, %v", tc.input, length, err, tc.length, tc.err)
		}
	}
}
