package pipeline

import (
	"math/rand"
	//"crypto/rand"
	"encoding/binary"
	//"fmt"
	"io"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		sort.Ints(a)
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}
func ReaderSource(r io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		buf := make([]byte, 8)
		for {
			n, err := r.Read(buf)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buf))
				out <- v
			}
			if err != nil {
				//fmt.Println("error:", err.Error())
				break
			}

		}
		close(out)
	}()
	return out
}
func WriteSink(w io.Writer, in <-chan int) {
	for v := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		w.Write(buf)
	}
}
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
