package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) {
	if body, err := io.ReadAll(r.Body); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(body, v); err != nil {
			log.Fatal(err)
		}
	}
}
