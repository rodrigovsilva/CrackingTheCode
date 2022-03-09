package fundamentals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse_reverseString(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []string
	}{
		"test 1": {
			input:    []string{"b", "a"},
			expected: []string{"a", "b"},
		},
		"test 2": {
			input:    []string{"d", "c", "b", "a"},
			expected: []string{"a", "b", "c", "d"},
		},
		"test 3": {
			input:    []string{"ywcicc", "pvwpwq", "tjonlu", "ptpcpy", "sjpuqo", "jvmmyi", "akkrwx", "jauefl", "flcdgm", "jsczwb", "qziblu", "muaxxu", "lzrbki", "jxhugo", "unbcgf", "ssopji", "hmxeot", "torxcw", "nzzods", "gffbhz", "adjnte", "erxrvt", "wwzrzx", "yczghl", "micnbc", "apgipf", "ipwbhr", "niwpze", "qkehgz", "eqebqi", "zaqgpk", "ipplqb", "vtdgox", "virvuy", "wzsnty", "nscfqu", "wbptbv", "eftqfv", "hvqbdc", "grxeaz", "nxdkwt", "hqmvfo", "zdhwcb", "asszut", "jqthkg", "lidijj", "zrzadt", "haqwvy", "lfalkn", "enrppw"},
			expected: []string{"enrppw", "lfalkn", "haqwvy", "zrzadt", "lidijj", "jqthkg", "asszut", "zdhwcb", "hqmvfo", "nxdkwt", "grxeaz", "hvqbdc", "eftqfv", "wbptbv", "nscfqu", "wzsnty", "virvuy", "vtdgox", "ipplqb", "zaqgpk", "eqebqi", "qkehgz", "niwpze", "ipwbhr", "apgipf", "micnbc", "yczghl", "wwzrzx", "erxrvt", "adjnte", "gffbhz", "nzzods", "torxcw", "hmxeot", "ssopji", "unbcgf", "jxhugo", "lzrbki", "muaxxu", "qziblu", "jsczwb", "flcdgm", "jauefl", "akkrwx", "jvmmyi", "sjpuqo", "ptpcpy", "tjonlu", "pvwpwq", "ywcicc"},
		},
		"test 4": {
			input:    []string{"b"},
			expected: []string{"b"},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, reverseString(tt.input))
		})
	}
}
