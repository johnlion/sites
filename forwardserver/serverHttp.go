package forwardserver

import (
	"net/http"

	"io"
)

func (f *Forwardserver)  ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World")
}
