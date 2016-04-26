package web

import (
	"net/http"
	"errors"
)

func ( w *Web ) Router( req *http.Request ) (error,string){
	if req.URL.Path == "/xxx" {
		return nil,"This is a valid router"
	}

	return errors.New( req.URL.Path + " is a unvalid router"),""

}