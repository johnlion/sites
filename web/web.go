package web

import (
	"flag"
	"net/http"
	"github.com/johnlion/sites/config"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月26日 下午3:28
 * File: web.go
 * Desc: web类 处理http相关业务
 * Example:
	import (
		"github.com/johnlion/sites/web"
		"fmt"

	)
	func main(){
		ObjWeb := web.Web_constract()
		ObjWeb.RunWebServer()
		fmt.Println( *ObjWeb.Domain )
	}

 *********************************************/

type Web struct {
	LocalDomain string              //local domain
	Domain *string                  //remote domain
	Protocol *string
	Server map[int]*http.Server
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
	web.LocalDomain = config.DOMAIN
	web.Domain = target
	web.Protocol = protocol
	//web.Header = map[string]string{
	//		"Server":"www.baidu.com",
	//		"Content-Type":"text/plain; charset=utf-8",
	//		"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.86 Safari/537.36",
	//		"Referer":"www.baidu.com",
	//}

	return &web
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



