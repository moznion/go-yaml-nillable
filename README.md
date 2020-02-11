# go-yaml-nillable

A library that provides nillable basic types for YAML marshaling and unmarshaling for golang.

## Synopsis

### non-nil value

```go
import (
	"fmt"

	"github.com/moznion/go-yaml-nillable"
	"gopkg.in/yaml.v2"
)

func main() {
	val := yamlnillable.Int64Of(12345)
	marshal, err := yaml.Marshal(val)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", marshal)
	// output: 12345
}
```

### nil value

```go
import (
	"fmt"

	"github.com/moznion/go-yaml-nillable"
	"gopkg.in/yaml.v2"
)

func main() {
	var val *yamlnillable.Int64
	marshal, err := yaml.Marshal(val)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", marshal)
	// output: 12345
}
```

## Supported types

- `bool` => `yamlnillable.Bool`
- `string` => `yamlnillable.String`
- `int` => `yamlnillable.Int`
- `int8` => `yamlnillable.Int8`
- `int16` => `yamlnillable.Int16`
- `int32` => `yamlnillable.Int32`
- `int64` => `yamlnillable.Int64`
- `uint` => `yamlnillable.Uint`
- `uint8` => `yamlnillable.Uint8`
- `uint16` => `yamlnillable.Uint16`
- `uint32` => `yamlnillable.Uint32`
- `uint64` => `yamlnillable.Uint64`
- `byte` => `yamlnillable.Byte`
- `rune` => `yamlnillable.Rune`
- `float32` => `yamlnillable.Float32`
- `float64` => `yamlnillable.Float64`

## Author

moznion (<moznion@gmail.com>)

## License

```
The MIT License (MIT)
Copyright © 2020 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

