// Code generated by go-yaml-nillable generator; DO NOT EDIT.

package yamlnillable

// Uint64 is a data type for nillable uint64 value for YAML marshaling/unmarshaling.
type Uint64 struct {
	Val        uint64
	IsAssigned bool
}

// Uint64Of makes a non-nil value with given uint64.
func Uint64Of(val uint64) *Uint64 {
	return &Uint64{Val: val, IsAssigned: true}
}

// MarshalYAML marshals Uint64 as YAML.
// This method used on marshaling YAML internally.
func (v *Uint64) MarshalYAML() (interface{}, error) {
	if v.IsAssigned {
		return v.Val, nil
	}
	return nil, nil
}

// UnarshalYAML unmarshals Uint64 as YAML.
// This method used on unmarshaling YAML internally.
func (v *Uint64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val uint64
	if err := unmarshal(&val); err != nil {
		return err
	}
	v.Val = val
	v.IsAssigned = true
	return nil
}
