package main
import (
	"github.com/johnlion/sites/web"
	"github.com/johnlion/sites/forwardserver"
)

func main(){
	ObjWeb := web.Web_constract()
	ObjForwardserver := forwardserver.Forwardserver_constract()

	done := make(chan bool)
		go ObjWeb.RunWebServer()
		go ObjForwardserver.RunWebServer()
	<-done

}


