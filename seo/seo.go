package seo

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月24日 上午11:56
 * File: seo.go
 * Desc:
 *********************************************/
import ()
import (
	"net/http"
	"os"
	"fmt"
	"github.com/johnlion/sites/config"
	"strings"
)

type Seo struct {
	Domain      string
	Title       string
	Description string
	Keywords    string
	RegParterns map[string]string
}


/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月24日 下午2:23
 * File:
 * Desc:构造函数
 *********************************************/
func Seo_constract( req *http.Request,target string, scheme string) *Seo {
	var seo Seo
	seo.SetKeywords("this is a keys for objSeo")
	seo.SetDescription("this is a desps for objSeo")
	//seo.RegParterns = map[string]string{ "A":"This is a", "B":"This is b"}
	seo.SetRegParterns(map[string]string{
		`href="`:        `href=[\"\']` + scheme + "://" + target,
		"charset=utf-8": `charset=[a-z0-9]{0,10}`,
	})

	invokeMethod := "Reg_" +  strings.Replace(  target, ".", "_", -1  )

	seo.Debug( invokeMethod )
	InvokeObjectMethod( &seo, invokeMethod )

	if config.SEO_DEBUG {
		for key,val := range seo.RegParterns{
			fmt.Println( "   " +  key + " [>] " + val )
		}
	os.Exit(1)
	}
	return &seo
}

func (s *Seo) SetKeywords(Keywords string) {
	s.Keywords = Keywords
}

func (s *Seo) SetDescription(Description string) {
	s.Description = Description
}

func (s *Seo) SetRegParterns(RegParterns map[string]string) {
	s.RegParterns = RegParterns
}
