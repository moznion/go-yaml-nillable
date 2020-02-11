package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyInt32(t *testing.T) {
	given := Int32{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Int32
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestInt32(t *testing.T) {
	var given int32 = 12345
	marshal, err := yaml.Marshal(Int32Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("12345\n"), marshal)

	var got Int32
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Int32{Val: given, IsAssigned: true}, got)
}
