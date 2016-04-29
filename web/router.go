package web

import (
	"errors"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月28日 上午11:49
 * Func: Router
 * Desc: 路由
 *********************************************/
func ( w *Web ) Router( requestURI string ) (error,string){

	if requestURI == "/xxx" {
		return nil,"This is a valid router"
	}

	return errors.New( requestURI + " is a unvalid router"),""

}

