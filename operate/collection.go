package operate

import (
	"log/slog"
	"math/rand"
	"strings"
	"sync"
)

func ForeachFn[E any](fn func(e E) E) func([]E) []E {
	return func(es []E) []E {
		for i, e := range es {
			es[i] = fn(e)
		}
		return es
	}
}

func ForeachE[E any](datas []E, fn func(e E) E) []E {
	for i, e := range datas {
		datas[i] = fn(e)
	}
	return datas
}

func GroupBy[K comparable, V any](datas []V, f func(v V) K) map[K][]V {
	groupByMap := make(map[K][]V, len(datas))
	for _, data := range datas {
		key := f(data)
		groupByMap[key] = append(groupByMap[key], data)
	}
	return groupByMap
}

func ShuffleE[E any](candidates []E) []E {
	lenC := len(candidates)
	if lenC <= 2 {
		return candidates
	}
	randNum := rand.Intn(lenC)
	return append(candidates[randNum:], candidates[:randNum]...)
}

func MaxE[E ~int | ~int32 | ~float64 | ~float32](es []E) E {
	e := es[0]
	for i := 1; i < len(es); i++ {
		if e < es[i] {
			e = es[i]
		}
	}
	return e
}

func MapVal2Slice[K comparable, V any](data map[K]V) []V {
	news := make([]V, 0, len(data))
	for _, v := range data {
		news = append(news, v)
	}
	return news
}

func MapKey2Slice[K comparable, V any](data map[K]V) []K {
	news := make([]K, 0, len(data))
	for k := range data {
		news = append(news, k)
	}
	return news
}

func Slice2Set[K comparable](data []K) map[K]bool {
	set := make(map[K]bool, len(data))
	for i := 0; i < len(data); i++ {
		set[data[i]] = true
	}
	return set
}
func SliceCutByStepE[T any](datas []T, step int) [][]T {
	lenDatas := len(datas)
	if step > lenDatas {
		return [][]T{datas}
	}
	cutSlice := make([][]T, 0, lenDatas/step+1)
	i := 0
	for ; i < lenDatas-step; i += step {
		cutSlice = append(cutSlice, datas[i:i+step])
	}
	if i < lenDatas {
		cutSlice = append(cutSlice, datas[i:])
	}
	return cutSlice
}
func SliceAdapterE[S, T any](data []S, fn func(S) T) []T {
	slice := make([]T, 0, len(data))
	for i := 0; i < len(data); i++ {
		slice = append(slice, fn(data[i]))
	}
	return slice
}

func CopySliceE[E any](es []E) []E {
	newES := make([]E, 0, len(es))
	for _, e := range es {
		newES = append(newES, e)
	}
	return newES
}

func Slice2SetE[K comparable, V any](data []V, fn func(v V) K) map[K]bool {
	set := make(map[K]bool, len(data))
	for i := 0; i < len(data); i++ {
		set[fn(data[i])] = true
	}
	return set
}

func Slice2MapE[K comparable, V, Val any](data []V, fn func(v V) (K, Val)) map[K]Val {
	set := make(map[K]Val, len(data))
	for i := 0; i < len(data); i++ {
		k, v := fn(data[i])
		set[k] = v
	}
	return set
}

func ReverseE[E comparable](words []E) []E {
	lenWords := len(words)
	if lenWords <= 1 {
		return words
	} else if lenWords == 2 {
		words[0], words[1] = words[1], words[0]
		return words
	}
	mid := len(words) / 2
	if lenWords%2 == 0 {
		mid = mid - 1
	}

	left, right := 0, len(words)-1
	for i := mid; i >= 0; i-- {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return words
}

func DistinctE[E any](entitys []E, fn func(e E) string) []E {
	newE := make([]E, 0)
	distincSet := make(map[string]bool, len(entitys))
	for _, e := range entitys {
		key := fn(e)
		if distincSet[key] {
			continue
		}
		newE = append(newE, e)
		distincSet[key] = true
	}
	return newE
}

func DistinctA[E any](entitys []E, fn func(e E) any) []E {
	newE := make([]E, 0)
	distincSet := make(map[any]bool, len(entitys))
	for _, e := range entitys {
		key := fn(e)
		if distincSet[key] {
			continue
		}
		newE = append(newE, e)
		distincSet[key] = true
	}
	return newE
}

func ContainerE[E any](entitys []E, e E, fn func(e E) string) bool {
	for _, v := range entitys {
		if fn(v) == fn(e) {
			return true
		}
	}
	return false
}

func ContainerSetE[E comparable, V any](set map[E]bool, vs []V, fn func(v V) E) bool {
	return IndexE(set, vs, fn) > 0
}

func IndexE[E comparable, V any](set map[E]bool, vs []V, fn func(v V) E) int {
	count := 0
	for _, v := range vs {
		if set[fn(v)] {
			count++
		}
	}
	return count
}

func FilterE[E any](filters []E, filterFn func(e E) bool) []E {
	newE := make([]E, 0, len(filters)/2)
	for _, e := range filters {
		if filterFn(e) {
			continue
		}
		newE = append(newE, e)
	}
	return newE
}

func TopE[E any](datas []E, top int) []E {
	if top > len(datas) {
		top = len(datas)
	}
	return datas[:top]
}

func AssertE[E, V any](e E, f func(e E) V) V {
	return f(e)
}

func FindE[E any](finds []E, findFn func(e E) bool) []E {
	newE := make([]E, 0)
	for _, e := range finds {
		if findFn(e) {
			newE = append(newE, e)
		}
	}
	return newE
}
func StrEmpty(e string) bool {
	return strings.TrimSpace(e) == ""
}

// 注意 现在有点问题E 可以是指针类型 也可以不是
func ForeachMMR[E any](es []E, doFn func(e E) (E, error), taskNum int) []E {
	slog.Info("ForeachMMR Start With", slog.Int("len", len(es)), slog.Int("threadNum", taskNum))
	if len(es) < taskNum {
		var err error
		for e, data := range es {
			es[e], err = doFn(data)
			if err != nil {
				slog.Error("ForeachMMR", slog.String("err", err.Error()))
			}
		}
		return es
	}
	chs := make(chan E, len(es))
	for _, e := range es {
		chs <- e
	}
	resultChs := make(chan E, len(es))
	close(chs)
	wg := &sync.WaitGroup{}
	wg.Add(taskNum)
	for i := 0; i < taskNum; i++ {
		go func() {
			defer wg.Done()
			for e := range chs {
				result, err := doFn(e)
				if err != nil {
					slog.Error("ForeachMMR", slog.String("err", err.Error()))
				}
				resultChs <- result
			}
		}()
	}
	wg.Wait()
	close(resultChs)
	rets := make([]E, 0, len(resultChs))
	for result := range resultChs {
		rets = append(rets, result)
	}
	return rets
}

// 注意 现在有点问题E 可以是指针类型 也可以不是
func ForeachMMRFn[E any](doFn func(e E) (E, error), taskNum int) func([]E) ([]E, error) {
	return func(es []E) ([]E, error) {
		slog.Info("ForeachMMRFn Start With", slog.Int("len", len(es)), slog.Int("threadNum", taskNum))
		if len(es) < taskNum {
			var err error
			for e, data := range es {
				es[e], err = doFn(data)
				if err != nil {
					slog.Error("ForeachMMRFn", slog.String("err", err.Error()))
				}
			}
			return es, err
		}
		chs := make(chan E, len(es))
		for _, e := range es {
			chs <- e
		}
		resultChs := make(chan E, len(es))
		close(chs)
		wg := &sync.WaitGroup{}
		wg.Add(taskNum)
		for i := 0; i < taskNum; i++ {
			go func() {
				defer wg.Done()
				for e := range chs {
					result, err := doFn(e)
					if err != nil {
						slog.Error("ForeachMMRFn", slog.String("err", err.Error()))
					}
					resultChs <- result
				}
			}()
		}
		wg.Wait()
		close(resultChs)
		rets := make([]E, 0, len(resultChs))
		for result := range resultChs {
			rets = append(rets, result)
		}
		return rets, nil
	}
}

// 注意 现在有点问题E 可以是指针类型 也可以不是
func ForeachMMRAdapterFn[D, T any](doFn func(e D) (T, error), taskNum int) func([]D) ([]T, error) {
	return func(ds []D) ([]T, error) {
		slog.Info("ForeachMMRAdapterFn Start With", slog.Int("len", len(ds)), slog.Int("threadNum", taskNum))

		if len(ds) < taskNum {
			var err error
			var target T
			tragets := make([]T, 0, len(ds))
			for _, data := range ds {
				target, err = doFn(data)
				if err != nil {
					slog.Error("ForeachMMRAdapterFn", slog.String("err", err.Error()))
				}
				tragets = append(tragets, target)
			}
			return tragets, nil
		}
		chs := make(chan D, len(ds))
		for _, e := range ds {
			chs <- e
		}
		resultChs := make(chan T, len(ds))
		close(chs)
		wg := &sync.WaitGroup{}
		wg.Add(taskNum)
		for i := 0; i < taskNum; i++ {
			go func() {
				defer wg.Done()
				for e := range chs {
					result, err := doFn(e)
					if err != nil {
						slog.Error("ForeachMMRAdapterFn", slog.String("err", err.Error()))
					}
					resultChs <- result
				}
			}()
		}
		wg.Wait()
		close(resultChs)
		rets := make([]T, 0, len(resultChs))
		for result := range resultChs {
			rets = append(rets, result)
		}
		return rets, nil
	}
}

func SliceAdapterFn[S, T any](fn func(S) T) func(ss []S) []T {
	return func(ss []S) []T {
		slice := make([]T, 0, len(ss))
		for i := 0; i < len(ss); i++ {
			slice = append(slice, fn(ss[i]))
		}
		return slice
	}

}

func SliceSelectedAdapterFn[S, T any](fn func(S) (T, bool)) func(ss []S) []T {
	return func(ss []S) []T {
		slice := make([]T, 0, len(ss))
		for i := 0; i < len(ss); i++ {
			if t, ok := fn(ss[i]); ok {
				slice = append(slice, t)
			}
		}
		return slice
	}

}

func SliceAdapterFnIndex[S, T any](fn func(s S, index int) T) func(ss []S) []T {
	return func(ss []S) []T {
		slice := make([]T, 0, len(ss))
		for i := 0; i < len(ss); i++ {
			slice = append(slice, fn(ss[i], i))
		}
		return slice
	}

}

// 获取数组的最后一个
func SliceLast[E any](es []E) (e E, err error) {
	lenES := len(es)
	if lenES == 0 {
		return e, ERR_PARAM_NIL
	} else if lenES == 1 {
		return es[0], nil
	} else {
		return es[lenES-1], nil
	}
}

func UnionSlice[K comparable](slicess ...[]K) []K {
	guessLen := len(slicess) * len(slicess[0])
	allSet := make(map[K]bool, guessLen)
	resultSlice := make([]K, 0, guessLen)

	for _, slices := range slicess {
		for _, itemU := range slices {
			if allSet[itemU] {
				continue
			}
			allSet[itemU] = true
			resultSlice = append(resultSlice, itemU)
		}
	}
	return resultSlice
}

func FilterEFn[E any](filterFn func(e E) bool) func(es []E) []E {
	return func(es []E) []E {
		newE := make([]E, 0, len(es))
		for _, e := range es {
			if filterFn(e) {
				continue
			}
			newE = append(newE, e)
		}
		return newE
	}
}

func CountEFn[E any](find func(e E) bool) func(es []E) int {
	return func(es []E) int {
		newE := make([]E, 0, len(es))
		for _, e := range es {
			if !find(e) {
				continue
			}
			newE = append(newE, e)
		}
		return len(newE)
	}
}

func FindEFn[E any](find func(e E) bool) func(es []E) []E {
	return func(es []E) []E {
		newE := make([]E, 0, len(es))
		for _, e := range es {
			if !find(e) {
				continue
			}
			newE = append(newE, e)
		}
		return newE
	}
}

// 求两个字符串数组的交集
func SliceIntersect(slice1, slice2 []string) []string {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil
	}
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times > 0 {
			nn = append(nn, v)
		}
	}
	return nn
}

func Join[T any](fn func(t T) string) func(ts []T, sep string) string {
	return func(ts []T, sep string) string {
		strs := SliceAdapterFn(fn)(ts)
		return strings.Join(strs, sep)
	}
}
