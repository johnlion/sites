package forwardserver

import (
	"net/http"

	"path/filepath"

	"github.com/johnlion/sites/config"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月13日 下午2:18
 * Func: ServeHTTP
 * Desc: load file stream to browser
 *********************************************/
func (f *Forwardserver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := config.TEMPLATE_BASE_DIR + req.URL.Path
	f.Debug( filepath.Join("novel", path) )
	http.ServeFile(res, req, path)
}