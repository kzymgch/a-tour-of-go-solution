package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (my MyReader) Read(rb []byte) (i int, e error) {
	for i, e = 0, nil; i < len(rb); i++ {
		rb[i] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
