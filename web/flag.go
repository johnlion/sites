package web

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午11:33
 * File: flag.go
 * Desc: web package 命令行参数配置
 *********************************************/
import "flag"

var (
	target *string = flag.String("target", TARGET, "target remote url")
	scheme *string = flag.String("scheme", PROTOCAL, "Secure Hypertext Transfer Scheme")
)
