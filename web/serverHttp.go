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

	/* 设置 Response Header */
	for i,v := range w.Header{
		res.Header().Set( i, v )
	}

	/* web 输出  */
	fmt.Fprint( res, "This is a test" )

	/* Debug */
	w.Debug( "This is a test" )

	/* Log */
	w.Log(  res.Header().Get( "User-Agent" ) )


}
