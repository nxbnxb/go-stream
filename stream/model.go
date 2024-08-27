package stream

type Stream[T any] struct {
	Option[T]
	Datum     chan *Data[T]
	ResultVal []T
}

type Option[T any] struct {
	options []func(*Data[T]) error
}

type Data[T any] struct {
	Val   T
	Index int
	Valid bool
	err   error
}
