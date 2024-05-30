package golang

import (
	"encoding/json"
)

func (g *Golang) JsonMarshal(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func (g *Golang) JsonUnmarshal(input []byte, dest interface{}) error {
	return json.Unmarshal(input, dest)
}
