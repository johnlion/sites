package web

import (
	"flag"
	"net/http"
	"fmt"
)

type Web struct {
	Domain *string
	Protocol *string

}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午10:56
 * Func: Web_contract
 * Desc: 类构造器
 *********************************************/
func Web_constract() *Web{
	var web Web
	flag.Parse()
	web.Domain = target
	web.Protocol = protocol
	return &web
}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午11:09
 * Func: RunWebServer
 * Desc: web package entry
 *********************************************/
func ( w *Web ) RunWebServer( ){
	http.ListenAndServe(":9090",  w )
}

func ( w*Web ) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf( res, "This is a test!" + *w.Domain )

}
