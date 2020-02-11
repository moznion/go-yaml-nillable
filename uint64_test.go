package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyUint64(t *testing.T) {
	given := Uint64{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Uint64
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestUint64(t *testing.T) {
	var given uint64 = 12345
	marshal, err := yaml.Marshal(Uint64Of(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("12345\n"), marshal)

	var got Uint64
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Uint64{Val: given, IsAssigned: true}, got)
}
