package util

import (
	"encoding/json"
)

func JSONMarshal(v any) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func JSONUnmarshal(s string, v any) error {
	err := json.Unmarshal([]byte(s), v)
	if err != nil {
		return err
	}
	return nil
}
