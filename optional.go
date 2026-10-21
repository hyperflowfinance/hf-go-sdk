package hyperflow

import "encoding/json"

// Optional distinguishes an omitted nullable GraphQL input from a supplied value.
type Optional[T any] struct {
	Set   bool
	value *T
}

// Some marks value as present, including when value is its zero value.
func Some[T any](value T) Optional[T] {
	return Optional[T]{Set: true, value: &value}
}

func (o Optional[T]) graphOptional() (any, bool) {
	if o.value == nil {
		return nil, o.Set
	}
	return *o.value, o.Set
}

// MarshalJSON encodes a present value, or null when no value was supplied.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.Set {
		return []byte("null"), nil
	}
	value, _ := o.graphOptional()
	return json.Marshal(value)
}
