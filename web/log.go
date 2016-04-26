package web

import(
	"os"
	"time"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月26日 下午1:55
 * Func: Log
 * Desc: 日志写入到文件,并返回true;否则,反回false
 *********************************************/
func ( w *Web ) Log( str string  ) bool{
	/* LOG */
	if ( !LOG_WRITE ){
		return false
	}

	/* 检测路径是否存在 */
	if _, err := os.Stat( LOG_BASE_DIR ); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll( LOG_BASE_DIR,  0777 )
		if ( err != nil ){
			panic( err )
		}
	}

	/* 打开文件 */
	fo,err := os.OpenFile( LOG_DIR_FILE , os.O_APPEND|os.O_WRONLY |os.O_CREATE  , 0777  )
	w.Check(err )

	/* 写入文件 */
	str = "[cgl][log][" + time.Now().Format("2006-01-02 15:04:05")  +  "] " +  str + "\n";
	 _, err = fo.WriteString(str);
	w.Check( err )

	defer w.Check( fo.Close() )
	return true







	// close fi on exit and check for its returned error
	//os.Exit(1)

}

