package test

import (
	"fmt"
	"github.com/nxbnxb/go-stream/stream"
	"strconv"
	"testing"
)

func TestStream(t *testing.T) {
	sc := stream.Of([]int{1, 2, 3, 4, 5, 6, 7, 8, 89, 2341, 23, 1, 5, 12, 31, 5, 12, 3, 12})
	top := sc.Filter(func(t int) bool {
		return t > 10
	}).
		Foreach(func(t int) (int, error) {
			return t + 1, nil
		}).Top(10)
	fmt.Println(top)
	fmt.Println(sc)
}

func TestStreamConvert(t *testing.T) {
	sc := stream.Of([]int{1, 2, 3, 4, 5, 6, 7, 8, 89, 2341, 23, 1, 5, 12, 31, 5, 12, 3, 12})
	top := sc.
		Filter(func(t int) bool {
			return t < 10
		}).
		Foreach(func(t int) (int, error) {
			return t * t, nil
		}).Top(11)

	scStr := stream.OfConvert(top, func(t int) string {
		return strconv.Itoa(t)
	})
	strings := scStr.Top(10)

	fmt.Println(top)
	fmt.Println(strings)
}
