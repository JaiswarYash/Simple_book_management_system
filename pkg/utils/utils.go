package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if r.Body == nil {
		return nil
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(body, x)
}
