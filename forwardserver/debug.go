package forwardserver

import (
"fmt"
"runtime"
"github.com/johnlion/sites/config"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月26日 上午10:38
 * Func: Debug
 * Desc: 输出调试信息到Command Line和LOG File
 *********************************************/
func (f *Forwardserver ) Debug( str string ){
	/* DEBUG */
	if ( config.DEBUG ){
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf( "[cgl][debug] %s:%d \n" , file,line )
		fmt.Printf( "[cgl][debug][content] ...\n" )
		fmt.Printf( "   %s\n", str )
		fmt.Printf( "[cgl][debug][content] ... end\n" )
	}
}

func (f *Forwardserver ) DebugArray( str []string ){
/* DEBUG */
	if ( config.DEBUG ){
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf( "[cgl][debug] %s:%d \n" , file,line )
		fmt.Printf( "[cgl][debug][content] ...\n" )
		fmt.Printf( "   %s\n", str )
		fmt.Printf( "[cgl][debug][content] ... end\n" )
	}
}







