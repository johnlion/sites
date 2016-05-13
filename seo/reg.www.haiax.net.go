package seo
import(
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"fmt"
	"github.com/johnlion/sites/config"
	"html/template"

)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月9日 上午10:22
 * Func: Reg_www_haiax_net
 * Desc: 正则替换规则
 *********************************************/
func ( s *Seo ) Reg_www_haiax_net(){
	s.RegParterns["test1"] = "test1"
	s.RegParterns["test2"] = "test2"
	s.RegParterns["test3"] = "test3"
	s.RegParterns["test4"] = "test4"
	s.RegParterns["test5"] = "test5"



}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月9日 上午10:23
 * Func: Reg_Spider_www_haiax_net
 * Desc: 正则查找规则
 *********************************************/
func ( s*Seo ) Reg_Spider_www_haiax_net( text string  ) {
	expr := s.RequestURI
	switch expr {
	case "/":       //首页
		s.Reg_Spider_www_haiax_net_index( text )
		break
	default:

		break
	}

}

func ( s*Seo ) Reg_Spider_www_haiax_net_default( text string  ) {

}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月10日 下午4:23
 * Func: 首页正则
 * Desc:
 *********************************************/
func ( s*Seo ) Reg_Spider_www_haiax_net_index( text string  ) {
	doc, err :=goquery.NewDocumentFromReader( bytes.NewBufferString( text ) )
	s.Check( err )
	if config.DEBUG{
		s.Debug(s.HOST )
		s.Debug(s.RequestURI )
	}

	/* Initial Var */
	map0 :=  make( map[string]map[string]map[int]map[string]string )
	map1 := make( map[string]map[int]map[string]string )
	map2 := make( map[int]map[string]string   )

	/* map0["site"] 页面基础资料 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	mapAttr := make( map[string]string ) //重构
	mapAttr["title"] = "海岸线文学网_海岸线小说网_txt电子书下载_最新章节"
	mapAttr["keywords"] = "海岸线文学网,免费小说阅读,免费小说下载,最新小说"
	mapAttr["description"] = "海岸线文学网提供海岸线文学网最热门的小说，海岸线文学网是众多网友自发建立，多位热心书友上传很多珍藏孤本小说，海岸线文学网最懂你的小说网。"
	mapAttr["host"] = "http://" + s.HOST
	mapAttr["resourceHost"] = "http://127.0.0.1:8090/"
	map2[0] = mapAttr
	map1["site"] = map2
	map0["siteList"] = map1


	/* map0["recentUpdateList"] 首页最新更新列表 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#newscontent .l .kind" ).Each(func(i int, g *goquery.Selection) {//类别
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		map2[i] = mapAttr
	})
	map1["kind"] =  map2

	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#newscontent .l .chap" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		mapAttr["href"],_ = g.Find("a").Attr("href")
		mapAttr["title"],_ = g.Find("a").Attr("title")
		mapAttr["target"],_ = g.Find("a").Attr("target")
		map2[i] = mapAttr
	})
	map1["chap"] =  map2

	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#newscontent .l .author" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		map2[i] = mapAttr
	})
	map1["author"] =  map2

	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#newscontent .l .time" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		map2[i] = mapAttr
	})
	map1["time"] =  map2

	map0["recentUpdateList"] = map1         //data assets

	/* map0["newStoragedList"] 首页最新入库列表 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#newscontent .r .s2" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		mapAttr["href"],_ = g.Find("a").Attr("href")
		mapAttr["title"],_ = g.Find("a").Attr("title")
		mapAttr["target"],_ = g.Find("a").Attr("target")
		map2[i] = mapAttr
	})
	map1["bookname"] =  map2

	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#newscontent .r .s3" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		map2[i] = mapAttr
	})
	map1["author"] =  map2

	map0["newStoragedList"] = map1         //data assets

	/* map0["firendlinkList"] 首页最友情链接 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( ".firendlink a" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		mapAttr["href"],_ = g.Attr("href")
		mapAttr["title"],_ = g.Attr("title")
		mapAttr["target"],_ = g.Attr("target")
		map2[i] = mapAttr
	})
	map1["friendlink"] =  map2

	map0["firendlinkList"] = map1         //data assets

	/* map0["hotbook"] 首页热门书排行 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#kmod-hotbook .kmod-body .c3" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["href"],_ = g.Find(".box .pic_txt_list .pic a").Attr("href")
		mapAttr["imageSrc"],_ = g.Find(".box .pic_txt_list .pic a img").Attr("src")
		mapAttr["imageWidth"],_ = g.Find(".box .pic_txt_list .pic a img").Attr("width")
		mapAttr["imageHeight"],_ = g.Find(".box .pic_txt_list .pic a img").Attr("height")
		mapAttr["imageAlt"],_ = g.Find(".box .pic_txt_list .pic a img").Attr("alt")
		mapAttr["title"] = g.Find(".box .pic_txt_list .pic a h3 a span").Text()
		mapAttr["info"] = g.Find(".box .pic_txt_list .pic a p .info span").Text()
		map2[i] = mapAttr
	})
	map1["hotbook"] =  map2
	map0["hotbookList"] = map1         //data assets


	/* map0["hotList"] 首页热门排行 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#main #hotlist .l .item" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构

		mapAttr["target"],_ = g.Find( ".image" ).Attr( "target" )
		mapAttr["title"],_ = g.Find( ".image" ).Attr( "title" )
		mapAttr["imageSrc"],_ = g.Find( ".image img" ).Attr( "src" )
		mapAttr["imageWidth"],_ = g.Find( ".image img" ).Attr( "width" )
		mapAttr["imageHeight"],_ = g.Find(".image img").Attr("height")
		mapAttr["imageAlt"],_ = g.Find(".image img").Attr("alt")
		mapAttr["title"] = g.Find("dl dt a").Text()
		mapAttr["href"],_ = g.Find( "dl dt a" ).Attr( "href" )
		mapAttr["info"] = g.Find("dl dd").Text()

		map2[i] = mapAttr

	})
	map1["hot"] =  map2
	map0["hotList"] = map1         //data assets

	/* map0["hotbook"] 首页小说热榜 */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#main #hotlist .r .s2" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		mapAttr["href"],_ = g.Find("a").Attr("href")
		mapAttr["title"],_ = g.Find("a").Attr("title")
		mapAttr["target"],_ = g.Find("a").Attr("target")
		map2[i] = mapAttr
	})

	map1["hotbook"] =  map2
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#main #hotlist .r .s3" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		map2[i] = mapAttr
	})
	map1["hotTop"] =  map2
	map0["hotTopList"] = map1         //data assets


	/* map0["menuList"] 首页Menu */
	map1 = make( map[string]map[int]map[string]string )//重构
	map2 = make( map[int]map[string]string   )//重构
	doc.Find( "#kui-nav ul li" ).Each(func(i int, g *goquery.Selection) {
		mapAttr := make( map[string]string ) //重构
		mapAttr["text"] = g.Text()
		mapAttr["href"],_ = g.Find("a").Attr("href")
		mapAttr["title"],_ = g.Find("a").Attr("title")
		mapAttr["target"],_ = g.Find("a").Attr("target")
		map2[i] = mapAttr
	})
	map1["menu"] =  map2
	map0["menuList"] = map1         //data assets



	fmt.Println( map0["menuList"] )




	/* 模板 */
	tpl, err := template.ParseFiles( "template/novel/index.html" )
	s.Check( err )
	s.tpl = tpl

	var html bytes.Buffer


	s.tpl.Execute( &html,map0 )
	s.text = html.String()








	//map1["recentUpdateList"] = map2

	//map2 = make( map[int]map[string]string  )
	//map3 = make( map[string]string  )
	///* map1["recentUpdateList"] 首页最新入库小说 */
	//doc.Find( "#newscontent .r .s2" ).Each(func(i int, g *goquery.Selection) {
	//	map3["title"] = g.Text()
	//	map2[i] = map3
	//
	//})
	//doc.Find( "#newscontent .r .s3" ).Each(func(i int, g *goquery.Selection) {
	//	map3["author"] = g.Text()
	//	map2[i] = map3
	//
	//})
	//
	//
	//



}