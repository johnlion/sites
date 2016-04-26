package web

import (
	"fmt"
	"runtime"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月26日 上午10:38
 * Func: Debug
 * Desc: 输出调试信息到Command Line和LOG File
 *********************************************/
func (w *Web ) Debug( str string ){
	/* DEBUG */
	if ( DEBUG ){
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf( "[cgl][debug] %s:%d \n" , file,line )
		fmt.Printf( "[cgl][debug][content] ...\n" )
		fmt.Printf( "   %s\n", str )
		fmt.Printf( "[cgl][debug][content] ... end\n" )
	}
}




