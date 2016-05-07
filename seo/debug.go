package seo
import (
	"fmt"
	"runtime"
	"github.com/johnlion/sites/config"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月7日 上午9:13
 * Func: Debug
 * Desc: 输出调试信息到Command Line和LOG File
 *********************************************/
func (s *Seo ) Debug( str string ){
	/* DEBUG */
	if ( config.DEBUG ){
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf( "[cgl][debug] %s:%d \n" , file,line )
		fmt.Printf( "[cgl][debug][content] ...\n" )
		fmt.Printf( "   %s\n", str )
		fmt.Printf( "[cgl][debug][content] ... end\n" )
	}
}

func (s *Seo ) DebugArray( str []string ){
	/* DEBUG */
	if ( config.DEBUG ){
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf( "[cgl][debug] %s:%d \n" , file,line )
		fmt.Printf( "[cgl][debug][content] ...\n" )
		fmt.Printf( "   %s\n", str )
		fmt.Printf( "[cgl][debug][content] ... end\n" )
	}
}



