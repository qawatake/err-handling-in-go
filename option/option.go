package main

type Option[T any] struct {
	v *T
}

func NewOpt[T any](v T) Option[T] {
	return Option[T]{v: &v}
}

func (o Option[T]) V() (T, bool) {
	if o.v == nil {
		var zero T
		return zero, false
	}
	return *o.v, true
}
