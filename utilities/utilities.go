package utilities

import (
	"encoding/json"
	"io"
)

func FromJson[T any](rw io.ReadCloser, body *T) error {

	encoder := json.NewDecoder(rw)
	return encoder.Decode(body)
}
