package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// i.e. Read() must implement the io.Reader interface
func (r MyReader) Read(b []byte) (n int, e error) {
    for i := range b {
        b[i] = 'A'
    }
    return len(b), nil
}

func main() {
    reader.Validate(MyReader{})
}