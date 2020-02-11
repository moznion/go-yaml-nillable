// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Int is a data type for nillable int value for YAML marshaling/unmarshaling.
type Int struct {
	Val        int
	IsAssigned bool
}

// IntOf makes a non-nil value with given int.
func IntOf(val int) *Int {
	return &Int{Val: val, IsAssigned: true}
}

// MarshalYAML marshals Int as YAML.
// This method used on marshaling YAML internally.
func (v *Int) MarshalYAML() (interface{}, error) {
	if v.IsAssigned {
		return v.Val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Int as YAML.
// This method used on unmarshaling YAML internally.
func (v *Int) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val int
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.Val = val
	v.IsAssigned = true
	return nil
}
