package forwardserver
import (
"net/http"
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

type Forwardserver struct {
	LocalDomain string  //local domain
	Domain      *string //remote domain
	Scheme      *string
	Server      map[int]*http.Server
	Header      map[string]string
}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午10:56
 * Func: Web_contract
 * Desc: 类构造器
 *********************************************/
func Forwardserver_constract() *Forwardserver {
/* 初始化属性 */
var forwardserver Forwardserver


return &forwardserver
}

