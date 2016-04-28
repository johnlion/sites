package main
import (
	"github.com/johnlion/sites/web"
	"fmt"
)

func main(){
	ObjWeb := web.Web_constract()
	ObjWeb.RunWebServer()
	fmt.Println( *ObjWeb.Domain )
}

