package operate

import (
	"log/slog"
	"reflect"
	"sync"
)

// MDoFn
// 注意 现在有点问题E必须是指针类型 才会对结果产生影响  否则就是用来执行任务的函数
func MDoFn[E any](doFns ...func(e E) error) func(es <-chan E, goNum int) {
	return func(es <-chan E, goNum int) {
		wg := &sync.WaitGroup{}
		wg.Add(goNum)
		for i := 0; i < goNum; i++ {
			go func() {
				defer wg.Done()
				var err error
				for e := range es {
					for _, doFn := range doFns {
						err = doFn(e)
						if err != nil {
							slog.Error("[MDoFn]doFn", "err", err.Error())
						}
					}

				}
			}()
		}
		wg.Wait()
	}
}

// DoFn
// 注意 现在有点问题E必须是指针类型 才会对结果产生影响  否则就是用来执行任务的函数
func DoFn[E any](doFns ...func(e E) error) func(es <-chan E) {
	return func(es <-chan E) {

		var err error
		for e := range es {
			for _, doFn := range doFns {
				err = doFn(e)
				if err != nil {
					slog.Error("[MDoFn]doFn", "err", err.Error())
				}
			}
		}

	}
}

// Foreach
func Foreach[E any](doFns ...func(e E) error) func(es <-chan E) <-chan E {
	return func(es <-chan E) <-chan E {
		results := make(chan E, len(es))
		defer close(results)
		foreach(doFns...)(es, results)
		return results
	}
}

// MForeach
func MForeach[E any](doFns ...func(e E) error) func(es <-chan E, goNum int) <-chan E {
	return func(es <-chan E, goNum int) <-chan E {
		wg := &sync.WaitGroup{}
		results := make(chan E, len(es))
		defer close(results)
		wg.Add(goNum)
		execFn := foreach(doFns...)
		for i := 0; i < goNum; i++ {
			go func() {
				defer wg.Done()
				execFn(es, results)
			}()
		}
		wg.Wait()
		return results
	}
}

// foreach
func foreach[E any](doFns ...func(e E) error) func(es <-chan E, results chan E) {
	return func(es <-chan E, results chan E) {

		var invoke E
		typeOf := reflect.TypeOf(invoke)

		var err error
		if typeOf.Kind() == reflect.Ptr {
			for item := range es {
				for _, doFn := range doFns {
					err = doFn(item)
					if err != nil {
						slog.Error("[MDoFn]doFn", "err", err.Error())
					}
				}
				results <- item
			}
		} else {
			for item := range es {
				itemP := &item
				for _, doFn := range doFns {
					err = doFn(item)
					if err != nil {
						slog.Error("[MDoFn]doFn", "err", err.Error())
					}
				}
				results <- *itemP
			}
		}
		return
	}
}
