package proxy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"github.com/garyburd/redigo/redis"
	//"path/filepath"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月27日 上午11:46
 * File: Proxy
 * Desc: 代理列表类
 * Example
	ObjProxy := proxy.Proxy_constract( *target, *shceme  , req, w.LocalDomain )             //实例化代理
	ObjProxy.ReProxy( res, req, req.RequestURI )
 *********************************************/
type Proxy struct {
	ProxyList   []string
	Target      string
	Scheme      string
	Url         string
	LocalDomain string
	RedisPool   *redis.Pool
}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月27日 下午2:35
 * Func: Proxy_constract
 * Desc: 构造函数
 * Param:
 * target string        远程主机
 * scheme string      http 协议
 * req *http.Request
 * localDomain string 本地域名
 *********************************************/
func Proxy_constract(target string, scheme string, req *http.Request, localDomain string) *Proxy {
	var proxy Proxy
	proxy.GetProxyList("")
	proxy.Target = target
	proxy.Scheme = scheme
	proxy.Url = scheme + "://" + target + req.RequestURI
	proxy.LocalDomain = localDomain
	proxy.RedisPool = proxy.NewPool()
	return &proxy
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc: 返回[]string
 * ****************************************/
func (p *Proxy) GetProxyList(file string) ([]string, error) {
	if strings.Contains(file, "") {
		file = "proxy.txt"
	}

	p.ReadLine(file, p.ProcessLine)
	return p.ProxyList, nil
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:文件行,返回[]bytes,否则反回 nil
 * ****************************************/
func (p *Proxy) ReadLine(filePath string, hookfn func([]byte)) error {
	//currentDir := p.GetCurrentPath()       //当前文件夹
	//filePath = currentDir + "/" + filePath //文件名称
	//filepath :=
	file, err := os.Open("proxy.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bfRd := bufio.NewReader(file)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				fmt.Println("ProxyList ReadLine ......End")
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
func (p *Proxy) GetCurrentPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	//fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return path.Dir(filename)
}

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc:处理行数据,并写入全局变量 proxyList
 * ****************************************/
func (p *Proxy) ProcessLine(line []byte) {
	p.ProxyList = append(p.ProxyList, strings.TrimSpace(string(line)))
	//p.Debug( "...ProxyList->... " +  strings.TrimSpace(  string( line)    )  )
}
