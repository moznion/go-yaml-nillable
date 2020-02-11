package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyUint32(t *testing.T) {
	given := Uint32{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Uint32
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestUint32(t *testing.T) {
	var given uint32 = 12345
	marshal, err := yaml.Marshal(Uint32Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("12345\n"), marshal)

	var got Uint32
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Uint32{Val: given, IsAssigned: true}, got)
}
