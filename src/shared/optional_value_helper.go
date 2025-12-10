package shared

import (
	"encoding/json"
	"fmt"
)

type Optional[T any] struct {
	Value   T
	IsSet   bool
	Default T
}

func NilValue[T any]() Optional[T] {
	return Optional[T]{IsSet: false}
}

// Function to create an optional value
func SetValue[T any](v T) Optional[T] {
	return Optional[T]{Value: v, IsSet: true}
}

// Function to create an optional value with a default value
func SetValueWithDefault[T any](v T, defaultValue T) Optional[T] {
	return Optional[T]{Value: v, IsSet: true, Default: defaultValue}
}

// Function to get the value of an optional value
func (o Optional[T]) GetValue() T {
	if !o.IsSet {
		return o.Default
	}
	return o.Value
}

func (o Optional[T]) ValueOrDefault(defaultValue T) T {
	if o.IsSet {
		return o.Value
	}
	return defaultValue
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.IsSet {
		// If the value is not set, return 'null' to omit the field's data structure
		return []byte("null"), nil
	}
	// If set, marshal only the underlying Value
	return json.Marshal(o.Value)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		// If the incoming JSON value is 'null', the field is not set.
		o.IsSet = false
		return nil
	}

	// Otherwise, the value is set. Try to unmarshal the data into the Value field.
	if err := json.Unmarshal(data, &o.Value); err != nil {
		return fmt.Errorf("failed to unmarshal optional value: %w", err)
	}

	// Successfully unmarshaled a non-null value
	o.IsSet = true
	return nil
}
