package stream

import "sort"

func chanDatumConvertVal[T any](resultDataCh <-chan *Data[T]) []T {
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

func valsConvertChanDatum[T any](vals []T) chan *Data[T] {
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
