package models

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
        "github.com/astaxie/beego/orm"
        _ "github.com/mattn/go-sqlite3"
)

type Inventory struct {
        Item  string    `orm:"pk" json:"item"`
        Price float32   `json:"price"`
	Category string `json:"category"`
}

func init() {
        orm.RegisterDriver("sqlite3", orm.DRSqlite)
        orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("inventoryDB"))
        orm.RegisterModel(new(Inventory))
	logs.Debug("Initiating Inventory DB: ", beego.AppConfig.String("inventoryDB"))
}


func ConvertpkLower(bulkitems []Inventory) ([]Inventory) {
	for key, item := range bulkitems {
		bulkitems[key].Item = strings.ToLower(item.Item)
	}
	return bulkitems
}

func ListInventory() ([]Inventory, error) {
        o := orm.NewOrm()
        o.Using("default")

	initTableErr := orm.RunSyncdb("default", false, false)
	if initTableErr != nil {
		logs.Error(initTableErr)
		return nil, initTableErr
	}

	logs.Debug("Executing List Inventory...")
        var listInventory []Inventory
        _, listErr := o.Raw("SELECT item, price, category FROM inventory").QueryRows(&listInventory)
        if listErr != nil {
		logs.Error(listErr)
                return nil, listErr
        }
        return listInventory, nil
}

func CreateInventory(bulkitems []Inventory) (int64, error){
        o := orm.NewOrm()
        o.Using("default")

        initTableErr := orm.RunSyncdb("default", false, false)
        if initTableErr != nil {
                logs.Error(initTableErr)
                return 0, initTableErr
        }

        logs.Debug("Bulk Insert of Records to Inventory...")

	successNums, bulkinsertErr := o.InsertMulti(1000, bulkitems)
	if bulkinsertErr != nil {
		logs.Error(bulkinsertErr)
		return 0, bulkinsertErr
	}
	return successNums, nil
}

func UpdateInventory(bulkitems []Inventory) (int64, error) {
        o := orm.NewOrm()
        o.Using("default")

        initTableErr := orm.RunSyncdb("default", false, false)
        if initTableErr != nil {
                logs.Error(initTableErr)
                return 0, initTableErr
        }

        logs.Debug("Updating records in the Inventory...")
	var updatedNum int64
	var err error
	for _, item := range bulkitems {
		sucessNum, updateErr := o.Update(&item)
		if updateErr != nil {
			logs.Error(updateErr)
			err = updateErr
			break
		}
		if sucessNum == 0 {
			_, insertErr := o.Insert(&item)
			if insertErr != nil {
				logs.Error(insertErr)
				err = insertErr
			} else {
				sucessNum = 1
			}
		}
		updatedNum = updatedNum + sucessNum
	}
	return updatedNum, err
}

func DeleteInventory(bulkitems []Inventory) (int64, error) {
        o := orm.NewOrm()
        o.Using("default")

        initTableErr := orm.RunSyncdb("default", false, false)
        if initTableErr != nil {
                logs.Error(initTableErr)
                return 0, initTableErr
        }

        logs.Debug("Deleting records in the Inventory...")
        var updatedNum int64
	var err error
	for _, item := range bulkitems {
		item.Item = strings.ToLower(item.Item)
		sucessNum, deleteErr := o.Delete(&item)
		if deleteErr != nil {
			logs.Error(deleteErr)
			err = deleteErr
			break
		}
		updatedNum = updatedNum + sucessNum
	}
	return updatedNum, err
}
