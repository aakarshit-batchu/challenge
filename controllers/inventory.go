package controllers

import (
        "challenge/models"
        "github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"strconv"
)

// Operations about inventory
type InventoryController struct {
        beego.Controller
}

// @Title Get
// @Description list items in the inventory
// @Success 200 {object} models.Response
// @Failure 400 cannot retrieve inventory
// @router / [get]
func (this *InventoryController) Get() {
	response, listErr := models.ListInventory()
	var status string
	var msg, result interface{}
	var res models.Response
        if listErr != nil {
		status = "Failed"
		msg = listErr
		result = ""
        } else {
		status = "Success"
		msg = ""
		result = struct{ Inventory []models.Inventory }{response}
        }
	res = models.Response{Status: status, Message: msg, Result: result}	
	logs.Debug("Response: ", res)
	this.Data["json"] = res
        this.ServeJSON()
}

// @Title Create
// @Description Add New Items to Inventory
// @Param       body            body    []models.Inventory   true            "Items as JSON Array to Inventory"
// @Success 200 {object} models.Response
// @Failure 400 cannot add items to inventory
// @router / [post]
func (this *InventoryController) Post() {
	req := []models.Inventory{}
        if unmarshalErr := json.Unmarshal(this.Ctx.Input.RequestBody, &req); unmarshalErr != nil {
                this.Ctx.Output.SetStatus(400)
                this.Ctx.Output.Body([]byte("empty body"))
                return
        }
	formattedreq := models.ConvertpkLower(req)
	response, createErr := models.CreateInventory(formattedreq)
	var status string
	var msg, result interface{}
	var res models.Response
	if createErr != nil {
                status = "Failed"
                msg = "Inventory items should be unique: " + createErr.Error()
                result = ""
		this.Ctx.Output.SetStatus(400)
	} else {
		status = "Success"
		msg = ""
		result = "Sucessfully inserted " + strconv.FormatInt(response, 10) + " records"
	}
	res = models.Response{Status: status, Message: msg, Result: result}
	logs.Debug("Response: ", res)
	this.Data["json"] = res
	this.ServeJSON()
}

// @Title Update
// @Description update the inventory
// @Param       body            body    []models.Inventory   true            "update items in inventory"
// @Success 200 {object} models.Response
// @Failure 400 failed to update items!
// @router / [put]
func (this *InventoryController) Put() {
        req := []models.Inventory{}
        if unmarshalErr := json.Unmarshal(this.Ctx.Input.RequestBody, &req); unmarshalErr != nil {
                this.Ctx.Output.SetStatus(400)
                this.Ctx.Output.Body([]byte("empty body"))
                return
        }
	formattedreq := models.ConvertpkLower(req)
        response, updateErr := models.UpdateInventory(formattedreq)
        var status string
        var msg, result interface{}
        var res models.Response
        if updateErr != nil {
                status = "Failed"
                msg = "Update Error: " + updateErr.Error() + " , updated " + strconv.FormatInt(response, 10) + " records"
                result = ""
		this.Ctx.Output.SetStatus(400)
        } else {
		status = "Success"
		msg = ""
                result = "Sucessfully updated " + strconv.FormatInt(response, 10) + " records"
        }
        res = models.Response{Status: status, Message: msg, Result: result}
	logs.Debug("Response: ", res)
        this.Data["json"] = res
        this.ServeJSON()
}

// @Title Delete
// @Description delete the item in inventory
// @Param       body            body    []models.Inventory   true            "delete items in inventory"
// @Success 200 {object} models.Response
// @Failure 400 failed to delete the item!
// @router / [delete]
func (this *InventoryController) Delete() {
        req := []models.Inventory{}
        if unmarshalErr := json.Unmarshal(this.Ctx.Input.RequestBody, &req); unmarshalErr != nil {
                this.Ctx.Output.SetStatus(400)
                this.Ctx.Output.Body([]byte("empty body"))
                return
        }
	formattedreq := models.ConvertpkLower(req)
        response, deleteErr := models.DeleteInventory(formattedreq)
        var status string
        var msg, result interface{}
        var res models.Response
        if deleteErr != nil {
		status = "Failed"
                msg = "Delete Error: " + deleteErr.Error() + " , delete " + strconv.FormatInt(response, 10) + " records"
		result = ""
		this.Ctx.Output.SetStatus(400)
        } else {
		status = "Success"
		msg = ""
                result = "Sucessfully deleted " + strconv.FormatInt(response, 10) + " records"
        }
        res = models.Response{Status: status, Message: msg, Result: result}
	logs.Debug("Response: ", res)
        this.Data["json"] = res
        this.ServeJSON()
}

