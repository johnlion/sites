package main
import (
	"github.com/johnlion/sites/web"

)

func main(){
	ObjWeb := web.Web_constract()


	done := make(chan bool)
	go ObjWeb.RunWebServer()

	<-done

}
