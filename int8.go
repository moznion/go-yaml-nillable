// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Int8 is a data type for nillable int8 value for YAML marshaling/unmarshaling.
type Int8 struct {
	Val        int8
	IsAssigned bool
}

// Int8Of makes a non-nil value with given int8.
func Int8Of(val int8) *Int8 {
	return &Int8{Val: val, IsAssigned: true}
}

// MarshalYAML marshals Int8 as YAML.
// This method used on marshaling YAML internally.
func (v *Int8) MarshalYAML() (interface{}, error) {
	if v.IsAssigned {
		return v.Val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Int8 as YAML.
// This method used on unmarshaling YAML internally.
func (v *Int8) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val int8
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.Val = val
	v.IsAssigned = true
	return nil
}
