package web

import (
	"net/http"
	//"log"
	"github.com/johnlion/sites/config"
	"strconv"

	"net"
	"log"
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

	/* 多服务器组设置 反向  */
	w.Server =  map[int]*http.Server{
		1: &http.Server{ Addr: config.LOCAL_DOMAIN_BIND_IP_PORT_1, Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//2: &http.Server{ Addr: config.LOCAL_RESOURCE_DOMAIN_BIND_IP_PORT_1, Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//3: &http.Server{ Addr: "127.0.0.1:9092", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//4: &http.Server{ Addr: "127.0.0.1:9093", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//5: &http.Server{ Addr: "127.0.0.1:9094", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//6: &http.Server{ Addr: "127.0.0.1:9095", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//7: &http.Server{ Addr: "127.0.0.1:9096", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//8: &http.Server{ Addr: "127.0.0.1:9097", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},
		//9: &http.Server{ Addr: "127.0.0.1:9098", Handler: w,ReadTimeout:    config.HTTP_READ_TIME_OUT, WriteTimeout:   config.HTTP_WRITE_TIME_OUT, MaxHeaderBytes: 1 << 20,},

		//more server...
	}

	/* 协程挂起服务器 */
	done := make(chan bool)
	for key,val := range w.Server{
		go w.Debug( "Server [" +  strconv.Itoa(key) + "] Start Listen And Server ...... " )
		go w.LogText( "Server [" +  strconv.Itoa(key) + "] Start Listen And Server ...... " )
		//go val.ListenAndServe()
		listener, err := net.Listen("tcp",  val.Addr )
		if err != nil {
			log.Fatal( err )

		}
		go val.Serve( listener )
	}

	<-done



}