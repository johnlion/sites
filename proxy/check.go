package proxy

import (
	"fmt"
	"os"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月27日 下午2:52
 * Func:
 * Desc:
 *********************************************/
func (p *Proxy) Check(err error) {
	if err != nil {
		p.Debug("error")
		fmt.Printf("%v", err)
		os.Exit(1)

	}
}
