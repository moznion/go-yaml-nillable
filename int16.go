// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Int16 is a data type for nillable int16 value for YAML marshaling/unmarshaling.
type Int16 struct {
	val        int16
	isAssigned bool
}

// Int16Of makes a non-nil value with given int16.
func Int16Of(val int16) *Int16 {
	return &Int16{val: val, isAssigned: true}
}

// MarshalYAML marshals Int16 as YAML.
// This method used on marshaling YAML internally.
func (v *Int16) MarshalYAML() (interface{}, error) {
	if v.isAssigned {
		return v.val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Int16 as YAML.
// This method used on unmarshaling YAML internally.
func (v *Int16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val int16
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.val = val
	v.isAssigned = true
	return nil
}