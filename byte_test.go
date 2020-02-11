package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyByte(t *testing.T) {
	given := Byte{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Byte
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestByte(t *testing.T) {
	var given byte = 0x67
	marshal, err := yaml.Marshal(ByteOf(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("103\n"), marshal)

	var got Byte
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Byte{Val: given, IsAssigned: true}, got)
}
