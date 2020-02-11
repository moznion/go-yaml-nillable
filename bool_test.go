package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyBool(t *testing.T) {
	given := Bool{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got Bool
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestBool(t *testing.T) {
	given := false
	marshal, err := yaml.Marshal(BoolOf(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("false\n"), marshal)

	var got Bool
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, Bool{Val: given, IsAssigned: true}, got)
}
