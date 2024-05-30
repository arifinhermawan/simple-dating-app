package golang

import "io"

func (g *Golang) ReadAll(input io.Reader) ([]byte, error) {
	return io.ReadAll(input)
}
