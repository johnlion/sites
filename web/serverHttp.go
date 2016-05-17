package web

import (
	"fmt"
	"github.com/johnlion/sites/config"
	"github.com/johnlion/sites/proxy"
	"net/http"
	"path/filepath"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 下午1:37
 * Func: ServeHTTP
 * Desc:
 *********************************************/
func (w *Web) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if req.Host == config.LOCAL_DOMAIN_DOMAIN_1 {
		res.WriteHeader(http.StatusOK)
		/* 从数据库到得HTML内容 */
		var html string
		err, html := w.Router(req.RequestURI)
		if err != nil { //没有取到,从远程服务器获取

			ObjProxy := proxy.Proxy_constract(*target, *scheme, req, w.LocalDomain) //实例化代理

			ObjProxy.ReProxy(res, req)
			//time.Sleep( time.Second )
		} else {
			fmt.Fprint(res, html)
		}

	}
	w.Debug(req.Host)

	if req.Host == config.LOCAL_RESOURCE_DOMAIN_1 {
		path := config.TEMPLATE_BASE_DIR + req.URL.Path
		w.Debug(filepath.Join("novel", path))
		http.ServeFile(res, req, path)
	}

	/* Log */
	//w.Log(  res.Header().Get( "User-Agent" ) )
}
