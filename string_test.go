package yamlnillable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEmptyString(t *testing.T) {
	given := String{}
	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.Equal(t, []byte("null\n"), marshal)

	var got String
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, given, got)
}

func TestString(t *testing.T) {
	given := "hello, world"
	marshal, err := yaml.Marshal(StringOf(given))
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello, world\n"), marshal)

	var got String
	err = yaml.Unmarshal(marshal, &got)
	assert.NoError(t, err)
	assert.Equal(t, String{Val: given, IsAssigned: true}, got)
}
