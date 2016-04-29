package proxy

import (
	"net/http"
	"github.com/johnlion/sites/config"
	"time"
	"math/rand"
	"fmt"
	"net/url"
	"io/ioutil"
	"github.com/johnlion/sites/seo"
	"github.com/johnlion/mahonia"

)

func ( p *Proxy ) ReProxy( res http.ResponseWriter, req *http.Request ){

	originReq, err := http.NewRequest( req.Method, p.Url, req.Body )
	p.Check( err )
	//p.copyHeader(req.Header, &originReq.Header)
	// Create a client and query the target

	for{
		//创建随机种子
		randNum := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := randNum.Intn( len( p.ProxyList ) )

		// --var transport http.Transport
		var transport http.Transport
		if  config.PROXY_STATUS {
			transport = *p.GetTransportFieldURL(p.ProxyList[i])
		}
		originResp, err := transport.RoundTrip( originReq )
		if ( err !=nil ){
			continue
		}
		fmt.Printf("Resp-Headers: %v\n", originResp.Header);
		defer originResp.Body.Close()

		//res.WriteHeader(originResp.StatusCode)
		body, err := ioutil.ReadAll(originResp.Body)
		if err !=nil  {
			continue
		}
		html := string( body )

		dH := res.Header()
		p.copyHeader(originResp.Header, &dH)
		dH.Add("Requested-Host", originReq.Host)


		ObjSeo := seo.Seo_constract( p.Target, p.Protocol )
		html = string( ObjSeo.RegProcess( html ) )

		//输出正常编码数据
		if mahonia.GetCharset(html) == nil {
			p.RequestUrLFileGroupSave( req.RequestURI ,html )
			//saveDataToRedis( replacedHtml )
			res.Write( []byte (html) )
			//res.WriteHeader( 200 )
			//fmt.Fprint( res, html )
			break
		}else{
			enc := mahonia.NewEncoder("utf8")
			dec := mahonia.NewDecoder("utf8")
			html , ok := dec.ConvertStringOK( html )
			if ok {
				p.RequestUrLFileGroupSave( req.RequestURI,html )
				res.Write( []byte (html) )
				break
			}
			dec = mahonia.NewDecoder("gb18030")
			html , ok = dec.ConvertStringOK( html )
			if ok {
				html , ok = enc.ConvertStringOK( html )
				if ok {
					p.RequestUrLFileGroupSave( req.RequestURI,html )
					res.Write( []byte (html) )
					break
				}
			}
		}
	}
}

func ( p *Proxy ) GetTransportFieldURL( proxy_addr string ) (transport *http.Transport)  {
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(  "http://" +  proxy_addr )
	fmt.Println( url_proxy )
	transport = &http.Transport{Proxy : http.ProxyURL(url_proxy)}
	return
}

func (p *Proxy) copyHeader(source http.Header, dest *http.Header) {
	for n, v := range source {
		for _, vv := range v {
			dest.Add(n, vv)
		}
	}
}
