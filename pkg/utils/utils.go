package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ong-gtp/book-mgt/pkg/errors"
)

type emptyOk struct {
	Message string
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func Ok(res []byte, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func OkEmpty(message string, w http.ResponseWriter) {
	m := emptyOk{message}
	res, err := json.Marshal(m)
	errors.ErrorCheck(err)
	Ok(res, w)
}
