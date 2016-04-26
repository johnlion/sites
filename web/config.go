package web

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

	HTTP_READ_TIME_OUT =  10 * time.Second
	HTTP_WRITE_TIME_OUT  =  10 * time.Second

	DEBUG = true
	LOG_WRITE = true
	LOG_BASE_DIR = "/tmp/log/web/"
	LOG_DIR_FILE = LOG_BASE_DIR + DOMAIN


)