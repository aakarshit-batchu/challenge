package controllers

import (
        "challenge/models"
        "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

// Operations about Orders
type OrderController struct {
        beego.Controller
}


// @Title Create
// @Description order items from inventory
// @Param       body            body    models.Customer   true            "order items from inventory"
// @Success 200 {object} models.Response
// @Failure 400 failed to place the order
// @router / [post]
func (this *OrderController) Post() {
	req := models.Customer{}
	var status string
	var msg, result interface{}
	var res models.Response
	logs.Debug("Request Body: ", string(this.Ctx.Input.RequestBody))
        if unmarshalErr := json.Unmarshal(this.Ctx.Input.RequestBody, &req); unmarshalErr != nil {
                this.Ctx.Output.SetStatus(400)
                this.Ctx.Output.Body([]byte("empty body"))
		logs.Error("Error: ", unmarshalErr)
                return
        }
	logs.Debug("Request: ", req)
	response, billErr := models.BillOrder(req)
	if billErr != nil {
		status = "Failed"
		msg = "Unable to place the Order: " + billErr.Error()
		result = ""
		this.Ctx.Output.SetStatus(400)
	} else {
		status = "Success"
		msg = ""
		result = struct{Invoice models.Invoice}{response}
	}
	res = models.Response{Status: status, Message: msg, Result: result}
	logs.Debug("Response: ", res)
	this.Data["json"] = res
	this.ServeJSON()
}

