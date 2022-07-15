package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Do(t *testing.T) {
	tests := map[string]struct {
		s      string
		i      int
		b      bool
		exp    string
		expErr string
	}{
		"success1":  {"a", 1, true, "[1A]", ""},
		"success2":  {"b", 3, true, "[3B]", ""},
		"success3":  {"c", 2, true, "[2C]", ""},
		"success4":  {"d", 5, true, "[5D]", ""},
		"success5":  {"a", 1, false, "[1a]", ""},
		"success6":  {"b", 3, false, "[3b]", ""},
		"success7":  {"c", 2, false, "[2c]", ""},
		"success8":  {"d", 5, false, "[5d]", ""},
		"success9":  {"c", 13, true, "C", ""},
		"success10": {"d", 21, true, "D", ""},
		"success11": {"c", 34, false, "c", ""},
		"success12": {"d", 13, false, "d", ""},

		"invalid s1":  {"a", 4, true, "", "invalid s"},
		"invalid s2":  {"b", 7, true, "", "invalid s"},
		"invalid s3":  {"c", 9, true, "", "invalid s"},
		"invalid s4":  {"d", 10, true, "", "invalid s"},
		"invalid s5":  {"a", 11, false, "", "invalid s"},
		"invalid s6":  {"b", 50, false, "", "invalid s"},
		"invalid s7":  {"f", 2, false, "", "invalid s"},
		"invalid s8":  {"g", 5, false, "", "invalid s"},
		"invalid s9":  {"H", 13, true, "", "invalid s"},
		"invalid s10": {"Q", 21, true, "", "invalid s"},
		"invalid s11": {"q", 34, false, "", "invalid s"},
		"invalid s12": {"n", 13, false, "", "invalid s"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Do(tt.s, tt.i, tt.b)
			assert.Equal(t, tt.exp, actual)

			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
