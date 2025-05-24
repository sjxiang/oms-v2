package app

import "github.com/sjxiang/oms-v2/stock/app/query"


type Application struct {
	Commands Commands
	Queries  Queries
}


type Commands struct {

}


type Queries struct {
	CheckIfItemsInStock query.CheckIfItemsInStockHandler
	GetItems            query.GetItemsHandler
}
