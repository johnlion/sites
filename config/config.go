package config


import "time"
/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午11:32
 * File: config.go
 * Desc: web package 配置文件
 *********************************************/



const (
	TARGET = "www.baidu.com"
	PROTOCAL = "http"       //"http|https|sockets|..."

	DOMAIN = "127.0.0.1"                    //域名
	DOMAIN_PORT = "9090"                    //域名端口
	DOMAIN_ADDR = DOMAIN + ":"  + DOMAIN_PORT//域名完整地址

	IMAGE_DOMAIN_1 = "127.0.0.1:9090"       //图片服务器1
	IMAGE_DOMAIN_2 = "127.0.0.1:9090"       //图片服务器2

	HTTP_READ_TIME_OUT =  10 * time.Second
	HTTP_WRITE_TIME_OUT  =  10 * time.Second

	DEBUG = true
	LOG_WRITE = true
	LOG_BASE_DIR = "/tmp/log/web/"
	LOG_DIR_FILE = LOG_BASE_DIR + DOMAIN

	PROXY_STATUS = false                    //开启/关闭代理

	CACHE_DIR = "cache/"

	FILE_DEFAULT_NAME = "index.html"                 //无文件名称使用此名称
	FIELSAVE = true                         //文件保存到本地开启|关闭


)

const(
	REG_FILE_SUFFIX = "(.vico)|(.png)|(.jpg)|(.jpeg)|(.gif)|(.svg)|(.bmp)|(.tiff)|(.webp)|(.ico)|(.js)|(.php)|(.asp)"
	REG_TOTAL_FILENAME ="[a-z0-9._:/#]{1,999}[" + REG_FILE_SUFFIX + "]"
)