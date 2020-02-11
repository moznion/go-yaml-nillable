package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyFloat64(t *testing.T) {
	given := Float64{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Float64
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestFloat64(t *testing.T) {
	var given float64 = 123.45
	marshal, err := yaml.Marshal(Float64Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("123.45\n"), marshal)

	var got Float64
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Float64{Val: given, IsAssigned: true}, got)
}
