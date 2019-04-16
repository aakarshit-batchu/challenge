package controllers

import (
        "github.com/astaxie/beego"
	"challenge/models"
)

type AboutController struct {
        beego.Controller
}


// @Title GetAll
// @Description list items in the inventory
// @Success 200 {object} models.Response
// @Failure 400 error response
// @router / [get]
func (this *AboutController) Get() {
        var status string = "Success"
        var msg, result interface{} = struct{Author string}{"Naga Sai Aakarshit Batchu"}, struct{About string}{"Welcome to Nokia Research and Development Order Management System"}
        var res models.Response
	
	res = models.Response{Status: status, Message: msg, Result: result}
	this.Data["json"] = res
        this.ServeJSON()
}

