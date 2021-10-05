package helpers

import (
	"encoding/json"
	"io"
)

// JsonEncodeForce force encode to json
func JsonEncodeForce(data interface{}) []byte {
	r, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return r
}

// JsonEncodeStringForce force encode to json
func JsonEncodeStringForce(data interface{}) string {
	r, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(r)
}

// JsonDecodeReader decode json from reader interface
func JsonDecodeReader(r io.Reader, data interface{}) error {
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return err
	}
	return nil
}
