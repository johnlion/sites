package web

import (
	"flag"
	"net/http"
	"log"
)

type Web struct {
	Domain *string
	Protocol *string
	Server *http.Server
	Header map[string]string

}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午10:56
 * Func: Web_contract
 * Desc: 类构造器
 *********************************************/
func Web_constract() *Web{
	flag.Parse()

	/* 初始化属性 */
	var web Web
	web.Domain = target
	web.Protocol = protocol
	web.Header = map[string]string{
			"Server":"www.baidu.com",
			"Content-Type":"text/plain; charset=utf-8",
			"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Safari/537.36",
			"Referer":"www.baidu.com",
	}

	return &web
}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午11:09
 * Func: RunWebServer
 * Desc: web package entry  运行并监听
 *********************************************/
func ( w *Web ) RunWebServer( ){
	//http.ListenAndServe( DOMAIN_ADDR ,  w )
	w.Server = &http.Server{
		Addr:           DOMAIN_ADDR,
		Handler:        w,
		ReadTimeout:    HTTP_READ_TIME_OUT,
		WriteTimeout:   HTTP_WRITE_TIME_OUT,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(w.Server.ListenAndServe())

}


func (w *Web ) SetHeader ( header map[string]string ){
	w.Header =  header
}






/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 下午2:55
 * Func: copyHeader
 * Desc: 复制Respose
 *********************************************/
func (w *Web ) CopyHeader(source http.Header, dest *http.Header) *http.Header{
	for n, v := range source {
		for _, vv := range v {
			dest.Add(n, vv)
		}
	}
	return dest
}



