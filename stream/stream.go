package stream

import (
	"github.com/nxbnxb/go-stream/operate"
	"runtime"
	"sort"
)

func (options *Option[T]) Filter(filter func(T) bool) *Option[T] {
	options.options = append(options.options, func(data *Data[T]) (err error) {
		if filter(data.Val) {
			data.Valid = false
		}
		return nil
	})
	return options
}

func (options *Option[T]) Foreach(handler func(T) (T, error)) *Option[T] {
	options.options = append(options.options, func(data *Data[T]) (err error) {
		data.Val, err = handler(data.Val)
		if err != nil {
			data.err = err
		}
		return err
	})
	return options
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

func (sc *Stream[T]) exec(pnum int) {
	if pnum == 0 {
		pnum = runtime.NumCPU()
	}

	var resultDataCh <-chan *Data[T]
	if pnum == 1 {
		resultDataCh = operate.Foreach(sc.options...)(sc.Datum)
		return
	}
	resultDataCh = operate.MForeach(sc.options...)(sc.Datum, pnum)

	sc.ResultVal = chanDatumAdapterVal(resultDataCh)
}

func chanDatumAdapterVal[T any](resultDataCh <-chan *Data[T]) []T {
	resultDatum := make([]Data[T], 0, len(resultDataCh))
	for data := range resultDataCh {
		if !data.checkDataValid() {
			continue
		}
		resultDatum = append(resultDatum, *data)
	}
	sort.SliceStable(resultDatum, func(i, j int) bool {
		return resultDatum[i].Index < resultDatum[j].Index
	})
	resultVal := make([]T, 0, len(resultDatum))

	for e := range resultDatum {
		resultVal = append(resultVal, resultDatum[e].Val)
	}
	return resultVal
}

func (sc *Stream[T]) Top(n int) []T {
	sc.exec(0)
	if len(sc.ResultVal) < n {
		n = len(sc.ResultVal)
	}
	sc.ResultVal = sc.ResultVal[:n]
	return sc.ResultVal
}

func Adapter[T, N any](sc *Stream[T], adapter func(t T) N) *Stream[N] {
	ns := make([]N, 0, len(sc.ResultVal))
	for _, v := range sc.ResultVal {
		ns = append(ns, adapter(v))
	}
	return Of(ns)
}
