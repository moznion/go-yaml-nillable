// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Uint16 is a data type for nillable uint16 value for YAML marshaling/unmarshaling.
type Uint16 struct {
	Val        uint16
	IsAssigned bool
}

// Uint16Of makes a non-nil value with given uint16.
func Uint16Of(val uint16) *Uint16 {
	return &Uint16{Val: val, IsAssigned: true}
}

// MarshalYAML marshals Uint16 as YAML.
// This method used on marshaling YAML internally.
func (v *Uint16) MarshalYAML() (interface{}, error) {
	if v.IsAssigned {
		return v.Val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Uint16 as YAML.
// This method used on unmarshaling YAML internally.
func (v *Uint16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint16
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.Val = val
	v.IsAssigned = true
	return nil
}
