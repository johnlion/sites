package proxy

import(
	"os"
	"time"
	"github.com/johnlion/sites/config"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月27日 下午3:41
 * Func: Log
 * Desc: 日志写入到文件,并返回true;否则,反回false
 *********************************************/
func ( p *Proxy ) Log( str string  ) bool{
	/* LOG */
	if ( !config.LOG_WRITE ){
		return false
	}

	/* 检测路径是否存在 */
	if _, err := os.Stat( config.LOG_BASE_DIR ); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll( config.LOG_BASE_DIR,  0777 )
		if ( err != nil ){
			panic( err )
		}
	}

	/* 打开文件 */
	fo,err := os.OpenFile( config.LOG_DIR_FILE , os.O_APPEND|os.O_WRONLY |os.O_CREATE  , 0777  )
	p.Check(err )

	/* 写入文件 */
	str = "[cgl][log][" + time.Now().Format("2006-01-02 15:04:05")  +  "] " +  str + "\n";
	_, err = fo.WriteString(str);
	p.Check( err )

	defer p.Check( fo.Close() )
	return true

	// close fi on exit and check for its returned error
	//os.Exit(1)

}
