package orderedobject

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
)

// Pair represents a key value pair.
type Pair[V any] struct {
	// Key is the pair's key.
	Key string

	// Value is the pair's value.
	Value V
}

// Object represents a JSON object that respects insertion order.
type Object[V any] []*Pair[V]

// NewObject creates a new Object with the specified capacity.
func NewObject[V any](capacity int) *Object[V] {
	obj := make(Object[V], 0, capacity)
	return &obj
}

// Set sets key in object with the given value.
//
// The key is replaced if it already exists.
func (object *Object[V]) Set(key string, value V) {
	for i, pair := range *object {
		if pair.Key == key {
			(*object)[i].Value = value
			return
		}
	}

	*object = append(*object, &Pair[V]{key, value})
}

// Has reports if the given key is set.
func (object *Object[V]) Has(key string) bool {
	for _, pair := range *object {
		if pair.Key == key {
			return true
		}
	}

	return false
}

// Get gets the value of key.
//
// The returned value is V's zero value if key isn't set.
func (object *Object[V]) Get(key string) V {
	for _, pair := range *object {
		if pair.Key == key {
			return pair.Value
		}
	}

	// "hack" to get V's zero value
	var empty V
	return empty
}

// MarshalJSON encodes the object into JSON format, respecting insertion order in the process.
func (object *Object[V]) MarshalJSON() ([]byte, error) {
	var builder bytes.Buffer
	encoder := jsoniter.NewEncoder(&builder)
	encoder.SetEscapeHTML(false)

	builder.WriteString("{")

	for i, pair := range *object {
		if i > 0 {
			builder.WriteString(",")
		}

		builder.WriteString(`"`)
		builder.WriteString(pair.Key)
		builder.WriteString(`":`)

		err := encoder.Encode(pair.Value)
		if err != nil {
			return nil, err
		}
	}

	builder.WriteString("}")

	return builder.Bytes(), nil
}
