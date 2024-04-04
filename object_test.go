package orderedobject

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	testCases := []struct {
		name     string
		object   *Object[any]
		expected string
	}{
		{
			name: "Empty object",
			object: func() *Object[any] {
				return NewObject[any](0)
			}(),
			expected: `{}`,
		},
		{
			name: "Single key-value pair",
			object: func() *Object[any] {
				obj := NewObject[any](1)
				obj.Set("key", "value")
				return obj
			}(),
			expected: `{"key":"value"}`,
		},
		{
			name: "Multiple key-value pairs",
			object: func() *Object[any] {
				obj := NewObject[any](3)
				obj.Set("name", "John")
				obj.Set("age", 30)
				obj.Set("city", "New York")
				return obj
			}(),
			expected: `{"name":"John","age":30,"city":"New York"}`,
		},
		{
			name: "Nested objects",
			object: func() *Object[any] {
				address := NewObject[any](2)
				address.Set("street", "123 Main St")
				address.Set("city", "London")

				person := NewObject[any](3)
				person.Set("name", "Alice")
				person.Set("age", 28)
				person.Set("address", address)

				return person
			}(),
			expected: `{"name":"Alice","age":28,"address":{"street":"123 Main St","city":"London"}}`,
		},
		{
			name: "Array of objects",
			object: func() *Object[any] {
				person1 := NewObject[any](2)
				person1.Set("name", "Bob")
				person1.Set("age", 35)

				person2 := NewObject[any](2)
				person2.Set("name", "Charlie")
				person2.Set("age", 40)

				people := NewObject[any](1)
				people.Set("people", []any{person1, person2})

				return people
			}(),
			expected: `{"people":[{"name":"Bob","age":35},{"name":"Charlie","age":40}]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := json.Marshal(tc.object)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if string(encoded) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, string(encoded))
			}
		})
	}
}
