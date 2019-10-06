package main

import (
	"github.com/astaxie/beego"
	"net/http"
	_ "phoneLocationOpenApi/routers"
)
func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("<html>" +
		"<h1 style=\"text-align:center\">" +
		"SORRY UNAUTHORIZED" +
		"</h1>" +
		"<h1 style=\"text-align:center\">" +
		"PLEASE CHECK AGAIN" +
		"</h1>" +
		"</html>"))
	return
}

func main() {

	beego.ErrorHandler("401", PageNotFound)
	beego.ErrorHandler("403", PageNotFound)
	beego.ErrorHandler("404", PageNotFound)
	beego.ErrorHandler("500", PageNotFound)
	beego.ErrorHandler("503", PageNotFound)

	//models.Init(time.Hour * 24)
	beego.Run()
}
