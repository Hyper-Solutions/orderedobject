# orderedobject

[![Go Reference](https://pkg.go.dev/badge/github.com/Hyper-Solutions/orderedobject.svg)](https://pkg.go.dev/github.com/Hyper-Solutions/orderedobject)

`orderedobject` is a simple Go package that provides an `Object` type for representing JSON objects that respect insertion order. It uses generics to allow for type-safe usage with any value type.

## Installation

To install the package, use `go get`:

```shell
go get github.com/Hyper-Solutions/orderedobject
```

## Usage

Here's a simple usage example:

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Hyper-Solutions/orderedobject"
)

func main() {
    // Create a new object with a capacity of 3
    obj := orderedobject.NewObject[any](3)

    // Set key-value pairs
    obj.Set("name", "John")
    obj.Set("age", 30)
    obj.Set("city", "New York")

    // Marshal the object to JSON
    encoded, err := json.Marshal(obj)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(encoded))
    // Output: {"name":"John","age":30,"city":"New York"}
}
```

The `Object` type provides the following methods:

- `NewObject[V any](capacity int) *Object[V]`: Creates a new `Object` with the specified capacity.
- `Set(key string, value V)`: Sets a key-value pair in the object. If the key already exists, its value is replaced.
- `Has(key string) bool`: Checks if a key is set in the object.
- `Get(key string) V`: Retrieves the value associated with a key. If the key is not set, the zero value of type `V` is returned.
- `MarshalJSON() ([]byte, error)`: Marshals the object to JSON, respecting the insertion order of key-value pairs.

## License

This package is licensed under the [MIT License](LICENSE).