package web

import (
	"net/http"
	"fmt"

)
/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 下午1:37
 * Func: ServeHTTP
 * Desc:
 *********************************************/
func ( w *Web ) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	w.Debug( req.URL.Path  )
	var html string
	err, html := w.Router( req )
	if err != nil{
	}
	if html ==""{
		html = "This is the default html page."
	}


	/* 设置 Response Header */
	for i,v := range w.Header{
		res.Header().Set( i, v )
	}

	/* web 输出  */
	fmt.Fprint( res, html )

	/* Log */
	w.Log(  res.Header().Get( "User-Agent" ) )


}
