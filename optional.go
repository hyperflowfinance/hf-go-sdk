package hyperflow

import "encoding/json"

type Optional[T any] struct {
	Set   bool
	value *T
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{Set: true, value: &value}
}

func (o Optional[T]) graphOptional() (any, bool) {
	if o.value == nil {
		return nil, o.Set
	}
	return *o.value, o.Set
}

func (o Optional[T]) IsZero() bool { return !o.Set }

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.Set {
		return []byte("null"), nil
	}
	value, _ := o.graphOptional()
	return json.Marshal(value)
}
