package web

import (
	"fmt"
	"os"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月26日 下午1:56
 * Func: Check
 * Desc: 错误检测,如果发生错误,终止代码,并输出err
 *********************************************/
func (w *Web) Check(err error) {
	if err != nil {
		w.Debug("error")
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
