package proxy

import (
	"path/filepath"
	"github.com/johnlion/sites/config"
	"regexp"

)
/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月28日 上午11:48
 * Func: RequestUrLFileGroup
 * Desc: 分组储存http流数据
 *********************************************/
func (p *Proxy ) RequestUrLFileGroupSave( requestURI string , content string ){
	if ( config.FIELSAVE ){
		extension := filepath.Base( requestURI )
		reg := regexp.MustCompile( config.REG_FILE_SUFFIX )
		extension = reg.FindString( extension )

		switch extension {
		case ".png":
			//do code
			p.SaveToImage( requestURI ,content  )
			break
		case ".jpg":
			//do code
			p.SaveToImage( requestURI ,content )
			break
		case ".jpeg":
			//do code
			p.SaveToImage( requestURI ,content )
			break
		case ".gif":
			p.SaveToImage( requestURI ,content )
			//do code
			break
		case ".svg":
			p.SaveToImage( requestURI ,content )
			//do code
			break
		case ".bmp":
			p.SaveToImage( requestURI ,content )
			//do code
			break
		case ".tiff":
			p.SaveToImage( requestURI ,content )
			break
		case ".webp":
			p.SaveToImage( requestURI ,content )
			//do code
			break
		case ".ico":
			p.SaveToImage( requestURI ,content )
			//do code
			break
		case ".vico":
			p.SaveToImage( requestURI ,content )
			//do code
			break
		case ".js":
			p.SaveToFile( requestURI ,content )
			//do code
			break
		default:


			p.HtmlToFile( requestURI, content )


			//save to mongodb
			break
		}
	}
}


