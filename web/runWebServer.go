package web

import (
	"net/http"
	//"log"
	"github.com/johnlion/sites/config"
	"strconv"
)
/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月25日 上午11:09
 * Func: RunWebServer
 * Desc: web package entry  运行并监听
 *********************************************/
func ( w *Web ) RunWebServer( ){
	//http.ListenAndServe( DOMAIN_ADDR ,  w )
	/*
	w.Server = &http.Server{
		Addr:           config.DOMAIN_ADDR,
		Handler:        w,
		ReadTimeout:    config.HTTP_READ_TIME_OUT,
		WriteTimeout:   config.HTTP_WRITE_TIME_OUT,
		MaxHeaderBytes: 1 << 20,
	}
	*/

	/* 多服务器组设置  */
	w.Server =  map[int]*http.Server{
		1: &http.Server{ Addr: config.DOMAIN_ADDR, Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,  },
		2: &http.Server{ Addr: "127.0.0.1:9091", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		3: &http.Server{ Addr: "127.0.0.1:9092", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
	}

	/* 协程挂起服务器 */
	done := make(chan bool)
	for key,val := range w.Server{
		go w.Debug( "Server [" +  strconv.Itoa(key) + "] Start Listen And Server ...... " )
		go w.Log( "Server [" +  strconv.Itoa(key) + "] Start Listen And Server ...... " )
		go val.ListenAndServe()


	}
	<-done



}