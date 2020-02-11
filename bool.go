// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Bool is a data type for nillable bool value for YAML marshaling/unmarshaling.
type Bool struct {
	val        bool
	isAssigned bool
}

// BoolOf makes a non-nil value with given bool.
func BoolOf(val bool) *Bool {
	return &Bool{val: val, isAssigned: true}
}

// MarshalYAML marshals Bool as YAML.
// This method used on marshaling YAML internally.
func (v *Bool) MarshalYAML() (interface{}, error) {
	if v.isAssigned {
		return v.val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Bool as YAML.
// This method used on unmarshaling YAML internally.
func (v *Bool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bool
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.val = val
	v.isAssigned = true
	return nil
}