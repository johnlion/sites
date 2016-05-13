package forwardserver

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
func ( f *Forwardserver ) RunWebServer( ){
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

	/* 多服务器组设置 正向  */
	f.Server =  map[int]*http.Server{
		1: &http.Server{ Addr: "127.0.0.1:9990", Handler: f,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},

		//more server...
	}

	/* 协程挂起服务器 */
	done := make(chan bool)
	for key,val := range f.Server{
		go f.Debug( "Server [" +  strconv.Itoa(key) + "] Start111 Listen And Server ...... " )
		go f.LogText( "Server [" +  strconv.Itoa(key) + "] Start Listen And Server ...... " )
		go val.ListenAndServe()
	}

	<-done



}