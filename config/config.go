package config

import "time"

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午11:32
 * File: config.go
 * Desc: web package 配置文件
 *********************************************/

/* Flag */
const (
	TARGET = "www.baidu.com"
	SCHEME = "http" //"http|https|sockets|..."
)

const (
	HTTP_READ_TIME_OUT  = 10 * time.Second
	HTTP_WRITE_TIME_OUT = 10 * time.Second

	DEBUG                  = true
	PROXY_STATUS = false //开启/关闭代理
	FILE_DEFAULT_NAME = "index.html" //无文件名称使用此名称
	FIELSAVE          = true         //文件保存到本地开启|关闭

)

const (
	CACHE_DIR = "cache/"
)

/* Log */
const(
	LOG_WRITE              = true
	LOG_BASE_DIR           = "/tmp/log/web"
	LOG_Default_ACCESS_DIR = LOG_BASE_DIR + "/" + "access.log"
)

/* REG */
const (
	REG_FILE_SUFFIX    = "(.vico)|(.png)|(.jpg)|(.jpeg)|(.gif)|(.svg)|(.bmp)|(.tiff)|(.webp)|(.ico)|(.js)|(.php)|(.asp)"
	REG_TOTAL_FILENAME = "[a-zA-Z0-9._:/#]{1,999}[" + REG_FILE_SUFFIX + "]"
)

const (
	DATABASE           = false
	DATABASE_MAIN_TYPE = "mysql"

)

/* REDIS */
const(
	REDIS = false
	REDIS_ADDR = "127.0.0.1"
	REDIS_PORT = "6379"
	REDIS_MAXIDLE = 80
	REDIS_MAXACTIVE = 12000
	REDIS_IMAGE = false
	REDIS_FILE = false
)

/* SEO */
const(
	SEO_DEBUG = true
)

/* Template */
const(
	TEMPLATE_BASE_DIR = "template/novel"
)
const(
	LOCAL_DOMAIN_DOMAIN_1 = "test.sydscience.com"
	LOCAL_DOMAIN_BIND_IP_PORT_1 = ":80"

	LOCAL_RESOURCE_DOMAIN_1 = "res.sydscience.com"
	LOCAL_RESOURCE_DOMAIN_BIND_IP_PORT_1 = ":80"
)