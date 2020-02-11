package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyUint8(t *testing.T) {
	given := Uint8{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Uint8
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestUint8(t *testing.T) {
	var given uint8 = 123
	marshal, err := yaml.Marshal(Uint8Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("123\n"), marshal)

	var got Uint8
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Uint8{Val: given, IsAssigned: true}, got)
}
