package proxy

import (
	"path/filepath"
	"github.com/johnlion/sites/config"
	"regexp"
	"net/http"
)
/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月28日 上午11:48
 * Func: RequestUrLFileGroup
 * Desc: 分组储存http流数据
 *********************************************/
func (p *Proxy ) RequestUrLFileGroupSave( req *http.Request , content string ){
	if ( config.FIELSAVE ){
		extension := filepath.Base( req.URL.RequestURI()  )
		reg := regexp.MustCompile( config.REG_FILE_SUFFIX )
		extension = reg.FindString( extension )

		switch extension {
		case ".png":
			//do code
			p.SaveToImage( req ,content  )
			break
		case ".jpg":
			//do code
			p.SaveToImage( req  ,content )
			break
		case ".jpeg":
			//do code
			p.SaveToImage( req  ,content )
			break
		case ".gif":
			p.SaveToImage( req  ,content )
			//do code
			break
		case ".svg":
			p.SaveToImage( req  ,content )
			//do code
			break
		case ".bmp":
			p.SaveToImage( req  ,content )
			//do code
			break
		case ".tiff":
			p.SaveToImage( req  ,content )
			break
		case ".webp":
			p.SaveToImage( req  ,content )
			//do code
			break
		case ".ico":
			p.SaveToImage( req  ,content )
			//do code
			break
		case ".vico":
			p.SaveToImage( req  ,content )
			//do code
			break
		case ".js":
			p.SaveToFile( req  ,content )
			//do code
			break
		case ".html":
			p.SaveToFile( req  ,content )
			//do code
			break
		case ".htm":
			p.SaveToFile( req  ,content )
			break
		default:
			p.DefaultToFile( req , content )


			//save to mongodb
			break
		}
	}
}


