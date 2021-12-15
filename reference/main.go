package main

import(
	"fmt"
)

type sample struct {
	arr []string
	m map[string]string
}

func newSample() sample {
	s := sample{
		arr: []string{},
		m: map[string]string{},
	}
	fmt.Printf("s: %p, arr: %p, m: %p\n", &s, &(s.arr), &(s.m))
	return s
}

func newSampleRef() *sample {
	fmt.Printf("newSampleRef\n")
	s := sample{
		arr: []string{},
		m: map[string]string{},
	}
	fmt.Printf("s: %p, arr: %p, m: %p\n", &s, &(s.arr), &(s.m))
	return &s
}	

func main() {
	s := newSample()
	fmt.Printf("s: %p, arr: %p, m: %p\n", &s, &s.arr, &s.m)

	sr := newSampleRef()
	fmt.Printf("s: %p, arr: %p, m: %p\n", sr, &sr.arr, &sr.m)
}
