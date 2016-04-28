package web

import (
	"errors"
)

func ( w *Web ) Router( requestURI string ) (error,string){
	if requestURI == "/xxx" {
		return nil,"This is a valid router"
	}

	return errors.New( requestURI + " is a unvalid router"),""

}