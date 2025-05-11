package main

type Result[V, E any] struct {
	v Option[V]
	e Option[E]
}

func (r *Result[V, E]) Or() (V, *E) {
	if v, ok := r.v.V(); ok {
		return v, nil
	}
	var zero V
	e, _ := r.e.V()
	return zero, &e
}

func NewValue[R Result[V, E], V, E any](v V) *R {
	return &R{v: NewOpt(v)}
}

func NewErr[R Result[V, E], V, E any](e E) *R {
	return &R{e: NewOpt(e)}
}

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
