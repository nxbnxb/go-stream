package stream

import (
	"github.com/nxbnxb/go-stream/operate"
	"runtime"
)

func Of[T any](vals []T) *Stream[T] {
	sc := &Stream[T]{
		Datum: valsConvertChanDatum(vals),
	}
	return sc
}

func OfConvert[dst any, src any](vals []src, convert func(src src) dst) *Stream[dst] {
	dsts := make([]dst, 0, len(vals))
	for _, val := range vals {
		dsts = append(dsts, convert(val))
	}
	sc := &Stream[dst]{
		Datum: valsConvertChanDatum(dsts),
	}
	return sc
}

func (stream *Stream[T]) Filter(filter func(T) bool) *Stream[T] {
	stream.options = append(stream.options, func(data *Data[T]) (err error) {
		if filter(data.Val) {
			data.Valid = false
		}
		return nil
	})
	return stream
}

func (stream *Stream[T]) Foreach(handler func(T) (T, error)) *Stream[T] {
	stream.options = append(stream.options, func(data *Data[T]) (err error) {
		data.Val, err = handler(data.Val)
		if err != nil {
			data.err = err
		}
		return err
	})
	return stream
}

func (stream *Stream[T]) ForeachIndex(handler func(index int, t T) (T, error)) *Stream[T] {
	stream.options = append(stream.options, func(data *Data[T]) (err error) {
		data.Val, err = handler(data.Index, data.Val)
		if err != nil {
			data.err = err
		}
		return err
	})
	return stream
}

func (data *Data[T]) checkDataValid() bool {
	if !data.Valid {
		return false
	}
	if data.err != nil {
		return false
	}
	return true
}

func (stream *Stream[T]) exec(pNum int) {
	if pNum == 0 {
		pNum = runtime.NumCPU()
	}

	var resultDataCh <-chan *Data[T]
	if pNum == 1 {
		resultDataCh = operate.Foreach(stream.options...)(stream.Datum)
		return
	}
	resultDataCh = operate.MForeach(stream.options...)(stream.Datum, pNum)
	stream.ResultVal = chanDatumConvertVal(resultDataCh)
}

func (stream *Stream[T]) Top(n int) []T {
	stream.exec(0)
	if len(stream.ResultVal) < n {
		n = len(stream.ResultVal)
	}
	stream.ResultVal = stream.ResultVal[:n]
	return stream.ResultVal
}
