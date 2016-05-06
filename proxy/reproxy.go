package proxy

import (
	"fmt"
	"github.com/johnlion/sites/config"
	"github.com/johnlion/sites/seo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"github.com/djimenez/iconv-go"

)


func (p *Proxy) ReProxy(res http.ResponseWriter, req *http.Request) {

	//数据库访问
	if  p.dataExisted( res , req  ) {
		p.visitByDatabase(res, req)
	}else{
		//非数据库访问
		p.visitByRemoteHost( res, req )
	}
}


func (p *Proxy) GetTransportFieldURL(proxy_addr string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse("http://" + proxy_addr)
	fmt.Println(url_proxy)
	transport = &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	return
}

func (p *Proxy) CopyHeader(source http.Header, dest *http.Header) {
	for n, v := range source {
		for _, vv := range v {
			dest.Add(n, vv)
		}
	}
}

func (p *Proxy) CopyCookies( source http.Cookie, dest *http.Cookie ){

}


/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月5日 下午2:32
 * Func: dataExisted
 * Desc: 数据是否存在; 存在返回true,否则反回false
 *********************************************/
func (p *Proxy) dataExisted ( res http.ResponseWriter, req *http.Request ) bool{
	return p.RedisKeyExisted( req )
}

func (p *Proxy) visitByDatabase( res http.ResponseWriter, req *http.Request ){
	p.Debug( "Redis model" )
	body := ""
	p.RedisServer( res, req ,body )

}

func (p *Proxy) visitByRemoteHost( res http.ResponseWriter, req *http.Request ){
	cookie1 := &http.Cookie{Name: "sample", Value: "sample", HttpOnly: false}
	http.SetCookie(res, cookie1)
	originReq, err := http.NewRequest(req.Method, p.Url, req.Body)
	p.Check(err)
	//p.copyHeader(req.Header, &originReq.Header)
	// Create a client and query the target

	for {
		//创建随机种子
		randNum := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := randNum.Intn(len(p.ProxyList))

		// --var transport http.Transport
		var transport http.Transport
		if config.PROXY_STATUS {
			transport = *p.GetTransportFieldURL(p.ProxyList[i])
		}
		originResp, err := transport.RoundTrip(originReq)
		if err != nil {
			continue
		}
		fmt.Printf("Resp-Headers: %v\n", originResp.Header)
		defer originResp.Body.Close()

		//res.WriteHeader(originResp.StatusCode)
		body, err := ioutil.ReadAll(originResp.Body)
		if err != nil {
			continue
		}


		dH := res.Header()


		p.CopyHeader(originResp.Header, &dH)
		dH.Add("Requested-Host", originReq.Host)


		out:=make([]byte,len(body) *2 )
		out=out[:]
		fmt.Println( len( string( body ) ) )



		_ , _, err = iconv.Convert( body, out, "gbk", "utf-8" )
		if err == nil {
			//ioutil.WriteFile("output.html", out, 0666)


			//SEO && DECODE
			html := string(out)
			ObjSeo := seo.Seo_constract(p.Target, p.Scheme)
			resourceHtml := ObjSeo.RegProcess(html)

			html = string(resourceHtml)

			p.RequestUrLFileGroupSave(req, html)
			res.Write(resourceHtml)
			break
		}else{
			//SEO && DECODE
			html := string(body)
			ObjSeo := seo.Seo_constract(p.Target, p.Scheme)
			resourceHtml := ObjSeo.RegProcess(html)

			html = string(resourceHtml)

			p.RequestUrLFileGroupSave(req, html)
			res.Write(resourceHtml)
			break
		}





		//输出正常编码数据
		/*
		if mahonia.GetCharset(html) == nil {
			p.RequestUrLFileGroupSave(req, html)
			//saveDataToRedis( replacedHtml )
			res.Write([]byte(html))
			//res.WriteHeader( 200 )
			//fmt.Fprint( res, html )
			break
		} else {
			enc := mahonia.NewEncoder("utf8")
			dec := mahonia.NewDecoder("utf8")
			html, ok := dec.ConvertStringOK(html)
			if ok {
				p.RequestUrLFileGroupSave(req, html)
				res.Write([]byte(html))
				break
			}
			dec = mahonia.NewDecoder("gb18030")
			html, ok = dec.ConvertStringOK(html)
			if ok {
				html, ok = enc.ConvertStringOK(html)
				if ok {
					p.RequestUrLFileGroupSave(req, html)
					res.Write([]byte(html))
					break
				}
			}
		}
		*/
	}
}
