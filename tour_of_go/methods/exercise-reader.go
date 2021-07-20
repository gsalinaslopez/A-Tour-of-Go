package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
    in_count := 0
    for i := 0; i < len(b); i++ {
        b[i] = 'A'
        in_count += 1
    }
    return in_count, nil
}

func main() {
    reader.Validate(MyReader{})
}
