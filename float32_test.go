package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyFloat32(t *testing.T) {
	given := Float32{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Float32
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestFloat32(t *testing.T) {
	var given float32 = 123.45
	marshal, err := yaml.Marshal(Float32Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("123.45\n"), marshal)

	var got Float32
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Float32{Val: given, IsAssigned: true}, got)
}
