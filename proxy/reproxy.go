package proxy

import (
	"fmt"
	"github.com/johnlion/mahonia"
	"github.com/johnlion/sites/config"
	"github.com/johnlion/sites/seo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
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

func (p *Proxy) copyHeader(source http.Header, dest *http.Header) {
	for n, v := range source {
		for _, vv := range v {
			dest.Add(n, vv)
		}
	}
}

func (p *Proxy) dataExisted ( res http.ResponseWriter, req *http.Request ) bool{
	return true
}

func (p *Proxy) visitByDatabase( res http.ResponseWriter, req *http.Request ){
	p.Debug( "Redis model" )
	p.RedisServer( res, req )

}

func (p *Proxy) visitByRemoteHost( res http.ResponseWriter, req *http.Request ){
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
		html := string(body)

		dH := res.Header()
		p.copyHeader(originResp.Header, &dH)
		dH.Add("Requested-Host", originReq.Host)

		ObjSeo := seo.Seo_constract(p.Target, p.Scheme)
		html = string(ObjSeo.RegProcess(html))

		//输出正常编码数据
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
	}
}
