package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyUint16(t *testing.T) {
	given := Uint16{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Uint16
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestUint16(t *testing.T) {
	var given uint16 = 12345
	marshal, err := yaml.Marshal(Uint16Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("12345\n"), marshal)

	var got Uint16
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Uint16{Val: given, IsAssigned: true}, got)
}
