// @APIVersion 1.0.0
// @Title Nokia Research and Development Order Management System
// @Description Order Management System Micro-Service which lets you create an Inventory and also generate Invoice on Orders
// @Contact uguessmyid@gmail.com
// @License Aakarshit 1.0

package routers

import (
        "github.com/astaxie/beego"
        "challenge/controllers"
)

func init() {
        ns := beego.NewNamespace("/v1",
                beego.NSNamespace("/inventory",
                        beego.NSInclude(
                                &controllers.InventoryController{},
                        ),
                ),
                beego.NSNamespace("/order",
                        beego.NSInclude(
                                &controllers.OrderController{},
                        ),
                ),
                beego.NSNamespace("/about",
                        beego.NSInclude(
                                &controllers.AboutController{},
                        ),
                ),
        )
        beego.AddNamespace(ns)
}
