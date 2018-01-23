package main

import "golang.org/x/tour/reader"

// MyReader : Emit only the character 'A' whatever it reads
type MyReader struct {
	read string
	done bool
}

// Read : Add a Read(p []byte) (int, error) method to MyReader.
func (r *MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(&MyReader{})
}
