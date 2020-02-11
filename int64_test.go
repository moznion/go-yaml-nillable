package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyInt64(t *testing.T) {
	given := Int64{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Int64
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestInt64(t *testing.T) {
	var given int64 = 12345
	marshal, err := yaml.Marshal(Int64Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("12345\n"), marshal)

	var got Int64
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Int64{Val: given, IsAssigned: true}, got)
}
