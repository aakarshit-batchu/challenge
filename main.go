package main

import (
	"github.com/astaxie/beego"
	_ "challenge/routers"
)

func main() {
        if beego.BConfig.RunMode == "dev" {
                beego.BConfig.WebConfig.DirectoryIndex = true
                beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
        }
	beego.Run()
}
