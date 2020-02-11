package yamlnillable

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func Example() {
	val := Int64Of(12345)
	marshal, err := yaml.Marshal(val)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", marshal)
	// output: 12345
}

func Example_nil_value() {
	var val *Int64
	marshal, err := yaml.Marshal(val)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", marshal)
	// output: null
}
