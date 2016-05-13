package seo
import (
	"regexp"
	"fmt"
	"strings"

	"path/filepath"
	"github.com/johnlion/sites/config"
)

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func: RegProcess
 * desc: 返回[]string
 * ****************************************/
func ( s *Seo )RegProcess( text string ) []byte {

	//invokeMethod1 := "Reg_" +  strings.Replace(  target, ".", "_", -1  )
	//seo.Debug( invokeMethod2 )
	//InvokeObjectMethod( &seo, invokeMethod1 )

	extension := filepath.Base( s.RequestURI  )
	reg := regexp.MustCompile( config.REG_FILE_SUFFIX )
	extension = reg.FindString( extension )

	switch extension {
	case ".png":
		//do code
		break
	case ".jpg":
		//do code
		break
	case ".jpeg":
		//do code
		break
	case ".gif":
		//do code
		break
	case ".svg":
		//do code
		break
	case ".bmp":
		//do code
		break
	case ".tiff":
		break
	case ".webp":
		//do code
		break
	case ".ico":
		//do code
		break
	case ".vico":
		//do code
		break
	case ".js":
		//do code
		break
	case ".html":
		//do code
		invokeMethod2 := "Reg_Spider_" +  strings.Replace(  s.Target, ".", "_", -1  )
		InvokeObjectMethod( s, invokeMethod2 ,text )
		text = s.text
		fmt.Println( text )
		break
	case ".htm":
		invokeMethod2 := "Reg_Spider_" +  strings.Replace(  s.Target, ".", "_", -1  )
		InvokeObjectMethod( s, invokeMethod2 ,text )
		text = s.text
		fmt.Println( text )
		break
	default:
		invokeMethod2 := "Reg_Spider_" +  strings.Replace(  s.Target, ".", "_", -1  )
		InvokeObjectMethod( s, invokeMethod2 ,text )
		text = s.text
		fmt.Println( text )
		break
	}







	/* 取得seo数据,并替换字符串 */
	for i,val := range s.RegParterns{
		fmt.Println( i )
		fmt.Println( val )
		reg := regexp.MustCompile( val )
		text = reg.ReplaceAllString(text,  i  )
	}
	return []byte(text)

}