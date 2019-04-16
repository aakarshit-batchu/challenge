package models

import (
	"strings"
	"github.com/astaxie/beego/logs"
        "github.com/astaxie/beego/orm"
        _ "github.com/mattn/go-sqlite3"
)

type OrderRequest struct {
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
}

type Customer struct {
        Name  string    `json:"name"`
        Address string   `json:"address"`
	Phonenumber int64 `json:"phonenumber"`
	Order []OrderRequest `json:"order"`
}

type Invoice struct {
	Name string `json:"name"`
	Total float32 `json:"total"`
	Taxes float32 `json:"taxes"`
	Charges float32 `json:"charges"`
	AmountTotal float32 `json:"amounttotal"`
}

type Response struct {
	Status string `json:"status"`
	Message interface{} `json:"message"`
	Result interface{} `json:"result"`
}

func BillOrder(customer Customer) (Invoice, error) {

        o := orm.NewOrm()
        o.Using("default")

        initTableErr := orm.RunSyncdb("default", false, false)
        if initTableErr != nil {
                logs.Error(initTableErr)
                return Invoice{}, initTableErr
        }

	logs.Debug("Billing the Order for customer: " + customer.Name)
	var total, amountTotal float32
	var taxes, charges float32 = 0.18, 0.05
	var err error
	for _, order := range customer.Order {
		inventory := Inventory{Item: strings.ToLower(order.Item)}
		readErr := o.Read(&inventory)
		if readErr != nil {
			if readErr == orm.ErrNoRows {
				logs.Error("Item doesn't exist in the Inventory: " + order.Item + " error: ", readErr)
				err = readErr
			} else {
				logs.Error(readErr)
				err = readErr
			}
			break
		}
		itemPrice := inventory.Price
		totalitemPrice := itemPrice * float32(order.Quantity)
		total = total + totalitemPrice
	}

	if total >= 500 {
		charges = 0
	}
	amountTotal = total + (total * taxes)	+ (total * charges)
	invoice := Invoice{Name: customer.Name, Total: total, Taxes: taxes*total, Charges: charges*total, AmountTotal: amountTotal}
	return invoice, err
}
