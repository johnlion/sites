package web

import (
	"net/http"
	"github.com/johnlion/sites/proxy"
	"fmt"

	//"time"
)
/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 下午1:37
 * Func: ServeHTTP
 * Desc:
 *********************************************/
func ( w *Web ) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	/* 从数据库到得HTML内容 */
	var html string
	err, html := w.Router( req.RequestURI )
	if err != nil{//没有取到,从远程服务器获取
		ObjProxy := proxy.Proxy_constract( *target, *protocol  , req, w.LocalDomain )             //实例化代理
		ObjProxy.ReProxy( res, req )
		//time.Sleep( time.Second )
	}else{
		fmt.Fprint( res, html )
	}
	/* Log */
	//w.Log(  res.Header().Get( "User-Agent" ) )


}
