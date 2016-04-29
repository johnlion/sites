package proxy

import (
	"github.com/johnlion/sites/config"

	"io/ioutil"
	"os"
	"path/filepath"
	//"strings"
	"regexp"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月28日 上午11:36
 * Func: SaveImage
 * Desc: 保存文件到指定位置
 *********************************************/
func ( p *Proxy ) SaveToFile( requestURI string ,body string  ){
	for{
		/*
		p.Debug( p.LocalDomain )
		p.Debug( requestURI )
		p.Debug( config.CACHE_DIR +  config.IMAGE_DOMAIN_1 + requestURI )
		*/

		url :=  config.IMAGE_DOMAIN_1 + requestURI

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
		p.Check( err )
		break
	}

}
