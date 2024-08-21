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

func Of[T any](vals []T) *Stream[T] {
	sc := &Stream[T]{
		Datum: adapter2Data(vals),
	}
	return sc
}

func adapter2Data[T any](vals []T) chan *Data[T] {
	var datumCh = make(chan *Data[T], len(vals))
	for e, val := range vals {
		datumCh <- &Data[T]{
			Val:   val,
			Valid: true,
			Index: e,
		}
	}
	close(datumCh)
	return datumCh
}
