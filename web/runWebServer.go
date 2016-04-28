package web

import (
	"net/http"
	"log"
	"github.com/johnlion/sites/config"
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
	w.Server = &http.Server{
		Addr:           config.DOMAIN_ADDR,
		Handler:        w,
		ReadTimeout:    config.HTTP_READ_TIME_OUT,
		WriteTimeout:   config.HTTP_WRITE_TIME_OUT,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(w.Server.ListenAndServe())

}