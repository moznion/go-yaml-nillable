package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyInt8(t *testing.T) {
	given := Int8{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Int8
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestInt8(t *testing.T) {
	var given int8 = 123
	marshal, err := yaml.Marshal(Int8Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("123\n"), marshal)

	var got Int8
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Int8{Val: given, IsAssigned: true}, got)
}
