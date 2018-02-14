package main

import (
	"bufio"
	"fmt"
	"os"
	"wp/pipeline"
)

func main() {
	const fname = "large.in"
	const n = 100000000
	f, err := os.Create(fname)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	p := pipeline.RandomSource(n)
	w := bufio.NewWriter(f)
	pipeline.WriteSink(w, p)
	w.Flush()
	f, err = os.Open(fname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	p = pipeline.ReaderSource(bufio.NewReader(f))
	for v := range p {
		fmt.Println(v)
	}
}
func MemDemo() {
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(2, 1, 6, 32, 67, 9)),
		pipeline.InMemSort(pipeline.ArraySource(2, 1, 6, 32, 67, 9)))
	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}
