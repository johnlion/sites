package proxy

import (
	"github.com/johnlion/sites/config"

	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"net/http"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月28日 上午11:36
 * Func: SaveImage
 * Desc: 保存图片到指定位置
 *********************************************/
func ( p *Proxy ) SaveToImage( req *http.Request ,body string  ){
	for{

		//p.Debug( p.LocalDomain )
		//p.Debug( requestURI )
		//p.Debug( config.CACHE_DIR +  config.IMAGE_DOMAIN_1 + requestURI )

		url :=  req.Host + req.URL.RequestURI()
		fpath :=  config.CACHE_DIR + url
		dir := filepath.Dir( fpath )


		reg := regexp.MustCompile( config.REG_TOTAL_FILENAME )
		fpath = reg.FindString( fpath )


		/* 检测路径是否存在 */
		if _, err := os.Stat( dir ); os.IsNotExist(err) {
			// path/to/whatever does not exist
			err = os.MkdirAll( dir,  0777 )
			if ( err != nil ){
				panic( err )
			}
		}


		err := ioutil.WriteFile( fpath ,[]byte ( body )  , 0644)
		if config.REDIS_IMAGE {
			p.RedisServerNoResponse( req , body )
		}else{
			p.RedisServerNoResponse( req , "1" )
		}
		p.Check( err )
		break
	}

}
