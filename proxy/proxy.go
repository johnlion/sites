package proxy

import (
	"strings"
	"os"
	"log"
	"bufio"
	"fmt"
	"io"
	"runtime"
	"path"
	"net/http"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月27日 上午11:46
 * File: Proxy
 * Desc: 代理列表类
 * Example
	ObjProxy := proxy.Proxy_constract( *target, *protocol )
 *********************************************/
type Proxy struct{
	ProxyList  []string
	Target string
	Protocol string
	Url string
}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月27日 下午2:35
 * Func: Proxy_constract
 * Desc: 构造函数
 * Param:
 * target string
 * protocol string
 *********************************************/
func Proxy_constract( target string, protocol string ,req *http.Request ) *Proxy{
	var proxy Proxy
	proxy.GetProxyList("")
	proxy.Target = target
	proxy.Protocol = protocol
	proxy.Url = protocol + "://" + target + req.RequestURI
	return &proxy
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc: 返回[]string
 * ****************************************/
func ( p *Proxy) GetProxyList( file string ) ([]string,error){
	if strings.Contains( file, ""){
		file = "proxy.txt"
	}

	p.ReadLine( file, p.ProcessLine )
	return p.ProxyList,nil
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:读取文件行,返回[]bytes,否则反回 nil
 * ****************************************/
func ( p *Proxy) ReadLine( filePath string, hookfn func( []byte ) )  error{
	currentDir := p.GetCurrentPath()//当前文件夹
	filePath = currentDir + "/"  + filePath//文件名称
	file, err := os.Open( filePath )

	if err != nil{
		log.Fatal( err )
	}
	defer  file.Close()
	bfRd := bufio.NewReader( file )
	fmt.Println( "ProxyList ReadLine ......Start" )
	for{
		line, err := bfRd.ReadBytes( '\n' )
		hookfn( line )  //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil{  //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF{
				fmt.Println( "ProxyList ReadLine ......End" )
				return nil
			}
			return err
		}
	}

	return nil
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:反回当前包文件路径,值为 string
 * ****************************************/
func ( p *Proxy) GetCurrentPath() string{
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return  path.Dir(filename)
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:处理行数据,并写入全局变量 proxyList
 * ****************************************/
func ( p *Proxy) ProcessLine( line []byte ){
	p.ProxyList = append( p.ProxyList , strings.TrimSpace(  string( line)    ) )
	fmt.Printf( "%s\n","...ProxyList->... " +  strings.TrimSpace(  string( line)    )  )
}


