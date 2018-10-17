package maputil

import (
	"testing"
	"gotest.tools/assert"
)

func TestMapKeysBasic(t *testing.T) {
	m := map[string]rune {
		"A" : 'A',
		"B" : 'B',
		"C" : 'C',
	}

	keys := MapKeys(m)
	realKeys, ok := keys.([]string)
	assert.Assert(t, ok)
	assert.Assert(t, len(realKeys) == len(m))
	for _, k := range realKeys {
		_, ok := m[k]
		assert.Assert(t, ok)
	}
}
