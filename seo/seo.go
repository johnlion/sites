
package seo

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月24日 上午11:56
 * File: seo.go
 * Desc:
 *********************************************/
import (

)

type Seo struct{
	Domain string
	Title string
	Description string
	Keywords string
	RegParterns map[string]string
}


/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月24日 下午2:23
 * File:
 * Desc:构造函数
 *********************************************/
func Seo_constract( target string, protocol string ) *Seo {
	var seo Seo
	seo.SetKeywords( "this is a keys for objSeo" )
	seo.SetDescription( "this is a desps for objSeo" )
	//seo.RegParterns = map[string]string{ "A":"This is a", "B":"This is b"}
	seo.SetRegParterns( map[string]string{
		`href="`:`href=[\"\']` + protocol + "://" + target,
		"charset=utf-8": `charset=[a-z0-9]{0,10}`,
	})
	return &seo
}

func ( s *Seo ) SetKeywords( Keywords string ){
	s.Keywords = Keywords
}

func ( s *Seo ) SetDescription( Description string ){
	s.Description = Description
}

func ( s *Seo ) SetRegParterns( RegParterns map[string]string  ){
	s.RegParterns = RegParterns
}






