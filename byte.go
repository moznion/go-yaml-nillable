// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Byte is a data type for nillable byte value for YAML marshaling/unmarshaling.
type Byte struct {
	val        byte
	isAssigned bool
}

// ByteOf makes a non-nil value with given byte.
func ByteOf(val byte) *Byte {
	return &Byte{val: val, isAssigned: true}
}

// MarshalYAML marshals Byte as YAML.
// This method used on marshaling YAML internally.
func (v *Byte) MarshalYAML() (interface{}, error) {
	if v.isAssigned {
		return v.val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Byte as YAML.
// This method used on unmarshaling YAML internally.
func (v *Byte) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val byte
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.val = val
	v.isAssigned = true
	return nil
}
