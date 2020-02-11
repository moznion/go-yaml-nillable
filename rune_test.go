package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyRune(t *testing.T) {
	given := Rune{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Rune
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestRune(t *testing.T) {
	given := 'g'
	marshal, err := yaml.Marshal(RuneOf(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("103\n"), marshal)

	var got Rune
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Rune{Val: given, IsAssigned: true}, got)
}
