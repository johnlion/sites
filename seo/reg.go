package seo
import (
	"regexp"
	"fmt"
)

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func: RegProcess
 * desc: 返回[]string
 * ****************************************/
func ( s *Seo )RegProcess( text string ) []byte {
	/* 取得seo数据 */
	for i,val := range s.RegParterns{
		fmt.Println( i )
		fmt.Println( val )
		reg := regexp.MustCompile( val )
		text = reg.ReplaceAllString(text,  i  )

	}
	return []byte(text)

}